package db

import "xorm.io/builder"

type StageExec struct {
	ID             int64  `xorm:"id autoincr pk"`
	ActId          int64  `xorm:"act_id"`
	StageId        int64  `xorm:"stage_id"`
	ExecPath       string `xorm:"exec_path"`
	StageStartUnix int64  `xorm:"stage_start_unix"`
	StageEndUnix   int64  `xorm:"stage_end_unix"`
}

func BranchQueryStageExec(iterationActionId int64, stageIds []int64)([]*StageExec, error)  {
	var execs []*StageExec
	err := x.Where(builder.Eq{"act_id": iterationActionId}.And(builder.In("stage_id", stageIds))).Find(&execs)
	if err != nil{
		return nil, err
	}
	return execs, nil
}

func InsertStageExec(exec StageExec) (int64, error){
	id, err := x.Insert(exec)
	return id, err
}
