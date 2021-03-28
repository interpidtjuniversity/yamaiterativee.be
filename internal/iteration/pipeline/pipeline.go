package pipeline

import (
	"encoding/json"
	"fmt"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/iteration/stage"
	"yama.io/yamaIterativeE/internal/util"
)

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
	PipelineId int64        `json:"pipelineId"`
	AvatarSrc  string       `json:"avatarSrc"`
	ActionInfo string       `json:"actionInfo"`
	ExtInfo    string       `json:"extInfo"`
	Nodes      []stage.Node `json:"nodes"`
	Edges      []edge       `json:"edges"`
	Groups     []group      `json:"groups"`
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


func StartPipeline(c *context.Context) ([]byte, error) {
	return nil,nil
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
	var pipelineInfos []pipelineInfo
	// create completely pipelineInfo with pipeline template and envActions
	stream.ForEach(func(action *db.IterationAction) {
		pipelineInfoTemplate := pipelineInfoMap[action.PipeLineId]
		// call stageInfo with actId and stageId
		var stageIds []int64
		for _,node := range pipelineInfoTemplate.Nodes {
			stageIds = append(stageIds, node.Id)
		}
		// agg StageExec and Stage
		stageExecs,_ := db.BranchQueryStageExec(action.ID, stageIds)
		stages,_ := db.BranchQueryStage(stageIds)
		stageExecMap := make(map[int64]*db.StageExec)
		stageMap := make(map[int64]*db.Stage)
		for _,exec := range stageExecs {
			stageExecMap[exec.StageId] = exec
		}
		for _,s :=range stages {
			stageMap[s.ID] = s
		}
		for i := 0; i<len(pipelineInfoTemplate.Nodes); i++ {
			pipelineInfoTemplate.Nodes[i].IconType = stageMap[pipelineInfoTemplate.Nodes[i].Id].IconType
			pipelineInfoTemplate.Nodes[i].ClassName = stageMap[pipelineInfoTemplate.Nodes[i].Id].ClassName
			pipelineInfoTemplate.Nodes[i].Group = stageMap[pipelineInfoTemplate.Nodes[i].Id].Group
			pipelineInfoTemplate.Nodes[i].Label = stageMap[pipelineInfoTemplate.Nodes[i].Id].Name
			pipelineInfoTemplate.Nodes[i].ExecId = stageExecMap[pipelineInfoTemplate.Nodes[i].Id].ID
			pipelineInfoTemplate.Nodes[i].StageIdExecId = fmt.Sprintf("%d_%d",pipelineInfoTemplate.Nodes[i].Id, stageExecMap[pipelineInfoTemplate.Nodes[i].Id].ID)
		}
		pipelineInfoTemplate.PipelineId=action.ID
		pipelineInfoTemplate.ActionInfo=action.ActionInfo
		pipelineInfoTemplate.AvatarSrc=action.AvatarSrc
		pipelineInfoTemplate.ExtInfo=action.ExtInfo
		// only one group exist so we get index of 0
		pipelineInfoTemplate.Groups[0].Options.Title = action.ActionGroupInfo

		pipelineInfos = append(pipelineInfos, pipelineInfoTemplate)
	})

	data,err := json.Marshal(pipelineInfos)
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

				startNodeId := dagMapLayout(pipeline.StageLayout,i+1,nodeNum)
				endNodeId := dagMapLayout(pipeline.StageLayout, j+1, nodeNum)
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

var dagMapLayout = func(layout [][]int64, index, nodeNum int) int64 {
	var counter int
	var nodeId int64
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
		nodeId := dagMapLayout(p.StageLayout,k+1,len(p.Stages))
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

type DataError struct {
	Args map[string]interface{}
}

func (err DataError) Error() string {
	return fmt.Sprintf("pipeline data error: %v", err.Args)
}
