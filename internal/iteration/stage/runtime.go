package stage

import (
	"sync"
	"yama.io/yamaIterativeE/internal/iteration/step"
)

/** business mapping from db.StageExec to RuntimeStage */
type RuntimeStage struct {
	Id                int64
	IterationActionId int64
	StageId           int64
	StageIndex        int
	PipelineId        int64

	Steps             []*step.RuntimeStep
	TaskNum           int
	ExecPath          string
	Lock              sync.Mutex
	Channel           chan step.Message
	IsRunning         bool
}

func (rs *RuntimeStage)NextStage() []*RuntimeStage {

	return nil
}
