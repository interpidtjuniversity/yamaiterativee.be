package db

import (
	"encoding/json"
	"fmt"
	"yama.io/yamaIterativeE/internal/context"
)

type Iteration struct {
	ID              int64    `xorm:"id autoincr pk"`
	IterCreatorId   string   `xorm:"iter_creator_uid"`
	IterType        string   `xorm:"iter_type"`
	IterAdmin       []string `xorm:"iter_admin"`
	IterState       []string `xorm:"iter_state"`
	IterBranch      string   `xorm:"iter_branch"`
	IterDevActGroup int64    `xorm:"iter_dev_act_group"`
	IterPreActGroup int64    `xorm:"iter_pre_act_group"`
	IterItgActGroup int64    `xorm:"iter_itg_act_group"`
	Application     string   `xorm:"application"`
	IterDevClc      float64  `xorm:"iter_dev_clc"`
	IterItgClc      float64  `xorm:"iter_itg_clc"`
	IterPreClc      float64  `xorm:"iter_pre_clc"`
	IterDevQs       float64  `xorm:"iter_dev_qs"`
	IterItgQs       float64  `xorm:"iter_itg_qs" `
	IterPreQs       float64  `xorm:"iter_pre_qs"`
	DevPr           []int    `xorm:"dev_pr"`
	ItgPr           []int    `xorm:"itg_pr"`
	PrePr           []int    `xorm:"pre_pr"`
}

// return iteration status, which is 'process'
func IterInfo(c *context.Context) []byte {
	iterationId := c.ParamsInt64(":iterationId")
	if iterationId == 1 {
		info := [][]string{
			{"开发阶段", "", "finish"},
			{"集成阶段", "", "finish"},
			{"预发阶段", "", "process"},
			{"灰度发布", "", "wait"},
			{"发布阶段", "", "wait"},
		}
		data, _ := json.Marshal(info)
		return data
	}
	return nil

}

func GetIterationById(id int64) (*Iteration, error){
	iteration := &Iteration{}
	has, _ := x.ID(id).Get(iteration)
	if !has {
		return nil, ErrIterationNotExist{Args: map[string]interface{}{"iterationId": id}}
	}
	return iteration, nil
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