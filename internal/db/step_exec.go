package db

type StepExec struct {
	ID          int64  	`xorm:"id autoincr pk"`
	StageExecId int64   `xorm:"stage_exec_id"`
	LogPath     string 	`xorm:"log_path"`
	StepId      int64   `xorm:"step_id"`
	Passed      bool    `xorm:"is_passed"`
	ExecPath    string  `xorm:"exec_path"`
}

func InsertStepExec(exec *StepExec) (int64, error){
	id, err := x.Insert(exec)
	return id, err
}

func PassStepExec(id int64) error{
	_, err := x.Table(&StepExec{}).ID(id).Update(map[string]interface{}{"is_passed": true})
	return err
}

func FailStepExec(id int64) error{
	_, err := x.Table(&StepExec{}).ID(id).Update(map[string]interface{}{"is_passed": false})
	return err
}