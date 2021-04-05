package db

import "xorm.io/builder"

type IterationActGroup struct {
	ID           int64  `xorm:"id autoincr pk"`
	GroupType    string `xorm:"group_type"`
	IterationId  int64  `xorm:"iter_id"`
	TargetBranch string `xorm:"target_branch"`
}

func GetIterationActGroupByIterationIdAndEnv(iterationId int64, envType string) (*IterationActGroup, error){
	iag := &IterationActGroup{}
	has, err := x.Where(builder.Eq{"group_type":envType}.And(builder.Eq{"iter_id":iterationId})).Get(iag)
	if !has {
		return nil, err
	}
	return iag, nil
}