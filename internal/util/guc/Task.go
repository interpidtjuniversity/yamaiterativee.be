package guc

type Task interface {
	Run() (interface{}, error)
	Success(result interface{})
	Fail()
	Cancel()
	IsCancel() bool
}
