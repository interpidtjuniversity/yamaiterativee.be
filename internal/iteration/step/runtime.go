package step

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/iteration/step/beanfactory"
)

type RuntimeStepState int
const (
	Init RuntimeStepState = iota
	Running
	Finish
	Failure
	Unknown
	Canceled
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
	case Canceled:
		return "Canceled"
	default:
		return "Unknown"
	}
}

func (r RuntimeStepState) FromString(s string) RuntimeStepState {
	switch s {
	case "Init":
		return Init
	case "Running":
		return Running
	case "Finish":
		return Finish
	case "Failure":
		return Failure
	case "Canceled":
		return Canceled
	default:
		return Unknown
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
	Cond           bool
	Type           string

	NeedUpdate     bool
	Canceled       bool

	Env            *map[string]interface{}
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
	step := db.StepExec{StepId: t.StepId, StageExecId: t.StageExecId, LogPath: t.LogPath, ExecPath: t.ExecPath, State: Running.ToString()}
	switch t.StepId {
	case 11:
		step.Link = (*(t.Env))["mergeRequestCodeReviewUrl"].(string)
		break
	}
	_, _ = db.InsertStepExec(&step)
	t.Id = step.ID
	// exec
	if t.Type == "callBack" {
		// 11. code review
		for !t.Cond && !t.Canceled {
			time.Sleep(time.Duration(5) * time.Second)
		}
	} else if t.Type == "command" {
		// other exec command
		// sleep to 10
		time.Sleep(time.Duration(5)*time.Second)
		basePath, _ := os.Getwd()
		// exec
		ctx := context.Background()
		t.transformArgs(t.Env)
		t.Command = fmt.Sprintf("%s%s",basePath,t.Command)
		commend := exec.CommandContext(ctx, t.Command, t.Args...)
		commend.Dir = t.ExecPath
		log, _ := os.OpenFile(t.LogPath, os.O_CREATE|os.O_WRONLY, 0777)
		commend.Stdout = log
		commend.Stderr = log
		err := commend.Run()
		if err != nil {
			return nil, err
		}
		return nil, err
	} else if t.Type == "code" {
		time.Sleep(time.Duration(5)*time.Second)
		os.OpenFile(t.LogPath, os.O_CREATE|os.O_WRONLY, 0777)
		bean := beanfactory.GetBean(t.Command)
		t.transformArgs(t.Env)
		t.Args = append(t.Args, t.LogPath)
		err := bean.Execute(t.Args, t.Env)
		return nil, err
	}
	return nil,nil
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
	return t.Canceled
}


func (t *RuntimeStep)transformArgs(env *map[string]interface{}) {
	for i:=0; i < len(t.Args); i++ {
		key := (*env)[t.Args[i]]
		t.Args[i] = key.(string)
	}
}
