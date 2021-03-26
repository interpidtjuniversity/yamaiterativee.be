package pipeline

import (
	"encoding/json"
	"fmt"
	"testing"
	"yama.io/yamaIterativeE/internal/db"
)

func Test_PipelineParse(t *testing.T) {
	pipeline := &db.Pipeline{ID: 1, Name: "经典MR", CreatorName: "interpidtjuniversity", IsPublic: true, Stages: []int64{1,2,3,4,5,6,7},
		StageDAG: [][]int64{
		{0,0,0,0,1,0,0},
		{0,0,1,0,0,0,0},
		{0,0,0,1,0,0,0},
		{0,0,0,0,1,0,0},
		{0,0,0,0,0,1,0},
		{0,0,0,0,0,0,1},
		{0,0,0,0,0,0,0},
		},
	StageLayout: [][]int64{
		{1,0,0,0,0,0,0},
		{2,3,4,5,6,7,0},
		{0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0},
	}}
	endpoints,_ := ParseEndpoints(pipeline)
	fmt.Print(json.Marshal(endpoints))
}
