package db

import (
	"fmt"
	"xorm.io/builder"
)

type StageExec struct {
	Id             int64  `xorm:"id autoincr pk"`
	ActId          int64  `xorm:"act_id"`
	StageId        int64  `xorm:"stage_id"`
	ExecPath       string `xorm:"exec_path"`
	StageStartUnix int64  `xorm:"stage_start_unix"`
	StageEndUnix   int64  `xorm:"stage_end_unix"`
	State          string `xorm:"state"`
}

func QueryStageExec(iterationActionId, stageId int64) (*StageExec, error){
	exec := &StageExec{}
	has, err := x.Table("stage_exec").Cols("id", "state").Where(builder.Eq{"act_id": iterationActionId}.And(builder.Eq{"stage_id": stageId})).Get(exec)
	if !has{
		return nil, err
	}
	return exec, nil
}

func BranchQueryStageExec(iterationActionId int64, stageIds []int64)([]*StageExec, error)  {
	var execs []*StageExec
	err := x.Where(builder.Eq{"act_id": iterationActionId}.And(builder.In("stage_id", stageIds))).Find(&execs)
	if err != nil{
		return nil, err
	}
	return execs, nil
}

func InsertStageExec(exec *StageExec) (int64, error){
	id, err := x.Insert(exec)
	return id, err
}

func UpdateStageExecState(id int64, state string) error {
	_, err := x.Table(&StageExec{}).ID(id).Update(map[string]interface{}{"state": state})
	return err
}

// (actionsIds[0],stageIds[0]), (actionsIds[1],stageIds[1])
func BatchQueryStateExecByActIdAndStageId(actionIds, stageIds []int64) ([]*StageExec, error) {
	var execs []*StageExec
	var filterResult []*StageExec
	condMap := make(map[string]bool)
	for i:=0; i < len(actionIds); i++ {
		condMap[fmt.Sprintf("%d-%d", actionIds[i], stageIds[i])] = true
	}

	err := x.Cols("id","act_id","stage_id", "state").Where(builder.In("act_id", actionIds)).And(builder.In("stage_id", stageIds)).Find(&execs)
	if err != nil{
		return nil, err
	}
	for _,v := range execs {
		key := fmt.Sprintf("%d-%d", v.ActId, v.StageId)
		if condMap[key] {
			filterResult = append(filterResult, v)
		}
	}
	return filterResult, nil
}

func GetStageExecIdByActIdAndStageId(actionId, stageId int64) (int64, error) {
	stageExec := new(StageExec)
	exist, err := x.Table("stage_exec").Where(builder.Eq{"act_id":actionId, "stage_id":stageId}).Limit(1).Get(stageExec)
	if !exist || err != nil {
		return 0, err
	}
	return stageExec.Id, nil
}