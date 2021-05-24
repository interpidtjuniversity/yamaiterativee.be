package db

import (
	"fmt"
	"xorm.io/builder"
)


type IterationState int

const (
	INIT_STATE    IterationState = -1
	DEV_STATE     IterationState = 0
	ITG_STATE     IterationState = 1
	PRE_STATE     IterationState = 2
	GRAY_STATE    IterationState = 3
	PROD_STATE    IterationState = 4
	FINISH_STATE  IterationState = 5
	UNKNOWN_STATE IterationState = 6
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
	case GRAY_STATE:
		return "gray"
	case PROD_STATE:
		return "prod"
	default:
		return "unknown"
	}
}

func (is IterationState)FromString(env string) IterationState {
	switch env {
	case "init":
		return INIT_STATE
	case "dev":
		return DEV_STATE
	case "itg":
		return ITG_STATE
	case "pre":
		return PRE_STATE
	case "gray":
		return GRAY_STATE
	case "prod":
		return PROD_STATE
	default:
		return UNKNOWN_STATE
	}
}

func (is IterationState) NextState() IterationState {
	switch is {
	case INIT_STATE:
		return DEV_STATE
	case DEV_STATE:
		return ITG_STATE
	case ITG_STATE:
		return PRE_STATE
	case PRE_STATE:
		return GRAY_STATE
	case GRAY_STATE:
		return PROD_STATE
	case PROD_STATE:
		return FINISH_STATE
	default:
		return UNKNOWN_STATE
	}
}

type Iteration struct {
	ID                int64          `xorm:"id autoincr pk"`
	IterCreator       string         `xorm:"iter_creator_uid"`
	IterType          string         `xorm:"iter_type"`
	IterAdmin         []string       `xorm:"iter_admin"`
	IterState         IterationState `xorm:"iter_state"` // 0,1,2,3,4   -> dev,itg,pre,gray.prod
	IterBranch        string         `xorm:"iter_branch"`
	IterDevActGroup   int64          `xorm:"iter_dev_act_group"`
	IterPreActGroup   int64          `xorm:"iter_pre_act_group"`
	IterItgActGroup   int64          `xorm:"iter_itg_act_group"`
	OwnerName         string         `xorm:"owner_name"`
	RepoName          string         `xorm:"repo_name"`
	IterDevClc        float64        `xorm:"iter_dev_clc"`
	IterItgClc        float64        `xorm:"iter_itg_clc"`
	IterPreClc        float64        `xorm:"iter_pre_clc"`
	IterDevQs         float64        `xorm:"iter_dev_qs"`
	IterItgQs         float64        `xorm:"iter_itg_qs" `
	IterPreQs         float64        `xorm:"iter_pre_qs"`
	DevPr             int            `xorm:"dev_pr"`
	ItgPr             int            `xorm:"itg_pr"`
	PrePr             int            `xorm:"pre_pr"`
	BaseCommit        string         `xorm:"base_commit"`
	Title             string         `xorm:"title"`
	Content           string         `xorm:"content"`
	TestConfig        string         `xorm:"test_config"`
	ProdConfig        string         `xorm:"prod_config"`
	DevConfig         string         `xorm:"dev_config"`
	PreConfig         string         `xorm:"pre_config"`
	StableConfig      string         `xorm:"stable_config"`
	GrayPercent       string         `xorm:"gray_percent"`
	GrayOrder         []string       `xorm:"gray_order"`
	GrayAdvanceState  bool           `xorm:"gray_advance_state"`
	GrayRollBackState bool           `xorm:"gray_rollback_state"`
}

func GetIterationById(id int64) (*Iteration, error){
	iteration := &Iteration{}
	exist, err := x.Table("iteration").Cols("id", "iter_creator_uid", "iter_type", "iter_admin","iter_state",
		"iter_branch","iter_dev_act_group","iter_pre_act_group","iter_itg_act_group","owner_name","repo_name","iter_dev_clc",
		"iter_itg_clc","iter_pre_clc","iter_dev_qs","iter_itg_qs","iter_pre_qs","dev_pr","itg_pr","pre_pr","base_commit",
		"title","content","gray_percent","gray_order","gray_advance_state","gray_rollback_state").Where(builder.Eq{"id":id}).Get(iteration)
	if !exist {
		return nil, err
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

func UpdateIterationState(iterId int64, s string) error {
	now := new(Iteration)
	state := UNKNOWN_STATE.FromString(s)
	exist, err := x.Table("iteration").Cols("iter_state").Where(builder.Eq{"id":iterId}).Get(now)
	if exist {
		if now.IterState < state {
			now.IterState = state
			_, err = x.Table("iteration").Cols("iter_state").Where(builder.Eq{"id":iterId}).Update(now)
		}
	}
	return err
}

func UpdateIterationDevActGroup(iterId, actGroupId int64) error {
	iteration := Iteration{IterDevActGroup: actGroupId}
	_, err := x.Table("iteration").Cols("iter_dev_act_group").Where(builder.Eq{"id":iterId}).Update(&iteration)
	return err
}
func UpdateIterationItgActGroup(iterId, actGroupId int64) error{
	iteration := Iteration{IterItgActGroup: actGroupId}
	_, err := x.Table("iteration").Cols("iter_itg_act_group").Where(builder.Eq{"id":iterId}).Update(&iteration)
	return err
}
func UpdateIterationPreActGroup(iterId, actGroupId int64) error{
	iteration := Iteration{IterPreActGroup: actGroupId}
	_, err := x.Table("iteration").Cols("iter_pre_act_group").Where(builder.Eq{"id":iterId}).Update(&iteration)
	return err
}

func GetIterationGrayPercent(iterId int64) string {
	iteration := Iteration{}
	exist, err := x.Table("iteration").Cols("gray_percent").Where(builder.Eq{"id":iterId}).Get(&iteration)
	if !exist || err!=nil {
		return ""
	}
	return iteration.GrayPercent
}

func UpdateIterationAdvanceGrayInfo(iterId int64, grayOrder []string, grayPercent string, advanceState bool) error {
	iteration := Iteration{GrayOrder: grayOrder, GrayPercent: grayPercent, GrayAdvanceState: advanceState}
	_, err := x.Table("iteration").Cols("gray_order", "gray_percent","gray_advance_state").Where(builder.Eq{"id":iterId}).Update(&iteration)
	return err
}

func UpdateIterationAdvanceState(iterId int64, advanceState bool) error {
	iteration := Iteration{GrayAdvanceState: advanceState}
	_, err := x.Table("iteration").Cols("gray_advance_state").Where(builder.Eq{"id":iterId}).Update(&iteration)
	return err
}

func UpdateIterationRollBackGrayInfo(iterId int64, grayOrder []string, grayPercent string, rollBackState bool) error {
	iteration := Iteration{GrayOrder: grayOrder, GrayPercent: grayPercent, GrayRollBackState: rollBackState}
	_, err := x.Table("iteration").Cols("gray_order", "gray_percent","gray_rollback_state").Where(builder.Eq{"id":iterId}).Update(&iteration)
	return err
}

func UpdateIterationRollBackState(iterId int64, rollBackState bool) error {
	iteration := Iteration{GrayRollBackState: rollBackState}
	_, err := x.Table("iteration").Cols("gray_rollback_state").Where(builder.Eq{"id":iterId}).Update(&iteration)
	return err
}

func GetIterationAdvanceGrayState(iterId int64) (string, bool) {
	iteration := Iteration{}
	exist, err := x.Table("iteration").Cols("gray_percent", "gray_advance_state").Where(builder.Eq{"id":iterId}).Get(&iteration)
	if !exist || err != nil {
		return "0", false
	}
	return iteration.GrayPercent, iteration.GrayAdvanceState
}

func GetIterationRollBackGrayState(iterId int64) (string, bool) {
	iteration := Iteration{}
	exist, err := x.Table("iteration").Cols("gray_percent", "gray_rollback_state").Where(builder.Eq{"id":iterId}).Get(&iteration)
	if !exist || err != nil {
		return "0", false
	}
	return iteration.GrayPercent, iteration.GrayRollBackState
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