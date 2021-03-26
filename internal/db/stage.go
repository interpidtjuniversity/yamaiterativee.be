package db

import "xorm.io/builder"

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
