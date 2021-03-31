package stage

import (
	"sync"
	"yama.io/yamaIterativeE/internal/iteration/step"
)

type RuntimeStageState int
const (
	Init RuntimeStageState = iota
	Ready
	Running
	// if a stage is Finished it will be removed, so Finish is useless
	Finish
	Failure
)

/** business mapping from db.StageExec to RuntimeStage */
type RuntimeStage struct {
	Id                int64
	IterationActionId int64
	StageId           int64
	StageIndex        int
	PipelineId        int64

	Steps     []*step.RuntimeStep
	TaskNum   int
	ExecPath  string
	Lock      sync.Mutex
	Success   chan step.Message
	Failure   chan step.Message
	State     RuntimeStageState
}


