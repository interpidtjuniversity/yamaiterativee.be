package step

/** business mapping from db.StepExec to RuntimeStep and task abstract*/
type RuntimeStep struct {
	IsCanceled bool
}

func (t *RuntimeStep) Run() (interface{}, error) {
	panic("implement me")
}

func (t *RuntimeStep) Success(result interface{}) {
	panic("implement me")
}

func (t *RuntimeStep) Fail() {
	panic("implement me")
}

func (t *RuntimeStep) Cancel() {
	panic("implement me")
}

func (t *RuntimeStep) IsCancel() bool {
	panic("implement me")
}
