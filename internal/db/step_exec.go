package db

type StepExec struct {
	ID int64
	LogPath string
	StepId int64
	Passed bool
}
