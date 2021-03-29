package step

import (
	"context"
	"os"
	"os/exec"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/iteration/stage"
)

/** business mapping from db.StepExec to RuntimeStep and task abstract*/
type RuntimeStep struct {
	Id          int64
	IsCanceled  bool
	StageExecId int64
	LogPath     string
	ExecPath    string
	IsPassed    bool
	Command     string
	Args        []string
	Channel     chan bool
}

// message when a step finish
type Message struct {
	Id                int64
	StageExecId       int64
	IterationActionId int64
}

func (t *RuntimeStep) Run() (interface{}, error) {
	ctx := context.Background()
	commend := exec.CommandContext(ctx, t.Command, t.Args...)
	commend.Dir = t.ExecPath
	log, _ := os.OpenFile(t.LogPath, os.O_CREATE|os.O_WRONLY, 0777)
	commend.Stdout = log
	commend.Stderr = log
	err := commend.Run()
	if err!=nil {
		return nil, err
	}

}

func (t *RuntimeStep) Success(result interface{}) {
	panic("implement me")
}

func (t *RuntimeStep) Failure() {
	panic("implement me")
}

func (t *RuntimeStep) Cancel() {
	panic("implement me")
}

func (t *RuntimeStep) IsCancel() bool {
	panic("implement me")
}

func FromStep(stage *stage.RuntimeStage, step *db.Step) *RuntimeStep {
	rs := &RuntimeStep{StageExecId: stage.Id, ExecPath: stage.ExecPath, LogPath: FormatLogPath(), Command: step.Command, Args: step.Args}
	return rs
}
