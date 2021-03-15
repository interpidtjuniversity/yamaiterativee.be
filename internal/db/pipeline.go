package db

type PipeLine struct {
	ID int64
	Name string
	CreatorId int64
	CreatorName string
	IsPublic bool
	Stages []int64
	StageDAG [][]int64
}