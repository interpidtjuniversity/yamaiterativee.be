package db

import (
	"fmt"
	"testing"
)

func Test_InsertActGroup(t *testing.T) {
	NewEngine()
	var actGroups []IterationActGroup
	actGroups = append(actGroups,IterationActGroup{GroupType: "dev", IterationId: 1})
	actGroups = append(actGroups,IterationActGroup{GroupType: "itg", IterationId: 1})
	actGroups = append(actGroups,IterationActGroup{GroupType: "pre", IterationId: 1})

	for _, group := range actGroups {
		if _, err := x.Insert(group); err != nil {
			fmt.Print(err)
		}
	}
}

