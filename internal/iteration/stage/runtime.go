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
	Canceled
	// if a stage is Finished it will be removed, so Finish is useless
	Finish
	Failure
	Unknown
)

func (rss RuntimeStageState) FromString(s string) RuntimeStageState {
	switch s {
	case "Init":
		return Init
	case "Ready":
		return Ready
	case "Running":
		return Running
	case "Failure":
		return Failure
	case "Finish":
		return Finish
	case "Canceled":
		return Canceled

	default:
		return Unknown
	}
}

func (rss RuntimeStageState)ToString() string {
	switch rss {
	case Init:
		return "Init"
	case Ready:
		return "Ready"
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

func (rss RuntimeStageState) CanUpdate() bool  {
	switch rss {
	case Running:
		return true
	case Canceled:
		return true
	case Finish:
		return true
	case Failure:
		return true
	default:
		return false
	}
}

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

	NeedUpdate       bool

	Env       *map[string]interface{}
}


