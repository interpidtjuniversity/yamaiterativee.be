package db

import "xorm.io/builder"

type StepExec struct {
	ID          int64  	`xorm:"id autoincr pk"`
	StageExecId int64   `xorm:"stage_exec_id"`
	LogPath     string 	`xorm:"log_path"`
	StepId      int64   `xorm:"step_id"`
	Passed      bool    `xorm:"is_passed"`
	ExecPath    string  `xorm:"exec_path"`
	State       string  `xorm:"state"`
	Link        string  `xorm:"link"`
}

func InsertStepExec(exec *StepExec) (int64, error){
	id, err := x.Insert(exec)
	return id, err
}

func UpdateStepExecState(id int64, state string) error {
	stepExec := &StepExec{State: state}
	_ ,err := x.Table("step_exec").Cols("state").Where(builder.Eq{"id":id}).Update(stepExec)
	return err
}

func PassStepExec(id int64) error{
	_, err := x.Table(&StepExec{}).ID(id).Update(map[string]interface{}{"is_passed": true, "state": "Finish"})
	return err
}

func FailStepExec(id int64) error{
	_, err := x.Table(&StepExec{}).ID(id).Update(map[string]interface{}{"is_passed": false, "state": "Failure"})
	return err
}

func GetStepExecByStageExecIdAndStepId(stageExecId, stepId int64) (*StepExec, error) {
	stepExec := &StepExec{}
	has, err := x.Where(builder.Eq{"stage_exec_id" : stageExecId}.And(builder.Eq{"step_id": stepId})).Get(stepExec)
	if !has {
		return nil, err
	}
	return stepExec, nil
}