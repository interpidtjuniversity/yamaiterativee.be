package db

import (
	"fmt"
	"xorm.io/builder"
)


type IterationState int

const (
	INIT_STATE = -1
	DEV_STATE IterationState = iota
	ITG_STATE
	PRE_STATE
	GARY_STATE
	PROD_STATE
)

func (is IterationState) ToString() string {
	switch is {
	case INIT_STATE:
		return "init"
	case DEV_STATE:
		return "dev"
	case ITG_STATE:
		return "itg"
	case PRE_STATE:
		return "pre"
	case GARY_STATE:
		return "gary"
	case PROD_STATE:
		return "prod"
	default:
		return "unknown"
	}
}

type Iteration struct {
	ID              int64          `xorm:"id autoincr pk"`
	IterCreator     string         `xorm:"iter_creator_uid"`
	IterType        string         `xorm:"iter_type"`
	IterAdmin       []string       `xorm:"iter_admin"`
	IterState       IterationState `xorm:"iter_state"` // 0,1,2,3,4   -> dev,itg,pre,gary.prod
	IterBranch      string         `xorm:"iter_branch"`
	IterDevActGroup int64          `xorm:"iter_dev_act_group"`
	IterPreActGroup int64          `xorm:"iter_pre_act_group"`
	IterItgActGroup int64          `xorm:"iter_itg_act_group"`
	OwnerName       string         `xorm:"owner_name"`
	RepoName        string         `xorm:"repo_name"`
	IterDevClc      float64        `xorm:"iter_dev_clc"`
	IterItgClc      float64        `xorm:"iter_itg_clc"`
	IterPreClc      float64        `xorm:"iter_pre_clc"`
	IterDevQs       float64        `xorm:"iter_dev_qs"`
	IterItgQs       float64        `xorm:"iter_itg_qs" `
	IterPreQs       float64        `xorm:"iter_pre_qs"`
	DevPr           int            `xorm:"dev_pr"`
	ItgPr           int            `xorm:"itg_pr"`
	PrePr           int            `xorm:"pre_pr"`
	BaseCommit      string         `xorm:"base_commit"`
	Title           string         `xorm:"title"`
	Content         string         `xorm:"content"`
	TestConfig      string         `xorm:"test_config"`
	ProdConfig      string         `xorm:"prod_config"`
	DevConfig       string         `xorm:"dev_config"`
	PreConfig       string         `xorm:"pre_config"`
	StableConfig    string         `xorm:"stable_config"`
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

func GetIterationByAdmin(adminId string) []*Iteration {
	var allIterations []*Iteration
	var filterIterIds []int64
	var filterIters []*Iteration
	x.Table("iteration").Cols("id","iter_admin").Find(&allIterations)
	for _, iter := range allIterations {
		for _, admin := range iter.IterAdmin {
			if admin == adminId {
				filterIterIds = append(filterIterIds, iter.ID)
				break
			}
		}
	}
	x.Table("iteration").Where(builder.In("id", filterIterIds)).Find(&filterIters)
	return filterIters
}

func GetIterationConfigByIterId(iterId int64) (*Iteration, error) {
	iteration := new(Iteration)
	exist, err := x.Table("iteration").Cols("dev_config", "stable_config", "test_config", "pre_config", "prod_config").Where(builder.Eq{"id": iterId}).Limit(1).Get(iteration)
	if err!=nil || !exist {
		return nil, err
	}
	return iteration,nil
}

func GetIterationDevConfig(iterId int64) string {
	iteration := new(Iteration)
	exist, err := x.Table("iteration").Cols("dev_config").Where(builder.Eq{"id": iterId}).Limit(1).Get(iteration)
	if err!=nil || !exist {
		return ""
	}
	return iteration.DevConfig
}

func GetIterationStableConfig(iterId int64) string {
	iteration := new(Iteration)
	exist, err := x.Table("iteration").Cols("stable_config").Where(builder.Eq{"id": iterId}).Limit(1).Get(iteration)
	if err!=nil || !exist {
		return ""
	}
	return iteration.StableConfig
}

func GetIterationTestConfig(iterId int64) string {
	iteration := new(Iteration)
	exist, err := x.Table("iteration").Cols("test_config").Where(builder.Eq{"id": iterId}).Limit(1).Get(iteration)
	if err!=nil || !exist {
		return ""
	}
	return iteration.TestConfig
}

func GetIterationPreConfig(iterId int64) string {
	iteration := new(Iteration)
	exist, err := x.Table("iteration").Cols("pre_config").Where(builder.Eq{"id": iterId}).Limit(1).Get(iteration)
	if err!=nil || !exist {
		return ""
	}
	return iteration.PreConfig
}

func GetIterationProdConfig(iterId int64) string {
	iteration := new(Iteration)
	exist, err := x.Table("iteration").Cols("prod_config").Where(builder.Eq{"id": iterId}).Limit(1).Get(iteration)
	if err!=nil || !exist {
		return ""
	}
	return iteration.ProdConfig
}

func UpdateIterationConfig(iterId int64, iteration *Iteration) (bool, error) {
	_, err := x.Table("iteration").Where(builder.Eq{"id": iterId}).Limit(1).Update(iteration)
	if err!=nil {
		return false, err
	}
	return true, nil
}

func GetIterationAllAdmins(iterId int64) (*Iteration, error) {
	iteration := new(Iteration)
	exist, err := x.Table("iteration").Cols("iter_admin").Where(builder.Eq{"id": iterId}).Limit(1).Get(iteration)
	if !exist && err!=nil {
		return nil, err
	}
	return iteration, err
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