package env

import (
	"encoding/json"
	"strconv"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/grpc/invoke/invokerImpl"
	"yama.io/yamaIterativeE/internal/iteration/step/log"
)

//1. const info for each step in each iteration
var ieah = &iterEnvActionHolder{m: map[string][]iterEnvAction{"dev": {{ButtonShowWords:"完成开发阶段", ID: "finishDev", Type: 0}, {ButtonShowWords:"提交MR", ID: "submitMRDev", Type: 0}, {ButtonShowWords:"Jar包管理", ID: "jarManageDev", Type: 0}, {ButtonShowWords:"配置变更",ID: "changeConfigDev", Type: 0}, {ButtonShowWords:"触发pipeline",ID: "triggerPipelineDev", Type: 0}, {ButtonShowWords:"申请服务器",ID: "applyServerDev", Type: 0}, {ButtonShowWords:"联调环境", ID: "jointDebuggingDev", Type: 0}},
	"itg":       {{ButtonShowWords:"完成集成阶段", ID: "finishItg", Type: 1}, {ButtonShowWords:"提交MR", ID: "submitMRItg", Type: 1}, {ButtonShowWords:"Jar包管理",ID: "jarManageItg", Type: 1}, {ButtonShowWords:"触发pipeline",ID: "triggerPipelineItg", Type: 1}, {ButtonShowWords:"同步master代码",ID: "syncMaster", Type: 1}},
	"pre":       {{ButtonShowWords:"完成预发阶段",ID: "finishPre", Type: 2}, {ButtonShowWords:"提交MR",ID: "submitMRPre", Type: 2}, {ButtonShowWords:"Jar包管理",ID: "jarManagePre", Type: 2}, {ButtonShowWords:"触发pipeline",ID: "triggerPipelinePre", Type: 2}},
	"grayscale": {{ButtonShowWords:"完成灰度阶段", ID: "finishGray", Type: 3}, {ButtonShowWords:"配置白名单", ID: "whiteList", Type: 3}, {ButtonShowWords:"配置黑名单", ID: "blackList", Type: 3}, {ButtonShowWords:"推进", ID: "advanceGray", Type: 3}, {ButtonShowWords:"回滚", ID: "rollBackGray", Type: 3}},
	"prod":      {{ButtonShowWords:"完成发布", ID: "finishProd", Type: 4}},
	"finish":    {{ButtonShowWords:"发布完成", ID: "prodFinish", Type: 4}},
	},
}
type iterEnvActionHolder struct {
	m map[string][]iterEnvAction
}
type iterEnvAction struct {
	ButtonShowWords string `json:"buttonShowWords"`
	ID              string `json:"id"`
	Type            int    `json:"type"`
}

func IterActionInfo(c *context.Context) []byte {
	envType := c.ParamsEscape(":envType")
	iterEnvActions := ieah.m[envType]
	data, _ := json.Marshal(iterEnvActions)

	return data
}

//2. mate data of an iteration

type iterEnvInfo struct {
	TargetBranch       string `json:"targetBranch"`
	LatestMode         string `json:"latestMode"`
	LatestCommit       string `json:"latestCommit"`
	LatestCommitLink   string `json:"latestCommitLink"`
	ServiceChange      string `json:"serviceChange"`
	PRCount            int    `json:"PRCount"`
	QualityScore       int    `json:"qualityScore"`
	ChangeLineCoverage string `json:"changeLineCoverage"`
	Type               string `json:"type"`
	GrayPercent        string `json:"grayPercent"`
	AdvanceGrayState   bool   `json:"advanceGrayState"`
	RollBackGrayState  bool   `json:"rollBackGrayState"`
}

func IterEnvInfo (c *context.Context) []byte {
	iterationId := c.ParamsInt64(":iterationId")
	envType := c.ParamsEscape(":envType")
	iteration, _ := db.GetIterationById(iterationId)

	if envType == "dev" || envType == "itg" || envType == "pre" {
		actGroup, _ := db.GetIterationActGroupByIterationIdAndEnv(iterationId, envType)

		if actGroup == nil {
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
		data, _ := json.Marshal(iterEnvInfo{Type: envType, TargetBranch: actGroup.TargetBranch, LatestCommit: latestCommit, LatestCommitLink: latestCommitLink, PRCount: pr, ChangeLineCoverage: strconv.Itoa(int(clc * 100)), QualityScore: int(qs), LatestMode: "MR"})
		return data
	} else if envType == "grayscale" || envType == "prod"{
		data, _ := json.Marshal(iterEnvInfo{Type: envType, GrayPercent: iteration.GrayPercent, RollBackGrayState: iteration.GrayRollBackState, AdvanceGrayState: iteration.GrayAdvanceState})
		return data
	}
	return nil
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
		{"发布完成", "", "wait"},
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


func SyncInfo(c *context.Context) ([]byte, error) {
	iterationId := c.ParamsInt64(":iterationId")
	envType := c.ParamsEscape(":envType")
	syncKey := c.Query("syncKey")
	appName := c.Query("appName")

	var actionIds []int64
	group, _ := db.GetIterationActGroupByIterationIdAndEnv(iterationId, envType)
	actions, _ := db.LightlyGetIterActionByActGroup(group.ID)
	for _, action := range actions {
		actionIds = append(actionIds, action.Id)
	}
	stageExec, _ := db.GetIterationEnvLatestSuccessStageExec(actionIds, []int64{7,18})
	report := log.ConstructTestReport(stageExec.ExecPath, appName)
	reportLog := log.Log{}
	json.Unmarshal(report, &reportLog)

	var result = 0.0
	if reportLog.Data != nil {
		var partIndex, allIndex = 0.0, 0.0
		if syncKey == "qualityScore" {
			for _, v := range (reportLog.Data).(map[string]interface{}) {
				for _, vI := range v.([]interface{}) {
					vv := vI.(map[string]interface{})
					//var part = vv.BranchCovered + vv.LineCovered + vv.InstructionCovered + vv.ComplexityCovered + vv.MethodCovered
					var part = vv["branchCovered"].(float64) + vv["complexityCovered"].(float64) + vv["instructionCovered"].(float64) + vv["lineCovered"].(float64) + vv["methodCovered"].(float64)
					var all = part + vv["branchMissed"].(float64) + vv["complexityMissed"].(float64) + vv["instructionMissed"].(float64) + vv["lineMissed"].(float64) + vv["methodMissed"].(float64)
					//var all = part + vv.MethodMissed + vv.LineMissed + vv.InstructionMissed + vv.ComplexityMissed + vv.BranchMissed
					partIndex += part
					allIndex += all
				}
			}
			result = partIndex / allIndex * 100
			//result = float64(partIndex) / float64(allIndex) *100
			db.UpdateIterationEnvQualityScore(iterationId, envType, result)
		} else if syncKey == "lineCoverage" {
			for _, v := range (reportLog.Data).(map[string]interface{}) {
				for _, vI := range v.([]interface{}) {
					vv := vI.(map[string]interface{})
					var part = vv["lineCovered"].(float64)
					var all = part + vv["lineMissed"].(float64)
					//var part = vv.LineCovered
					//var all  = part + vv.LineMissed
					partIndex += part
					allIndex += all
				}
			}
			result = partIndex / allIndex
			//result = float64(partIndex) / float64(allIndex)
			db.UpdateIterationEnvLineCoverage(iterationId, envType, result)
			result *= 100
		}
	}
	data, _ := json.Marshal(int(result))
	return data, nil
}


