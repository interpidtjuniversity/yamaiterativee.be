package stage

import (
	"encoding/json"
	"fmt"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/iteration/step"
)


type Node struct {
	ActionIdStageId string     `json:"actionId_stageId"`
	StageId         int64      `json:"stageId"`
	ActionId        int64      `json:"actionId"`
	Id              int64      `json:"id"`
	Label           string     `json:"label"`
	ClassName       string     `json:"className"`
	IconType        string     `json:"iconType"`
	Top             int        `json:"top"`
	Left            int        `json:"left"`
	Group           string     `json:"group"`
	Endpoints       []Endpoint `json:"endpoints"`
}
//
type Endpoint struct {
	Id          string     `json:"id"`
	Orientation []int     `json:"orientation"`
	Pos         []float64 `json:"pos"`
}

func (ep *Endpoint)FormatId(id int){
	if ep.Id!="" {
		ep.Id = fmt.Sprintf(ep.Id, id)
	}
}

func IterStageInfo(c *context.Context) ([]byte, error){
	stageId := c.ParamsInt64(":stageId")

	var infoSteps []step.InfoStep
	stage,_ := db.GetStageById(stageId)
	steps,_ := db.BranchQueryStepsByIds(stage.Steps)
	for i:=0; i<len(steps); i++ {
		infoStep := step.InfoStep{Index: i, Image: steps[i].Img, Title: steps[i].Name, StepId: steps[i].ID}
		infoSteps = append(infoSteps, infoStep)
	}

	data, err := json.Marshal(infoSteps)
	return data, err
}
