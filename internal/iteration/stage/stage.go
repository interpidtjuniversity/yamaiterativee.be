package stage

import (
	"encoding/json"
	"fmt"
	"sync"
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

var p *ping = &ping{counter: 0, m: map[int64]string{0: "loading", 1:"warning", 2:"success", 3:"error", 4:"ellipsis"}, colorM: map[int64]string{0: "#1F49E0", 1:"#FFA003", 2:"#1DC11D", 3:"#FF3333"}}
type ping struct {
	mu sync.Mutex
	counter int64
	m map[int64]string
	colorM map[int64]string
}

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

func StageExecInfo(c *context.Context) {
	//stageId := c.ParamsInt64(":stage")
	//execId := c.ParamsInt64(":exec")
}

type stepStatus struct {
	Type   string `json:"type"`
	Color  string `json:"color"`
	Status string `json:"status"`
}


func Ping(c *context.Context) []byte{
	p.mu.Lock()
	msg := stepStatus{Type: p.m[int64(int(p.counter)%len(p.m))], Color: p.colorM[int64(int(p.counter)%len(p.colorM))]}
	p.counter++
	p.mu.Unlock()
	data, _ := json.Marshal(msg)
	return data
}

var ieah = &iterEnvActionHolder{m: map[string][]iterEnvAction{"dev": {{"完成开发阶段"}, {"提交MR"}, {"Jar包管理"}, {"配置变更"}, {"触发pipeline"}, {"申请服务器"}, {"新建联调环境"}},
                                                            "itg": {{ "完成集成阶段"}, { "提交MR"}, {"Jar包管理"}, {"触发pipeline"}},
                                                            "pre": {{ "完成预发阶段"}, { "提交MR"}, {"Jar包管理"}, {"触发pipeline"}},
                                                            "grayscale":{{ "完成灰度阶段"}, {"配置白名单"}, {"配置黑名单"}, {"流量控制"}},
                                                            "prod": {{"完成发布"}},
                                                            },
}
type iterEnvActionHolder struct {
	m map[string][]iterEnvAction
}
type iterEnvAction struct {
	ButtonShowWords string `json:"buttonShowWords"`
}

func IterActionInfo(c *context.Context) []byte {
	envType := c.ParamsEscape(":envType")
	iterEnvActions := ieah.m[envType]
	data, _ := json.Marshal(iterEnvActions)

	return data
}


var iei = &iterEnvInfoHolder{m: map[string]iterEnvInfo{
	"dev": {TargetBranch: "E14621321654894_20210319", LatestCommit: "ce7894s", LatestMode: "MR", ServiceChange: "小", PRCount: 100, QualityScore: 90, ChangeLineCoverage: "90%"},
	"pre": {TargetBranch: "master", LatestCommit: "b7s855", LatestMode: "MR", ServiceChange: "小", PRCount: 1000, QualityScore: 95, ChangeLineCoverage: "95%"},
}}

type iterEnvInfoHolder struct {
	m map[string]iterEnvInfo
}

type iterEnvInfo struct {
	TargetBranch       string          `json:"targetBranch"`
	LatestMode         string          `json:"latestMode"`
	LatestCommit       string          `json:"latestCommit"`
	ServiceChange      string          `json:"serviceChange"`
	PRCount            int             `json:"PRCount"`
	QualityScore       int             `json:"qualityScore"`
	ChangeLineCoverage string          `json:"changeLineCoverage"`
}

func IterEnvInfo (c *context.Context) []byte {
	envType := c.ParamsEscape(":envType")
	info := iei.m[envType]
	data, _ := json.Marshal(info)
	return data
}