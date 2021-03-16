package db

type Pipeline struct {
	ID          int64     `xorm:"id autoincr pk"`
	Name        string    `xorm:"name"`
	CreatorId   int64     `xorm:"creator_id"`
	CreatorName string    `xorm:"creator_name"`
	IsPublic    bool      `xorm:"is_public"`
	Stages      []int64   `xorm:"stages" json:"-"`
	StageDAG    [][]int64 `xorm:"stage_dag" json:"-"`
}