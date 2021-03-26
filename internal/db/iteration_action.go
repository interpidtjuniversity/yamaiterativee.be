package db

import (
	"xorm.io/builder"
)

type IterationAction struct {
	ID          int64  `xorm:"id autoincr pk"`
	ActorName   string `xorm:"actor_name"`
	FinallyPass bool   `xorm:"finally_pass" json:"-"`
	PipeLineId  int64  `xorm:"pipeline_id"`
	EnvGroup    int64  `xorm:"env_group"`
	State       string `xorm:"state"`
	ActionInfo  string `xorm:"action_info"`
	AvatarSrc   string `xorm:"avatar_src"`
	ExtInfo     string `xorm:"ext_info"`
}

func GetIterActionByActGroup(actGroupId int64) ([]*IterationAction, error) {
	var actions []*IterationAction
	err := x.Where(builder.Eq{"env_group": actGroupId}).Find(&actions)
	if err!=nil {
		return nil, err
	}
	return actions, nil
}

type IterationMergeRequest struct {
	IterationAction
	SponsorPassId []int64  `xorm:"sponsor_pass_id" json:"-"`
	SponsorID     int64    `xorm:"sponsor_id"`
	SponsorName   string   `xorm:"sponsor_name"`
	ReviewersID   []int64  `xorm:"reviewers_id" json:"-"`
	ReviewersName []string `xorm:"reviewers_name" json:"-"`
}

type IterationServerApply struct {
	IterationAction
}

type IterationResourceRelease struct {
	IterationAction
}

