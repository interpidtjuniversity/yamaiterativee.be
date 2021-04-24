package db

import (
	"fmt"
)

type Iteration struct {
	ID              int64    `xorm:"id autoincr pk"`
	IterCreatorId   string   `xorm:"iter_creator_uid"`
	IterType        string   `xorm:"iter_type"`
	IterAdmin       []string `xorm:"iter_admin"`
	IterState       int      `xorm:"iter_state"`        // 0,1,2,3,4   -> dev,itg,pre,gary.prod
	IterBranch      string   `xorm:"iter_branch"`
	IterDevActGroup int64    `xorm:"iter_dev_act_group"`
	IterPreActGroup int64    `xorm:"iter_pre_act_group"`
	IterItgActGroup int64    `xorm:"iter_itg_act_group"`
	OwnerName       string   `xorm:"owner_name""`
	RepoName        string   `xorm:"repo_name"`
	IterDevClc      float64  `xorm:"iter_dev_clc"`
	IterItgClc      float64  `xorm:"iter_itg_clc"`
	IterPreClc      float64  `xorm:"iter_pre_clc"`
	IterDevQs       float64  `xorm:"iter_dev_qs"`
	IterItgQs       float64  `xorm:"iter_itg_qs" `
	IterPreQs       float64  `xorm:"iter_pre_qs"`
	DevPr           int      `xorm:"dev_pr"`
	ItgPr           int      `xorm:"itg_pr"`
	PrePr           int      `xorm:"pre_pr"`
}

func GetIterationById(id int64) (*Iteration, error){
	iteration := &Iteration{}
	has, _ := x.ID(id).Get(iteration)
	if !has {
		return nil, ErrIterationNotExist{Args: map[string]interface{}{"iterationId": id}}
	}
	return iteration, nil
}

func InsertIteration(iteration Iteration) (int64, error){
	_,err := x.Insert(&iteration)
	if err != nil {
		return 0, err
	}
	return iteration.ID, nil
}

func UpdateIteration(iteration Iteration) error {
	_,err := x.ID(iteration.ID).Update(iteration)
	return err
}


type ErrIterationNotExist struct {
	Args map[string]interface{}
}

func (err ErrIterationNotExist) Error() string {
	return fmt.Sprintf("iteration does not exist: %v", err.Args)
}

func (ErrIterationNotExist) NotFound() bool {
	return true
}