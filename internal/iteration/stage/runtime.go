package stage

import "yama.io/yamaIterativeE/internal/iteration/step"

/** business mapping from db.StageExec to RuntimeStage */
type RuntimeStage struct {
	Steps []*step.RuntimeStep
	PipelineId int64

}

func (rs *RuntimeStage)NextStage() []*RuntimeStage {

	return nil
}
