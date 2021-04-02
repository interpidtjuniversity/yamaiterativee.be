package db

import "xorm.io/builder"

type Step struct {
	ID          int64    `xorm:"id autoincr pk"`
	Name        string   `xorm:"name"`
	CreatorId   int64    `xorm:"creator_id"`
	CreatorName string   `xorm:"creator_name"`
	Requires    []string `xorm:"requires" json:"-"`
	Command     string   `xorm:"command"`
	Args        []string `xorm:"args" json:"-"`
	IsPublic    bool     `xorm:"is_public"`
	Img         string   `xorm:"img"`
}

func BranchQueryStepsByIds(stageId []int64) ([]*Step, error) {
	var steps []*Step
	err := x.Where(builder.In("id", stageId)).Find(&steps)
	if err != nil {
		return nil, err
	}
	return steps, nil
}