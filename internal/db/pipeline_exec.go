package db

import "time"

type PipeLineExec struct {
	ID         int64
	PipeLineId int64
	UserId     int64
	StartUnix  time.Time
	EndUnix    time.Time

	/** k: the id of step, v the id of step_exec */
	StepLogs   map[int64]int64
}