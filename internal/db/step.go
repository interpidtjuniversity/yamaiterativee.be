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
	Type        string   `xorm:"type"`
	LogType     int      `xorm:"log_type"`
}

func BranchQueryStepsByIds(stepId []int64) ([]*Step, error) {
	var steps []*Step
	err := x.Where(builder.In("id", stepId)).Find(&steps)
	if err != nil {
		return nil, err
	}
	return steps, nil
}

func GetStepLogTypeById(stepId int64) int {
	step := new(Step)
	exist, err := x.Table("step").Cols("log_type").Where(builder.Eq{"id": stepId}).Get(step)
	if !exist || err != nil {
		return -1
	}
	return step.LogType

}