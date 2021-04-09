package env

import (
	"encoding/json"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/grpc/invoke/invokerImpl"
)

//1. const info for each step in each iteration
var ieah = &iterEnvActionHolder{m: map[string][]iterEnvAction{"dev": {{"完成开发阶段"}, {"提交MR"}, {"Jar包管理"}, {"配置变更"}, {"触发pipeline"}, {"申请服务器"}, {"新建联调环境"}},
	"itg":       {{"完成集成阶段"}, {"提交MR"}, {"Jar包管理"}, {"触发pipeline"}},
	"pre":       {{"完成预发阶段"}, {"提交MR"}, {"Jar包管理"}, {"触发pipeline"}},
	"grayscale": {{"完成灰度阶段"}, {"配置白名单"}, {"配置黑名单"}, {"流量控制"}},
	"prod":      {{"完成发布"}},
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



//3. status of an iteration
func IterInfo(c *context.Context) []byte {
	iterationId := c.ParamsInt64(":iterationId")
	iteration, _ := db.GetIterationById(iterationId)
	info := [][]string{
		{"开发阶段", "", "finish"},
		{"集成阶段", "", "finish"},
		{"预发阶段", "", "process"},
		{"灰度发布", "", "wait"},
		{"发布阶段", "", "wait"},
	}
	for i := 0; i<iteration.IterState;i++ {
		info[i][2] = "finish"
	}
	info[iteration.IterState][2] = "process"
	for i := iteration.IterState + 1; i < 5; i++ {
		info[i][2] = "wait"
	}
	data, _ := json.Marshal(info)
	return data

}


