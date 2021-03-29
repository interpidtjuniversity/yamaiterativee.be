package guc

type Task interface {
	Run() (interface{}, error)
	Success(result interface{})
	Failure()
	Cancel()
	IsCancel() bool
}
