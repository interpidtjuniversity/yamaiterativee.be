package db

import (
	"fmt"
	"testing"
)

func Test_CreateIteration(t *testing.T) {
	NewEngine()

	iter := Iteration{
		IterCreatorId: "interpidtjuniversity",
		IterAdmin: []string{"interpidtjuniversity"},
		IterState: []string{"dev"},
		IterBranch: "E9876543210_20210226",
		Application: "test",
		IterType: "basic MR",
	}
	_, err := x.Insert(iter)
	if err != nil {
		fmt.Print(err)
	}

}
