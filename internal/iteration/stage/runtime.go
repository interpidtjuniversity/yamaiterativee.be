package stage

import (
	"sync"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/iteration/pipeline"
	"yama.io/yamaIterativeE/internal/iteration/step"
)

/** business mapping from db.StageExec to RuntimeStage */
type RuntimeStage struct {
	Id                int64
	Steps             []*step.RuntimeStep
	IterationActionId int64
	StageId           int64
	TaskNum           int
	ExecPath          string
	lock              sync.Mutex
}

func FromStage(pipeline *pipeline.RuntimePipeline, stage *db.Stage) *RuntimeStage {
	rs := &RuntimeStage{IterationActionId: pipeline.ID, StageId: stage.ID, TaskNum: len(stage.Steps)}
	steps, _ := db.BranchQueryStepsByIds(stage.Steps)
	var runtimeSteps []*step.RuntimeStep
	for _,s := range steps {
		runtimeSteps = append(runtimeSteps, step.FromStep(rs, s))
	}
	rs.Steps = runtimeSteps
	return rs
}

func (rs *RuntimeStage)NextStage() []*RuntimeStage {

	return nil
}
