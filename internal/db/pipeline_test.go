package db

import (
	"fmt"
	"testing"
)

func Test_PipelineInsert(t *testing.T) {
	NewEngine()

	var pipelines []Pipeline
	pipelines = append(pipelines,Pipeline{Name: "经典MR",CreatorId: 0, CreatorName: "interpidtjuniversity", IsPublic: true, Stages: []int64{1,2,3,4,5,6,7}, StageDAG: [][]int64{
		{0,0,0,0,1,0,0},
		{0,0,1,0,0,0,0},
		{0,0,0,1,0,0,0},
		{0,0,0,0,1,0,0},
		{0,0,0,0,0,1,0},
		{0,0,0,0,0,0,1},
		{0,0,0,0,0,0,0},
	}})

	for _, pipeline := range pipelines {
		if _, err := x.Insert(pipeline); err != nil {
			fmt.Print(err)
		}
	}
}
