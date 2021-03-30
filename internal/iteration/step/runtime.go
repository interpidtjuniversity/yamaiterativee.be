package step

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"yama.io/yamaIterativeE/internal/db"
)

/** business mapping from db.StepExec to RuntimeStep and task abstract*/
type RuntimeStep struct {
	Id                int64
	StageExecId       int64
	IterationActionId int64
	StepId            int64
	StageId           int64
	PipelineId        int64
	StageIndex        int

	IsCanceled        bool
	LogPath           string
	ExecPath          string
	IsPassed          bool
	Command           string
	Args              []string
	Channel           chan Message
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
	return nil, err
}

func (t *RuntimeStep) Success(result interface{}) {
	fmt.Print(fmt.Sprintf("channel length is:%d\n",len(t.Channel)))
	t.Channel <- Message{
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

}

func (t *RuntimeStep) Cancel() {

}

func (t *RuntimeStep) IsCancel() bool {
	return false
}

