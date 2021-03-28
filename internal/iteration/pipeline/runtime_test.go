package pipeline

import (
	"fmt"
	"testing"
)

func Test_DAGParallel(t *testing.T) {
	array := [][]int64{
		{0,0,0,0,1,0,0},
		{0,0,1,0,0,0,0},
		{0,0,0,1,0,0,0},
		{0,0,0,0,1,0,0},
		{0,0,0,0,0,1,0},
		{0,0,0,0,0,0,1},
		{0,0,0,0,0,0,0},
	}

	array = [][]int64{
		{0,0,1,0,0},
		{0,0,1,0,0},
		{0,0,0,1,0},
		{0,0,0,0,0},
		{0,0,1,0,0},
	}

	array = [][]int64{
		{0,1,0},
		{0,0,1},
		{0,0,0},
	}
	fmt.Print(GetDAGMaxParallel(array))
}
