package db

import "xorm.io/builder"

type Pipeline struct {
	ID          int64     `xorm:"id autoincr pk"`
	Name        string    `xorm:"name"`
	CreatorId   int64     `xorm:"creator_id"`
	CreatorName string    `xorm:"creator_name"`
	IsPublic    bool      `xorm:"is_public"`
	Stages      []int64   `xorm:"stages"`
	StageDAG    [][]int64 `xorm:"stage_dag"`       //edge info
	StageLayout [][]int64 `xorm:"stage_layout"`    //node position info
}

func BranchQueryPipelineByIds(ids []int64) ([]*Pipeline, error) {
	var pipelines []*Pipeline
	err := x.Where(builder.In("id", ids)).Find(&pipelines)
	if err!=nil {
		return nil, err
	}
	return pipelines, nil
}

func GetPipelineById(id int64) (*Pipeline, error) {
	pipeline := &Pipeline{}
	has, err := x.ID(id).Get(pipeline)
	if !has {
		return nil, err
	}
	return pipeline, nil
}
