package db

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateIteration(t *testing.T) {
	NewEngine()

	iter := Iteration{
		IterCreator: "interpidtjuniversity",
		IterAdmin:   []string{"interpidtjuniversity"},
		IterState:   3,
		IterBranch:  "E9876543210_20210226",
		OwnerName:   "interpidtjuniversity",
		RepoName:    "test",
		IterType:    "basic MR",
	}
	_, err := x.Insert(iter)
	if err != nil {
		fmt.Print(err)
	}

}

func Test_GetGrayOrder(t *testing.T) {
	NewEngine()

	iteration, err := GetIterationById(18)
	assert.Nil(t, err)
	fmt.Print(iteration.GrayOrder)
}
