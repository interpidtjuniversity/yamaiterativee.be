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
	has, err := x.Where(builder.Eq{"iter_id":iterationId, "group_type":envType}).Get(iag)
	if !has {
		return nil, err
	}
	return iag, nil
}

func InsertIterationActGroup(group *IterationActGroup) (int64, error){
	_,err := x.Insert(group)
	if err != nil {
		return 0, err
	}
	return group.ID, nil
}

func GetOrGenerateIterationActGroup(iterationId int64, envType string) (int64, error) {
	var iags []*IterationActGroup
	_ = x.Table("iteration_act_group").Where(builder.Eq{"iter_id": iterationId}).Find(&iags)
	var developIag *IterationActGroup
	for _, iag := range iags {
		if iag.GroupType == "dev" {
			developIag = iag
		}
		if iag.GroupType == envType {
			return iag.ID, nil
		}
	}
	if envType == "pre" {
		preIag := &IterationActGroup{
			GroupType: "pre",
			IterationId: iterationId,
			TargetBranch: "master",
		}
		InsertIterationActGroup(preIag)
		return preIag.ID, nil
	} else if envType == "itg" && developIag != nil {
		itgIag := &IterationActGroup{
			GroupType:    "itg",
			IterationId:  iterationId,
			TargetBranch: developIag.TargetBranch,
		}
		InsertIterationActGroup(itgIag)
		return itgIag.ID, nil
	}
	return 0,nil
}