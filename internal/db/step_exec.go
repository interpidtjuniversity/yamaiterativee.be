package db

type StepExec struct {
	ID          int64  	`xorm:"id autoincr pk"`
	StageExecId int64   `xorm:"stage_exec_id"`
	LogPath     string 	`xorm:"log_path"`
	StepId      int64   `xorm:"is_passed"`
	Passed      bool    `xorm:"is_passed"`
	ExecPath    string  `xorm:"exec_path"`
}
