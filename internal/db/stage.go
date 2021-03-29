package db

import (
	"fmt"
	"xorm.io/builder"
)

type Stage struct {
	ID          int64   `xorm:"id autoincr pk"`
	Name        string  `xorm:"name"`
	CreatorId   int64   `xorm:"creator_id"`
	CreatorName string  `xorm:"creator_name"`
	IsPublic    bool    `xorm:"is_public"`
	Steps       []int64 `xorm:"steps" json:"-"`
	ClassName   string  `xorm:"class_name"`
	IconType    string  `xorm:"icon_type"`
	Group       string  `xorm:"group"`
}

func BranchQueryStage(stageIds []int64)([]*Stage,error) {
	var stages []*Stage
	err := x.Where(builder.In("id", stageIds)).Find(&stages)
	if err != nil {
		return nil, err
	}
	return stages,nil
}

func GetStageById(id int64)(*Stage,error) {
	stage := &Stage{}
	has, _ := x.ID(id).Get(stage)
	if !has {
		return nil, ErrIterationNotExist{Args: map[string]interface{}{"stageId":id}}
	}
	return stage, nil
}

type ErrStageNotExist struct {
	Args map[string]interface{}
}

func (err ErrStageNotExist) Error() string {
	return fmt.Sprintf("stage does not exist: %v", err.Args)
}

func (ErrStageNotExist) NotFound() bool {
	return true
}