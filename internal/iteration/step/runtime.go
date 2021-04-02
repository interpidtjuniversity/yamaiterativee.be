package step

import (
	"context"
	"os"
	"os/exec"
	"yama.io/yamaIterativeE/internal/db"
)

type RuntimeStepState int
const (
	Init RuntimeStepState = iota
	Running
	Finish
	Failure
	Unknown
)

func (r RuntimeStepState) ToString() string {
	switch r {
	case Init:
		return "Init"
	case Running:
		return "Running"
	case Finish:
		return "Finish"
	case Failure:
		return "Failure"
	default:
		return "Unknown"
	}
}

/** business mapping from db.StepExec to RuntimeStep and task abstract*/
type RuntimeStep struct {
	Id                int64
	StageExecId       int64
	IterationActionId int64
	StepId            int64
	StageId           int64
	PipelineId        int64
	StageIndex        int

	IsCanceled     bool
	LogPath        string
	ExecPath       string
	IsPassed       bool
	Command        string
	Args           []string
	SuccessChannel chan Message
	FailureChannel chan Message
	State          RuntimeStepState
}

// message when a step finish
type Message struct {
	Id                int64
	StageExecId       int64
	IterationActionId int64
	StepId            int64
	StageId           int64
	PipelineId        int64
	StageIndex        int
}

func (t *RuntimeStep) Run() (interface{}, error) {
	// write db
	t.State = Running
	step := db.StepExec{StepId: t.StepId, StageExecId: t.StageExecId, LogPath: t.LogPath, ExecPath: t.ExecPath}
	_, _ = db.InsertStepExec(&step)
	t.Id = step.ID
	// exec
	ctx := context.Background()
	commend := exec.CommandContext(ctx, t.Command, t.Args...)
	commend.Dir = t.ExecPath
	log, _ := os.OpenFile(t.LogPath, os.O_CREATE|os.O_WRONLY, 0777)
	commend.Stdout = log
	commend.Stderr = log
	err := commend.Run()
	if err != nil {
		return nil, err
	}
	err = commend.Wait()
	return nil, err
}

func (t *RuntimeStep) Success(result interface{}) {
	t.State = Finish
	t.SuccessChannel <- Message{
		StageIndex: t.StageIndex,
		PipelineId: t.PipelineId,
		StageExecId: t.StageExecId,
		IterationActionId: t.IterationActionId,
		StepId: t.StepId,
		Id: t.Id,
		StageId: t.StageId,
	}
	// update db
	db.PassStepExec(t.Id)
}

func (t *RuntimeStep) Failure() {
	t.State = Failure
	t.FailureChannel <- Message{
		StageIndex: t.StageIndex,
		PipelineId: t.PipelineId,
		StageExecId: t.StageExecId,
		IterationActionId: t.IterationActionId,
		StepId: t.StepId,
		Id: t.Id,
		StageId: t.StageId,
	}
	// update db
	db.FailStepExec(t.Id)
}

func (t *RuntimeStep) Cancel() {

}

func (t *RuntimeStep) IsCancel() bool {
	return false
}

