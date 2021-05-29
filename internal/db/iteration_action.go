package db

import (
	"xorm.io/builder"
)

type IterationAction struct {
	Id              int64  `xorm:"id autoincr pk"`
	ActorName       string `xorm:"actor_name"`
	FinallyPass     bool   `xorm:"finally_pass" json:"-"`
	PipeLineId      int64  `xorm:"pipeline_id"`
	EnvGroup        int64  `xorm:"env_group"`
	State           string `xorm:"state"`
	ActionInfo      string `xorm:"action_info"`
	AvatarSrc       string `xorm:"avatar_src"`
	ExtInfo         string `xorm:"ext_info"`
	ActionGroupInfo string `xorm:"action_group_info"`
	ExecPath        string `xorm:"exec_path"`
}

func GetIterActionByActGroup(actGroupId int64) ([]*IterationAction, error) {
	var actions []*IterationAction
	err := x.Where(builder.Eq{"env_group": actGroupId}).Find(&actions)
	if err!=nil {
		return nil, err
	}
	return actions, nil
}

func LightlyGetIterActionByActGroup(actGroupId int64) ([]*IterationAction, error) {
	var actions []*IterationAction
	err := x.Cols("id", "actor_name", "finally_pass", "pipeline_id").Where(builder.Eq{"env_group": actGroupId}).Find(&actions)
	if err!=nil {
		return nil, err
	}
	return actions, nil
}

func GetIterActionById (actionId int64) (*IterationAction, error) {
	iterationAction := &IterationAction{}
	has, err := x.ID(actionId).Get(iterationAction)
	if !has {
		return nil, err
	}
	return iterationAction, nil
}

func InsertIterationAction(action *IterationAction) (int64, error) {
	id, err := x.Insert(action)
	return id, err
}

func UpdateIterationAction(id int64, state string) error {
	_, err := x.Table(&IterationAction{}).ID(id).Update(map[string]interface{}{"state": state})
	return err
}

func GetIterActionByUserName(userName string, limit int) ([]*IterationAction, error){
	var actions []*IterationAction
	err := x.Table("iteration_action").Where(builder.Eq{"actor_name": userName}).Desc("id").Limit(limit).Find(&actions)
	return actions, err
}

func GetIterActionByUserNameAndPipelineId(userName string, pipelineId int64, limit int) ([]*IterationAction, error){
	var actions []*IterationAction
	err := x.Table("iteration_action").Where(builder.Eq{"actor_name": userName, "pipeline_id": pipelineId}).Desc("id").Limit(limit).Find(&actions)
	return actions, err
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

