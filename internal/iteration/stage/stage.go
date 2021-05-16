package stage

import (
	"fmt"
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
	State           string     `json:"state"`
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
