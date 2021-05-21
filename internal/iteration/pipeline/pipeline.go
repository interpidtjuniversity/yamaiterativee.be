package pipeline

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/grpc/invoke/invokerImpl"
	"yama.io/yamaIterativeE/internal/iteration/stage"
	"yama.io/yamaIterativeE/internal/iteration/step"
	"yama.io/yamaIterativeE/internal/util"
)

var PIPELINE_EXEC_PATH = "/root/yamaIterativeE/yamaIterativeE-pipeline-exec/%s"
var STAGE_EXEC_PATH = "/root/yamaIterativeE/yamaIterativeE-pipeline-exec/%s/%s"
var PMD_SCAN_PATH = "%s/%s"

var EndpointType = []stage.Endpoint{
	{Id: "%d_left", Orientation: []int{-1, 0}, Pos: []float64{0, 0.5}},
	{Id: "%d_right", Orientation: []int{1, 0}, Pos: []float64{0, 0.5}},
	{Id: "%d_top", Orientation: []int{0, -1}, Pos: []float64{0.5, 0}},
	{Id: "%d_bottom", Orientation: []int{0, 1}, Pos: []float64{0.5, 0}},
}

var EdgeType = [][]edge{
	{
		edge{Source: "%d_left", Target: "%d_left", Arrow: true, Type: "endpoint", ArrowPosition:0.5, ShapeType: "Flow"},
		edge{Source: "%d_left", Target: "%d_right", Arrow: true, Type: "endpoint", ArrowPosition:0.5, ShapeType: "Flow"},
		edge{Source: "%d_left", Target: "%d_top", Arrow: true, Type: "endpoint", ArrowPosition:0.5, ShapeType: "Flow"},
		edge{Source: "%d_left", Target: "%d_bottom", Arrow: true, Type: "endpoint", ArrowPosition:0.5, ShapeType: "Flow"},
	},
	{
		edge{Source: "%d_right", Target: "%d_left", Arrow: true, Type: "endpoint", ArrowPosition:0.5, ShapeType: "Flow"},
		edge{Source: "%d_right", Target: "%d_right", Arrow: true, Type: "endpoint", ArrowPosition:0.5, ShapeType: "Flow"},
		edge{Source: "%d_right", Target: "%d_top", Arrow: true, Type: "endpoint", ArrowPosition:0.5, ShapeType: "Flow"},
		edge{Source: "%d_right", Target: "%d_bottom", Arrow: true, Type: "endpoint", ArrowPosition:0.5, ShapeType: "Flow"},
	},
	{
		edge{Source: "%d_top", Target: "%d_left", Arrow: true, Type: "endpoint", ArrowPosition:0.5, ShapeType: "Flow"},
		edge{Source: "%d_top", Target: "%d_right", Arrow: true, Type: "endpoint", ArrowPosition:0.5, ShapeType: "Flow"},
		edge{Source: "%d_top", Target: "%d_top", Arrow: true, Type: "endpoint", ArrowPosition:0.5, ShapeType: "Flow"},
		edge{Source: "%d_bottom", Target: "%d_bottom", Arrow: true, Type: "endpoint", ArrowPosition:0.5, ShapeType: "Flow"},
	},
	{
		edge{Source: "%d_bottom", Target: "%d_left", Arrow: true, Type: "endpoint", ArrowPosition:0.5, ShapeType: "Flow"},
		edge{Source: "%d_bottom", Target: "%d_right", Arrow: true, Type: "endpoint", ArrowPosition:0.5, ShapeType: "Flow"},
		edge{Source: "%d_bottom", Target: "%d_top", Arrow: true, Type: "endpoint", ArrowPosition:0.5, ShapeType: "Flow"},
		edge{Source: "%d_bottom", Target: "%d_bottom", Arrow: true, Type: "endpoint", ArrowPosition:0.5, ShapeType: "Flow"},
	},
}

// reset these value to chang the stage position of pipeline
const (
	FirstStagePX = 55
	FirstStagePY = 50
	StageXStep   = 70
	StageYStep   = 175

	GroupTop     = 0
	GroupLeft    = 130

	FirstGroupPY = 320
	FirstGroupPX = 75
	GroupYStep   = 125
	GroupXStep   = 75
)

type EndpointPosition int
const (
	LEFT  EndpointPosition = iota
	RIGHT
	TOP
	BOTTOM
)

func (ep EndpointPosition)String(id int) string{
	switch ep {
	case LEFT:
		return fmt.Sprintf("%d_left", id)
	case RIGHT:
		return fmt.Sprintf("%d_right", id)
	case TOP:
		return fmt.Sprintf("%d_top", id)
	case BOTTOM:
		return fmt.Sprintf("%d_bottom", id)
	}
	return "N/A"
}

type pipelineInfo struct {
	PipelineId  int64        `json:"pipelineId"`
	AvatarSrc   string       `json:"avatarSrc"`
	ActionInfo  string       `json:"actionInfo"`
	ExtInfo     string       `json:"extInfo"`
	Nodes       []stage.Node `json:"nodes"`
	Edges       []edge       `json:"edges"`
	Groups      []group      `json:"groups"`
	State       string       `json:"state"`
	IterationId int64        `json:"iterationId"`
}

func (pi pipelineInfo)Clone() pipelineInfo  {
	clone := pipelineInfo{}
	clone.PipelineId = pi.PipelineId
	clone.AvatarSrc = pi.AvatarSrc
	clone.ActionInfo = pi.ActionInfo
	clone.State = pi.State
	clone.IterationId = pi.IterationId
	clone.Nodes = make([]stage.Node, len(pi.Nodes))
	clone.Edges = make([]edge, len(pi.Edges))
	clone.Groups = make([]group, len(pi.Groups))

	for i:=0; i<len(pi.Nodes); i++ {
		clone.Nodes[i].StageId = pi.Nodes[i].StageId
		clone.Nodes[i].Id = pi.Nodes[i].Id
		clone.Nodes[i].ActionId = pi.Nodes[i].ActionId
		clone.Nodes[i].Group = pi.Nodes[i].Group
		clone.Nodes[i].Label = pi.Nodes[i].Label
		clone.Nodes[i].ClassName = pi.Nodes[i].ClassName
		clone.Nodes[i].IconType = pi.Nodes[i].IconType
		clone.Nodes[i].ActionIdStageId = pi.Nodes[i].ActionIdStageId
		clone.Nodes[i].Left = pi.Nodes[i].Left
		clone.Nodes[i].Top = pi.Nodes[i].Top
		clone.Nodes[i].Endpoints = make([]stage.Endpoint, len(pi.Nodes[i].Endpoints))
		for j := 0; j < len(pi.Nodes[i].Endpoints); j++ {
			clone.Nodes[i].Endpoints[j].Id = pi.Nodes[i].Endpoints[j].Id
			clone.Nodes[i].Endpoints[j].Pos = pi.Nodes[i].Endpoints[j].Pos
			clone.Nodes[i].Endpoints[j].Orientation = pi.Nodes[i].Endpoints[j].Orientation
		}
	}
	for i:=0; i<len(pi.Edges); i++ {
		clone.Edges[i].TargetNode = pi.Edges[i].TargetNode
		clone.Edges[i].Target = pi.Edges[i].Target
		clone.Edges[i].SourceNode = pi.Edges[i].SourceNode
		clone.Edges[i].Source = pi.Edges[i].Source
		clone.Edges[i].ShapeType = pi.Edges[i].ShapeType
		clone.Edges[i].Type = pi.Edges[i].Type
		clone.Edges[i].Arrow = pi.Edges[i].Arrow
		clone.Edges[i].ArrowPosition = pi.Edges[i].ArrowPosition
	}
	for i:=0; i<len(pi.Groups); i++ {
		clone.Groups[i].Id = pi.Groups[i].Id
		clone.Groups[i].Top = pi.Groups[i].Top
		clone.Groups[i].Left = pi.Groups[i].Left
		clone.Groups[i].Height = pi.Groups[i].Height
		clone.Groups[i].Width = pi.Groups[i].Width
		clone.Groups[i].Resize = pi.Groups[i].Resize
		clone.Groups[i].Draggable = pi.Groups[i].Draggable
		clone.Groups[i].Options = groupOptions{Title: pi.Groups[i].Options.Title}
	}
	return clone
}

type edge struct {
	Source        string  `json:"source"`
	Target        string  `json:"target"`
	SourceNode    int   `json:"sourceNode"`
	TargetNode    int   `json:"targetNode"`
	Arrow         bool    `json:"arrow"`
	Type          string  `json:"type"`
	ArrowPosition float64 `json:"arrowPosition"`
	ShapeType     string  `json:"shapeType"`
}

func (e *edge) Format(startIndex, endIndex int) {
	if e.Source!=""{
		e.Source = fmt.Sprintf(e.Source, startIndex)
	}
	if e.Target!=""{
		e.Target = fmt.Sprintf(e.Target, endIndex)
	}
	e.SourceNode = startIndex
	e.TargetNode = endIndex
}

type group struct {
	Id        string       `json:"id"`
	Draggable bool         `json:"draggable"`
	Top       int          `json:"top"`
	Left      int          `json:"left"`
	Width     int          `json:"width"`
	Height    int          `json:"height"`
	Resize    bool         `json:"resize"`
	Options   groupOptions `json:"options"`
}

type groupOptions struct {
	Title string `json:"title"`
}

func StartNewServerPipelineInternal(c *context.Context) ([]byte, error) {
	pipelineId := c.ParamsInt64(":pipelineId")
	owner := c.Query("owner")
	iterId := c.QueryInt64("iterId")
	env := c.Query("env")
	appType := c.Query("appType")
	appOwner := c.Query("appOwner")
	appName := c.Query("appName")
	serverType := c.Query("serverType")
	actionInfo := fmt.Sprintf("%s 申请了服务器", owner)
	avatarSrc,_ := db.GetUserAvatarByUserName(owner)
	envGroup, _ := db.GetOrGenerateIterationActGroup(iterId, env)

	// reg pipeline
	pipeExec := db.IterationAction{ActorName: owner, PipeLineId: pipelineId, EnvGroup: envGroup, State: Init.ToString(),
		ActionInfo: actionInfo, AvatarSrc: avatarSrc, ExecPath: fmt.Sprintf(PIPELINE_EXEC_PATH, util.GenerateRandomStringWithSuffix(10,"")),
	}

	// prepare workspace
	os.MkdirAll(pipeExec.ExecPath, os.ModePerm)
	_, _ = db.InsertIterationAction(&pipeExec)
	pipeline,_ := db.GetPipelineById(pipelineId)

	runtimePipeline := FromIterationAction(pipeExec, *pipeline, &map[string]interface{}{
		"appOwner":appOwner,
		"appName":appName,
		"execPath": pipeExec.ExecPath,
		"appType": appType,
		"serverType": serverType,
		"env":env,
		"iterId": strconv.Itoa(int(iterId)),
	})
	_ = e.Reg(runtimePipeline)
	return nil, nil
}

func StartDeployPipelineInternal(c *context.Context) ([]byte, error) {
	pipelineId := c.ParamsInt64(":pipelineId")
	actorName := c.Query("actorName")
	iterId := c.QueryInt64("iterId")
	branchName := c.Query("branchName")
	env := c.Query("env")
	actionInfo := fmt.Sprintf("%s 手动触发pipeline", actorName)
	appOwner := c.Query("appOwner")
	appName := c.Query("appName")
	serverName := c.Query("serverName")
	repoURL := db.GetApplicationRepoURLByOwnerAndRepo(appOwner, appName)
	serverIP := db.GetServerIPByServerName(serverName)
	avatarSrc,_ := db.GetUserAvatarByUserName(actorName)
	envGroup, _ := db.GetOrGenerateIterationActGroup(iterId, env)

	// reg pipeline
	pipeExec := db.IterationAction{ActorName: actorName, PipeLineId: pipelineId, EnvGroup: envGroup, State: Init.ToString(),
		ActionInfo: actionInfo, AvatarSrc: avatarSrc, ExecPath: fmt.Sprintf(PIPELINE_EXEC_PATH, util.GenerateRandomStringWithSuffix(10,"")),
	}
	// prepare workspace
	os.MkdirAll(pipeExec.ExecPath, os.ModePerm)
	_, _ = db.InsertIterationAction(&pipeExec)
	pipeline,_ := db.GetPipelineById(pipelineId)

	runtimePipeline := FromIterationAction(pipeExec, *pipeline, &map[string]interface{}{
		"appOwner":appOwner,
		"appName":appName,
		"repoURL":repoURL,
		"branchName": branchName,
		"execPath": pipeExec.ExecPath,
		"env":env,
		"serverName": serverName,
		"serverIP": serverIP,
		"iterId": strconv.Itoa(int(iterId)),
		"appPath":appName,
		"pmdScanPath":fmt.Sprintf(PMD_SCAN_PATH, appName, "src"),
	})
	_ = e.Reg(runtimePipeline)
	return nil, nil
}

func StartBasicMRPipeline(c *context.Context) ([]byte, error) {
	// get pipeline
	pipelineId := c.ParamsInt64(":pipelineId")
	actorName := c.Query("actorName")
	envGroup := c.QueryInt64("envGroup")
	actionInfo := c.Query("actionInfo")
	avatarSrc := c.Query("avatarSrc")
	extInfo := c.Query("extInfo")
	actionGroupInfo := c.Query("actionGroupInfo")


	pipeExec := db.IterationAction{ActorName: actorName, PipeLineId: pipelineId, EnvGroup: envGroup, State: Init.ToString(),
		ActionGroupInfo: actionGroupInfo, ActionInfo: actionInfo, AvatarSrc: avatarSrc, ExtInfo: extInfo, ExecPath: "/root/yamaIterativeE/yamaIterativeE-pipeline-exec",
	}
	_, _ = db.InsertIterationAction(&pipeExec)

	pipeline,_ := db.GetPipelineById(pipelineId)

	runtimePipeline := FromIterationAction(pipeExec, *pipeline, &map[string]interface{}{})
	err := e.Reg(runtimePipeline)

	return nil,err
}

func ReStartBasicMRPipelineWithArgs(pipelineId, iterId, actionId int64, actorName, iterDevelopBranch, iterTargetBranch string,
	env, actionInfo, appOwner, appName string, mrCodeReviews []string) error {
	actionGroupInfo := fmt.Sprintf("%s -> %s", iterDevelopBranch, iterTargetBranch)
	avatarSrc,_ := db.GetUserAvatarByUserName(actorName)
	envGroup, _ := db.GetOrGenerateIterationActGroup(iterId, env)
	db.UpdateIterationState(iterId, env)
	repoURL := db.GetApplicationRepoURLByOwnerAndRepo(appOwner, appName)

	// reg pipeline
	pipeExec := db.IterationAction{ActorName: actorName, PipeLineId: pipelineId, EnvGroup: envGroup, State: Init.ToString(),
		ActionGroupInfo: actionGroupInfo, ActionInfo: actionInfo, AvatarSrc: avatarSrc, ExecPath: fmt.Sprintf(PIPELINE_EXEC_PATH, util.GenerateRandomStringWithSuffix(10,"")),
	}
	// prepare workspace
	os.MkdirAll(pipeExec.ExecPath, os.ModePerm)
	_, _ = db.InsertIterationAction(&pipeExec)
	pipeline,_ := db.GetPipelineById(pipelineId)

	// reg merge request code review
	_, mergeRequestCodeReviewUrl := invokerImpl.InvokeRegisterMergeRequestService(appOwner, appName, iterDevelopBranch, iterTargetBranch,
		pipeExec.Id, 1, 11, mrCodeReviews, pipelineId, actorName, iterId, env, actionInfo)
	mergeReq := MergeRequest{UserName: appOwner, Repository: appName, SourceBranch: iterDevelopBranch, TargetBranch: iterTargetBranch, MergeInfo: actionInfo}
	mergeReqData, _ := json.Marshal(mergeReq)
	runtimePipeline := FromIterationAction(pipeExec, *pipeline, &map[string]interface{}{
		"mergeRequestCodeReviewUrl":mergeRequestCodeReviewUrl,
		"sourceBranch":iterDevelopBranch,
		"targetBranch":iterTargetBranch,
		"appOwner":appOwner,
		"appName":appName,
		"repoURL":repoURL,
		"appPath":appName,
		"pmdScanPath":fmt.Sprintf(PMD_SCAN_PATH, appName, "src"),
		"mergeArg":fmt.Sprintf("'%s'",string(mergeReqData)),
		"mergeInfo":actionInfo,
		"yamaHubAddr":"localhost:8000",
		"mergeService":"proto.YaMaHubBranchService/Merge2Branch",
	})
	e.UnReg(actionId)
	_ = e.Reg(runtimePipeline)

	return nil
}

func StartBasicMRPipelineInternal(c *context.Context) ([]byte, error) {
	// form data
	pipelineId := c.ParamsInt64(":pipelineId")
	actorName := c.Query("actorName")
	iterId := c.QueryInt64("iterId")
	iterTargetBranch := c.Query("iterTargetBranch")
	iterDevelopBranch := c.Query("iterDevelopBranch")
	var mrCodeReviews []string
	json.Unmarshal([]byte(c.Query("mrCodeReviews")), &mrCodeReviews)
	env := c.Query("env")
	actionInfo := c.Query("mrInfo")
	appOwner := c.Query("appOwner")
	appName := c.Query("appName")

	// auto data
	actionGroupInfo := fmt.Sprintf("%s -> %s", iterDevelopBranch, iterTargetBranch)
	avatarSrc,_ := db.GetUserAvatarByUserName(actorName)
	envGroup, _ := db.GetOrGenerateIterationActGroup(iterId, env)
	db.UpdateIterationState(iterId, env)
	repoURL := db.GetApplicationRepoURLByOwnerAndRepo(appOwner, appName)

	// reg pipeline
	pipeExec := db.IterationAction{ActorName: actorName, PipeLineId: pipelineId, EnvGroup: envGroup, State: Init.ToString(),
		ActionGroupInfo: actionGroupInfo, ActionInfo: actionInfo, AvatarSrc: avatarSrc, ExecPath: fmt.Sprintf(PIPELINE_EXEC_PATH, util.GenerateRandomStringWithSuffix(10,"")),
	}
	// prepare workspace
	os.MkdirAll(pipeExec.ExecPath, os.ModePerm)
	_, _ = db.InsertIterationAction(&pipeExec)
	pipeline,_ := db.GetPipelineById(pipelineId)

	// reg merge request code review
	_, mergeRequestCodeReviewUrl := invokerImpl.InvokeRegisterMergeRequestService(appOwner, appName, iterDevelopBranch, iterTargetBranch,
		pipeExec.Id, 1, 11, mrCodeReviews, pipelineId, actorName, iterId, env, actionInfo)
	mergeReq := MergeRequest{UserName: appOwner, Repository: appName, SourceBranch: iterDevelopBranch, TargetBranch: iterTargetBranch, MergeInfo: actionInfo}
	mergeReqData, _ := json.Marshal(mergeReq)
	runtimePipeline := FromIterationAction(pipeExec, *pipeline, &map[string]interface{}{
		"mergeRequestCodeReviewUrl":mergeRequestCodeReviewUrl,
		"sourceBranch":iterDevelopBranch,
		"targetBranch":iterTargetBranch,
		"appOwner":appOwner,
		"appName":appName,
		"repoURL":repoURL,
		"appPath":appName,
		"pmdScanPath":fmt.Sprintf(PMD_SCAN_PATH, appName, "src"),
		"mergeArg":fmt.Sprintf("'%s'",string(mergeReqData)),
		"mergeInfo":actionInfo,
		"yamaHubAddr":"localhost:8000",
		"mergeService":"proto.YaMaHubBranchService/Merge2Branch",
	})
	_ = e.Reg(runtimePipeline)

	return nil,nil
}

func CancelPipeline(c *context.Context) []byte{
	actionId := c.ParamsInt64(":actionId")
	e.UnReg(actionId)
	return []byte("success")
}

// callBack for other system
func PassStep(actionId, stageId, stepId int64) []byte {

	for action:=e.actions.Front(); action!=nil; action = action.Next() {
		runtimePipeline, _ := action.Value.(*RuntimePipeline)
		if runtimePipeline.ID == actionId {
			for _, se := range runtimePipeline.Buckets {
				if se.StageId == stageId {
					for _, sp := range se.Steps {
						if sp.StepId == stepId {
							sp.Cond = true
							return []byte("success")
						}
					}
				}
			}
		}
	}
	return []byte("success")
}

func IterPipelineInfo(c *context.Context) ([]byte,error) {
	// check param
	iterationId := c.ParamsInt64(":iterationId")
	envType := c.ParamsEscape(":envType")

	// get all pipeline from this iterationId in envType env
	iteration,err := db.GetIterationById(iterationId)
	if err!=nil {
		return nil, err
	}

	var envGroup int64
	switch envType {
	case "dev":
		envGroup = iteration.IterDevActGroup
		break
	case "itg":
		envGroup = iteration.IterItgActGroup
		break
	case "pre":
		envGroup = iteration.IterPreActGroup
		break
	}

	envActions, err := db.GetIterActionByActGroup(envGroup)
	pipelineInfoMap := make(map[int64]pipelineInfo)
	stream, _ := util.New(envActions)
	var pipelineIds []int64
	_ = stream.Map(func(action *db.IterationAction) int64{
		return action.PipeLineId
	}).ToSlice(&pipelineIds)
	pipelines, _ := db.BranchQueryPipelineByIds(pipelineIds)
	// agg build
	for _,p := range pipelines {
		// 1.parse endpoint and edge
		endpoints,edges,_ := ParseDagAndLayout(p)
		// 2.parse node
		nodes,_ := makeNode(p, endpoints)
		// 3.create group
		groups,_ := makeGroups(p)
		// 4.init pipelineInfo template, no point so we can deep copy
		pli := pipelineInfo{Nodes: nodes, Edges: edges, Groups: groups}
		// 5.save result
		pipelineInfoMap[p.ID] = pli
	}
	// agg actId and stageId
	var actionIds []int64
	var stageIds []int64
	filterStateMap := make(map[string]string)
	stream.ForEach(func(action *db.IterationAction) {
		pipelineInfoTemplate := pipelineInfoMap[action.PipeLineId]
		for i := 0; i<len(pipelineInfoTemplate.Nodes); i++ {
			actionIds = append(actionIds, action.Id)
			stageIds = append(stageIds, pipelineInfoTemplate.Nodes[i].StageId)
		}
	})
	filterResults, _ := db.BatchQueryStateExecByActIdAndStageId(actionIds, stageIds)
	for _, filterResult := range filterResults {
		filterStateMap[fmt.Sprintf("%d-%d", filterResult.ActId, filterResult.StageId)] = filterResult.State
	}

	var pipelineInfos []pipelineInfo
	stageIdsMap := make(map[string][]*db.Stage)
	// create completely pipelineInfo with pipeline template and envActions
	stream.ForEach(func(action *db.IterationAction) {
		pipelineInfoTemplate := pipelineInfoMap[action.PipeLineId].Clone()
		// call stageInfo with actId and stageId
		var stageIds []int64
		for _,node := range pipelineInfoTemplate.Nodes {
			stageIds = append(stageIds, node.StageId)
		}
		// agg StageExec and Stage
		// stageExecs,_ := db.BranchQueryStageExec(action.Id, stageIds)
		var stages []*db.Stage
		key := concatInt64Slice(stageIds)
		if stageIdsMap[key] == nil {
			stages, _ = db.BranchQueryStage(stageIds)
			stageIdsMap[key] = stages
		} else{
			stages = stageIdsMap[key]
		}
		stageMap := make(map[int64]*db.Stage)
		for _,s :=range stages {
			stageMap[s.ID] = s
		}
		for i := 0; i<len(pipelineInfoTemplate.Nodes); i++ {
			pipelineInfoTemplate.Nodes[i].IconType = stageMap[pipelineInfoTemplate.Nodes[i].StageId].IconType
			pipelineInfoTemplate.Nodes[i].ClassName = stageMap[pipelineInfoTemplate.Nodes[i].StageId].ClassName
			pipelineInfoTemplate.Nodes[i].Group = stageMap[pipelineInfoTemplate.Nodes[i].StageId].Group
			pipelineInfoTemplate.Nodes[i].Label = stageMap[pipelineInfoTemplate.Nodes[i].StageId].Name
			pipelineInfoTemplate.Nodes[i].ActionId = (*action).Id
			pipelineInfoTemplate.Nodes[i].ActionIdStageId = fmt.Sprintf("%d_%d", (*action).Id, pipelineInfoTemplate.Nodes[i].StageId)
			pipelineInfoTemplate.Nodes[i].State = filterStateMap[fmt.Sprintf("%d-%d", pipelineInfoTemplate.Nodes[i].ActionId, pipelineInfoTemplate.Nodes[i].StageId)]
		}
		pipelineInfoTemplate.PipelineId=action.Id
		pipelineInfoTemplate.ActionInfo=action.ActionInfo
		pipelineInfoTemplate.AvatarSrc=action.AvatarSrc
		pipelineInfoTemplate.ExtInfo=action.ExtInfo
		pipelineInfoTemplate.State = action.State
		pipelineInfoTemplate.IterationId = iterationId
		// only one group exist so we get index of 0
		pipelineInfoTemplate.Groups[0].Options.Title = action.ActionGroupInfo

		pipelineInfos = append(pipelineInfos, pipelineInfoTemplate)
	})

	data,err := json.Marshal(pipelineInfos)
	return data, err
}

func IterActionState(c *context.Context) ([]byte, error) {
	actionId := c.ParamsInt64(":actionId")
	var action *RuntimePipeline
	for i:=e.actions.Front(); i!=nil; i=i.Next() {
		cur := (i.Value).(*RuntimePipeline)
		if cur.ID == actionId {
			action = cur
		}
	}
	actionState := Unknown
	if action != nil {
		actionState = action.Status
	}else {
		iterationAction, _ := db.GetIterActionById(actionId)
		actionState = actionState.FromString(iterationAction.State)
	}
	return json.Marshal(actionState.ToString())
}

func IterStageState(c *context.Context) ([]byte, error) {
	// iterationId, stageId
	actionId := c.ParamsInt64(":actionId")
	stageId := c.ParamsInt64(":stageId")
	// 1.search e.actions
	// 2.query db
	var action *RuntimePipeline
	for i := e.actions.Front(); i != nil; i=i.Next() {
		cur := (i.Value).(*RuntimePipeline)
		if cur.ID == actionId {
			// cur is a pointer, insure keeping is read only
			action = cur
			break
		}
	}

	stageState := stage.Canceled
	if action != nil {
		for i := 0; i < len(action.Buckets); i++ {
			if action.Buckets[i].StageId == stageId {
				stageState = action.Buckets[i].State
				break
			}
		}

	} else{
		// compatible time delay
		// query from db
		stageExec,_ := db.QueryStageExec(actionId, stageId)
		if stageExec != nil {
			stageState = stageState.FromString(stageExec.State)
		}
	}

	data, err := json.Marshal(stageState.ToString())

	return data,err
}

func IterStepLog(c *context.Context) ([]byte, error) {
	appName := c.Query("appName")
	actionId := c.ParamsInt64(":actionId")
	stageId := c.ParamsInt64(":stageId")
	stepId := c.ParamsInt64(":stepId")
	logType := db.GetStepLogTypeById(stepId)
	stageExec,_ := db.QueryStageExec(actionId, stageId)
	if stageExec != nil {
		stepExec,_ := db.GetStepExecByStageExecIdAndStepId(stageExec.Id, stepId)
		return step.GetLog(stepExec.LogPath, stepExec.ExecPath, logType, appName), nil
	} else {
		return nil, nil
	}
}

func IterStepTestCodeCovered(c *context.Context) ([]byte, error) {
	appName := c.Query("appName")
	packageName := c.Query("package")
	fileName := c.Query("fileName")
	actionId := c.ParamsInt64(":actionId")
	stageId := c.ParamsInt64(":stageId")
	stepId := c.ParamsInt64(":stepId")
	stageExec,_ := db.QueryStageExec(actionId, stageId)
	if stageExec != nil {
		stepExec,_ := db.GetStepExecByStageExecIdAndStepId(stageExec.Id, stepId)
		return step.GetFileCovered(stepExec.ExecPath, appName, packageName, fileName), nil
	} else {
		return nil, nil
	}
}

func IterStepState(c *context.Context) ([]byte, error) {
	actionId := c.ParamsInt64(":actionId")
	stageId := c.ParamsInt64(":stageId")
	stepId := c.ParamsInt64(":stepId")

	var action *RuntimePipeline
	runtimeStepState := step.Unknown

	for i := e.actions.Front(); i != nil; i=i.Next() {
		cur := (i.Value).(*RuntimePipeline)
		if cur.ID == actionId {
			// cur is a pointer, insure keeping is read only
			action = cur
			break
		}
	}

	if action != nil {
		for j:=0; j<len(action.Buckets); j++ {
			for k:=0; k<len(action.Buckets[j].Steps); k++ {
				if action.Buckets[j].StageId == stageId && action.Buckets[j].Steps[k].StepId == stepId {
					runtimeStepState = action.Buckets[j].Steps[k].State
					break
				}
			}
		}
	} else {
		stageExec,_ := db.QueryStageExec(actionId, stageId)
		if stageExec!=nil {
			stepExec,_ := db.GetStepExecByStageExecIdAndStepId(stageExec.Id, stepId)
			if stepExec!=nil {
				runtimeStepState = runtimeStepState.FromString(stepExec.State)
			}
		}
	}
	data, err := json.Marshal(runtimeStepState.ToString())
	return data, err
}


func IterStageInfo(c *context.Context) ([]byte, error){
	actionId := c.ParamsInt64(":actionId")
	stageId := c.ParamsInt64(":stageId")

	var infoSteps []step.InfoStep
	stage,_ := db.GetStageById(stageId)
	steps,_ := db.BranchQueryStepsByIds(stage.Steps)
	stageExec, _ := db.GetStageExecIdByActIdAndStageId(actionId, stageId)
	for i:=0; i<len(steps); i++ {
		infoStep := step.InfoStep{Index: i, Image: steps[i].Img, Title: steps[i].Name, StepId: steps[i].ID}
		stepExec,_ := db.GetStepExecByStageExecIdAndStepId(stageExec, steps[i].ID)
		if stepExec != nil {
			infoStep.Link = stepExec.Link
		}
		infoSteps = append(infoSteps, infoStep)
	}

	data, err := json.Marshal(infoSteps)
	return data, err
}

func ParseDagAndLayout(pipeline *db.Pipeline)([][]stage.Endpoint, []edge, error){
	// 1. generate endpoints
	// 2. calculate position of each block
	// layout is a matrix which is same size as dag
	nodeNum := len(pipeline.Stages)
	endpoints := make([][]stage.Endpoint, nodeNum)
	var edges []edge

	for i := 0; i < nodeNum; i++ {
		for j := 0; j < nodeNum; j++ {
			// edge exist from i+1 to j+1
			if pipeline.StageDAG[i][j] == 1 {

				startNodeId := dagMapLayout(pipeline.StageLayout,i+1)
				endNodeId := dagMapLayout(pipeline.StageLayout, j+1)
				startNodeX, startNodeY,_ := findNodePosition(pipeline, nodeNum, startNodeId)
				endNodeX, endNodeY,_ := findNodePosition(pipeline, nodeNum, endNodeId)
				// startNode endpoints
				s,e := findEndpointPosition(startNodeX, startNodeY, endNodeX, endNodeY)
				makeEndpoint(&endpoints[i], &endpoints[j],i+1, j+1, s, e)
				makeEdge(&edges,i+1, j+1, s, e)
			}
		}
	}


	return endpoints, edges, nil
}

func ParseLayout(p *db.Pipeline) (maxRow, maxCol int) {
	nodeNum := len(p.Stages)
	var max = func(x,y int) int{
		if x > y{
			return x
		}
		return y
	}
	for i := 0; i < nodeNum; i++ {
		for j := 0; j < nodeNum; j++ {
			if p.StageLayout[i][j]!=0 {
				maxRow = max(maxRow, i+1)
				maxCol = max(maxCol, j+1)
			}
		}
	}
	return
}

var findNodePosition = func(pipeline *db.Pipeline, nodeNum int, nodeId int64)(int, int, error) {
	for i := 0; i < nodeNum; i++ {
		for j := 0; j < nodeNum; j++ {
			if pipeline.StageLayout[i][j] == nodeId {
				return i,j,nil
			}
		}
	}
	return -1,-1, DataError{Args: map[string]interface{}{"pipelineId":pipeline.ID,"pipelineName":pipeline.Name, "creatorName":pipeline.CreatorName}}
}

var dagMapLayout = func(layout [][]int64, index int) int64 {
	var counter int
	var nodeId int64
	nodeNum := len(layout)
	for i := 0; i < nodeNum; i++ {
		for j := 0; j < nodeNum; j++ {
			if layout[i][j] > 0 {
				counter++
			}
			if counter ==index {
				nodeId = layout[i][j]
				return nodeId
			}
		}
	}
	return nodeId
}

var makeEndpoint = func(sEndpoints,eEndpoints *[]stage.Endpoint, sIndex, eIndex int, s,e EndpointPosition) {
	/** 0.for obeying the rules, we a not allow two same endpoint in a node
	  1.if node a above node b and exist an edge from a to b, make a top endpoint to node b, make a right endpoint to node a
	  2.if node a below node b and exist an edge from a to b, make a bottom endpoint to node b, make a right endpoint to node a
	  3.if node a left at node b and exist an edge from a to b, make a left endpoint to node b, make a right endpoint to node a
	  4.if node a right at node b ans exist an edge from a to b, make a right endpoint to node b, make a right endpoint to node a
	  5.traverse all edges and judge each node between the two side of the edge
	*/
	// nodeId this as same as stageId which is unique in an pipeline
	var sExist, eExist bool

	for _, v := range *sEndpoints {
		if v.Id == s.String(sIndex) {
			sExist = true
			break
		}
	}
	if !sExist {
		sNewEndpoint := EndpointType[s]
		sNewEndpoint.FormatId(sIndex)
		*sEndpoints = append(*sEndpoints, sNewEndpoint)
	}

	for _, v := range *eEndpoints {
		if v.Id == e.String(eIndex) {
			eExist = true
			break
		}
	}
	if !eExist {
		eNewEndpoint := EndpointType[e]
		eNewEndpoint.FormatId(eIndex)
		*eEndpoints = append(*eEndpoints, eNewEndpoint)
	}

}

var makeEdge = func(edges *[]edge, sIndex, eIndex int, s, e EndpointPosition) {
	newEdge := EdgeType[s][e]
	newEdge.Format(sIndex,eIndex)
	*edges = append(*edges, newEdge)
}

var makeGroups = func(p *db.Pipeline) ([]group, error){
	var groups []group
	// fill Top Left Width, Height
	// maxRow width of p.StageLayout, maxCol height of p.StageLayout
	maxRow,maxCol := ParseLayout(p)
	groups = append(groups, group{Id: "group", Draggable:false, Resize: false, Options: groupOptions{}, Top: GroupTop, Left: GroupLeft,
		Width: FirstGroupPY+maxCol*GroupYStep, Height: FirstGroupPX+maxRow*GroupXStep})
	return groups,nil
}

var makeNode = func(p *db.Pipeline, endpoints [][]stage.Endpoint)([]stage.Node, error) {
	var nodes []stage.Node
	for k,ep := range endpoints {
		// index of the node is k+1
		nodeId := dagMapLayout(p.StageLayout,k+1)
		nodeX,nodeY,_ := findNodePosition(p, len(p.Stages), nodeId)
		node := stage.Node{Id: int64(k+1), Endpoints: ep, Group: "group", Top: FirstStagePX+nodeX*StageXStep, Left: FirstStagePY+nodeY*StageYStep, StageId: nodeId}
		// get node info(stage info)
		nodes = append(nodes, node)
	}
	return nodes, nil
}

var findEndpointPosition = func(sX, sY, eX, eY int) (s, e EndpointPosition) {
	if sX<eX && sY<eY {
		//s right, e top
		s = RIGHT
		e = TOP
	} else if sX<eX && sY==eY {
		//s bottom, e top
		s = BOTTOM
		e = TOP
	} else if sX<eX && sY>eY {
		//s left, e top
		s = LEFT
		e = TOP
	} else if sX==eX && sY<eY {
		//s right, e left
		s = RIGHT
		e = LEFT
	} else if sX==eX && sY>eY {
		//s left, e right
		s = LEFT
		e = RIGHT
	} else if sX>eX && sY<eY {
		//s right, e bottom
		s = RIGHT
		e = BOTTOM
	} else if sX>eX && sY==eY {
		//s top, e bottom
		s = TOP
		e = BOTTOM
	} else if sX>eX && sY>eY {
		//s left, e bottom
		s = LEFT
		e = BOTTOM
	}

	return s,e
}

func buildPipelineEnv() map[string]interface{}{
	return nil
}

type MergeRequest struct {
	UserId       int64  `json:"userId"`
	UserName     string `json:"userName"`
	Repository   string `json:"repository"`
	SourceBranch string `json:"sourceBranch"`
	TargetBranch string `json:"targetBranch"`
	MergeInfo    string `json:"mergeInfo"`
}

func concatInt64Slice(array []int64) string{
	var arrayString []string
	for _, v := range array {
		arrayString = append(arrayString, strconv.FormatInt(v, 10))
	}
	return strings.Join(arrayString, "-")
}

type DataError struct {
	Args map[string]interface{}
}

func (err DataError) Error() string {
	return fmt.Sprintf("pipeline data error: %v", err.Args)
}
