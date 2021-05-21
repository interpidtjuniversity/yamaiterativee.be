package env

import (
	"encoding/json"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/grpc/invoke/invokerImpl"
)

//1. const info for each step in each iteration
var ieah = &iterEnvActionHolder{m: map[string][]iterEnvAction{"dev": {{ButtonShowWords:"完成开发阶段", ID: "finishDev"}, {ButtonShowWords:"提交MR", ID: "submitMRDev"}, {ButtonShowWords:"Jar包管理", ID: "jarManageDev"}, {ButtonShowWords:"配置变更",ID: "changeConfigDev"}, {ButtonShowWords:"触发pipeline",ID: "triggerPipelineDev"}, {ButtonShowWords:"申请服务器",ID: "applyServerDev"}, {ButtonShowWords:"新建联调环境", ID: "newJointDebuggingDev"}},
	"itg":       {{ButtonShowWords:"完成集成阶段", ID: "finishItg"}, {ButtonShowWords:"提交MR", ID: "submitMRItg"}, {ButtonShowWords:"Jar包管理",ID: "jarManageItg"}, {ButtonShowWords:"触发pipeline",ID: "triggerPipelineItg"}},
	"pre":       {{ButtonShowWords:"完成预发阶段",ID: "finishPre"}, {ButtonShowWords:"提交MR",ID: "submitMRPre"}, {ButtonShowWords:"Jar包管理",ID: "jarManagePre"}, {ButtonShowWords:"触发pipeline",ID: "triggerPipelinePre"}},
	"grayscale": {{ButtonShowWords:"完成灰度阶段", ID: "finishGray"}, {ButtonShowWords:"配置白名单", ID: "whiteList"}, {ButtonShowWords:"配置黑名单", ID: "blackList"}, {ButtonShowWords:"流量控制", ID: "flowControl"}},
	"prod":      {{ButtonShowWords:"完成发布", ID: "finishProd"}},
	},
}
type iterEnvActionHolder struct {
	m map[string][]iterEnvAction
}
type iterEnvAction struct {
	ButtonShowWords string `json:"buttonShowWords"`
	ID              string `json:"id"`
}

func IterActionInfo(c *context.Context) []byte {
	envType := c.ParamsEscape(":envType")
	iterEnvActions := ieah.m[envType]
	data, _ := json.Marshal(iterEnvActions)

	return data
}

//2. mate data of an iteration

type iterEnvInfo struct {
	TargetBranch       string          `json:"targetBranch"`
	LatestMode         string          `json:"latestMode"`
	LatestCommit       string          `json:"latestCommit"`
	LatestCommitLink   string          `json:"latestCommitLink"`
	ServiceChange      string          `json:"serviceChange"`
	PRCount            int             `json:"PRCount"`
	QualityScore       int             `json:"qualityScore"`
	ChangeLineCoverage string          `json:"changeLineCoverage"`
}

func IterEnvInfo (c *context.Context) []byte {
	iterationId := c.ParamsInt64(":iterationId")
	envType := c.ParamsEscape(":envType")
	iteration, _ := db.GetIterationById(iterationId)
	actGroup,_ := db.GetIterationActGroupByIterationIdAndEnv(iterationId, envType)

	if actGroup == nil{
		return nil
	}
	latestCommit, latestCommitLink := invokerImpl.InvokeQueryRepoBranchCommitService(iteration.OwnerName, iteration.RepoName, actGroup.TargetBranch)

	var pr int
	var qs float64
	var clc float64
	switch envType {
	case "dev":
		 pr = iteration.DevPr
		 qs = iteration.IterDevQs
		 clc = iteration.IterDevClc
		 break
	case "itg":
		pr = iteration.ItgPr
		qs = iteration.IterItgQs
		clc = iteration.IterItgClc
		break
	case "pre":
		pr = iteration.PrePr
		qs = iteration.IterPreQs
		clc = iteration.IterPreClc
		break
	}
	data, _ := json.Marshal(iterEnvInfo{TargetBranch: actGroup.TargetBranch, LatestCommit: latestCommit, LatestCommitLink: latestCommitLink, PRCount: pr, ChangeLineCoverage: string(rune(int(clc * 100))), QualityScore: int(qs), LatestMode: "MR"})
	return data
}


type IterationInfo struct {
	StateArray  [][]string `json:"stateArray"`
	Owner       string     `json:"owner"`
	Application string     `json:"application"`
	IterBranch  string     `json:"iterBranch"`
	IterTitle   string     `json:"iterTitle"`
	IterState   int        `json:"iterState"`
	ServerType  string     `json:"serverType"`
}
//3. status of an iteration
func IterInfo(c *context.Context) []byte {
	iterationId := c.ParamsInt64(":iterationId")
	info := IterationInfo{}
	iteration, _ := db.GetIterationById(iterationId)
	stateArray := [][]string{
		{"开发阶段", "", "wait"},
		{"集成阶段", "", "wait"},
		{"预发阶段", "", "wait"},
		{"灰度发布", "", "wait"},
		{"发布阶段", "", "wait"},
	}
	for i := 0; i<int(iteration.IterState);i++ {
		stateArray[i][2] = "finish"
	}
	if iteration.IterState >= 0 {
		stateArray[iteration.IterState][2] = "process"
	}
	for i := iteration.IterState + 1; i < 5; i++ {
		stateArray[i][2] = "wait"
	}
	info.StateArray = stateArray
	info.Owner = iteration.OwnerName
	info.Application = iteration.RepoName
	info.IterBranch = iteration.IterBranch
	info.IterTitle = iteration.Title
	info.IterState = int(iteration.IterState)
	info.ServerType = db.GetApplicationTypeByOwnerAndRepo(iteration.OwnerName, iteration.RepoName)

	data, _ := json.Marshal(info)
	return data

}


