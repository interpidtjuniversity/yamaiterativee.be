package iteration

import (
	"encoding/json"
	"sync"
	"yama.io/yamaIterativeE/internal/context"
)

type stageInfo struct {
	Title string `json:"title"`
	Img   string `json:"img"`
}

type pipelineInfo struct {
	PipelineId int64   `json:"pipelineId"`
	AvatarSrc  string  `json:"avatarSrc"`
	ActionInfo string  `json:"actionInfo"`
	ExtInfo    string  `json:"extInfo"`
	Nodes      []node  `json:"nodes"`
	Edges      []edge  `json:"edges"`
	Groups      []group `json:"groups"`
}

type node struct {
	StageIdExecId string     `json:"stageId_execId"`
	StageId       int64      `json:"stageId"`
	ExecId        int64      `json:"execId"`
	Id            int64      `json:"id"`
	Label         string     `json:"label"`
	ClassName     string     `json:"className"`
	IconType      string     `json:"iconType"`
	Top           int        `json:"top"`
	Left          int        `json:"left"`
	Group         string     `json:"group"`
	Endpoints     []endpoint `json:"endpoints"`
}

type endpoint struct {
	Id          string     `json:"id"`
	Orientation []int     `json:"orientation"`
	Pos         []float64 `json:"pos"`
}

type edge struct {
	Source        string  `json:"source"`
	Target        string  `json:"target"`
	SourceNode    int64   `json:"sourceNode"`
	TargetNode    int64   `json:"targetNode"`
	Arrow         bool    `json:"arrow"`
	Type          string  `json:"type"`
	ArrowPosition float64 `json:"arrowPosition"`
	ShapeType     string  `json:"shapeType"`
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

var p *ping = &ping{counter: 0, m: map[int64]string{0:"loading", 1:"warning", 2:"success", 3:"error"}, colorM: map[int64]string{0:"1F49E0", 1:"#FFA003", 2:"#1DC11D", 3:"#FF3333"}}
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

func IterPipelineInfo(c *context.Context) []byte {
	iterationId := c.ParamsInt64(":iterationId")
	envType := c.ParamsEscape(":envType")

	var endpoint1_nodes []endpoint
	var endpoint2_nodes []endpoint
	var endpoint3_nodes []endpoint
	var endpoint4_nodes []endpoint
	var endpoint5_nodes []endpoint
	var endpoint6_nodes []endpoint
	var endpoint7_nodes []endpoint

	endpoint1_nodes = append(endpoint1_nodes, endpoint{Id: "1_right", Orientation: []int{1,0}, Pos: []float64{0, 0.5}})
	endpoint2_nodes = append(endpoint2_nodes, endpoint{Id: "2_right", Orientation: []int{1,0}, Pos: []float64{0, 0.5}})
	endpoint3_nodes = append(endpoint3_nodes, endpoint{Id: "3_left", Orientation: []int{-1,0}, Pos: []float64{0, 0.5}})
	endpoint3_nodes = append(endpoint3_nodes, endpoint{Id: "3_right", Orientation: []int{1,0}, Pos: []float64{0, 0.5}})
	endpoint4_nodes = append(endpoint4_nodes, endpoint{Id: "4_left", Orientation: []int{-1,0}, Pos: []float64{0, 0.5}})
	endpoint4_nodes = append(endpoint4_nodes, endpoint{Id: "4_right", Orientation: []int{1,0}, Pos: []float64{0, 0.5}})
	endpoint5_nodes = append(endpoint5_nodes, endpoint{Id: "5_left", Orientation: []int{-1,0}, Pos: []float64{0, 0.5}})
	endpoint5_nodes = append(endpoint5_nodes, endpoint{Id: "5_right", Orientation: []int{1,0}, Pos: []float64{0, 0.5}})
	endpoint5_nodes = append(endpoint5_nodes, endpoint{Id: "5_top", Orientation: []int{0,-1}, Pos: []float64{0.5, 0}})
	endpoint6_nodes = append(endpoint6_nodes, endpoint{Id: "6_left", Orientation: []int{-1,0}, Pos: []float64{0, 0.5}})
	endpoint6_nodes = append(endpoint6_nodes, endpoint{Id: "6_right", Orientation: []int{1,0}, Pos: []float64{0,0.5}})
	endpoint7_nodes = append(endpoint7_nodes, endpoint{Id: "7_left", Orientation: []int{-1,0}, Pos: []float64{0,0.5}})

	var nodes []node
	nodes = append(nodes, node{Endpoints: endpoint1_nodes, StageIdExecId: "1_1", StageId: 1, ExecId: 1, Id: 1, Label: "代码评审", ClassName: "icon-background-color", IconType: "icon-kaifa", Top: 55, Left: 50, Group: "group"})
	nodes = append(nodes, node{Endpoints: endpoint2_nodes, StageIdExecId: "2_1", StageId: 2, ExecId: 1, Id: 2, Label: "冲突检测", ClassName: "icon-background-color", IconType: "icon-kaifa", Top: 125, Left: 50, Group: "group"})
	nodes = append(nodes, node{Endpoints: endpoint3_nodes, StageIdExecId: "3_1", StageId: 3, ExecId: 1, Id: 3, Label: "代码扫描", ClassName: "icon-background-color", IconType: "icon-kaifa", Top: 125, Left: 225, Group: "group"})
	nodes = append(nodes, node{Endpoints: endpoint4_nodes, StageIdExecId: "4_1", StageId: 4, ExecId: 1, Id: 4, Label: "预编译", ClassName: "icon-background-color", IconType: "icon-kaifa", Top: 125, Left: 400, Group: "group"})
	nodes = append(nodes, node{Endpoints: endpoint5_nodes, StageIdExecId: "5_1", StageId: 5, ExecId: 1, Id: 5, Label: "合并", ClassName: "icon-background-color", IconType: "icon-kaifa", Top: 125, Left: 575, Group: "group"})
	nodes = append(nodes, node{Endpoints: endpoint6_nodes, StageIdExecId: "6_1", StageId: 6, ExecId: 1, Id: 6, Label: "合并后编译", ClassName: "icon-background-color", IconType: "icon-kaifa", Top: 125, Left: 750, Group: "group"})
	nodes = append(nodes, node{Endpoints: endpoint7_nodes, StageIdExecId: "7_1", StageId: 7, ExecId: 1, Id: 7, Label: "质量检测", ClassName: "icon-background-color", IconType: "icon-kaifa", Top: 125, Left: 925, Group: "group"})

	var edges []edge
	edges = append(edges, edge{Source: "1_right", Target: "5_top", SourceNode: 1, TargetNode: 5, Arrow: true, Type: "endpoint", ArrowPosition: 0.5, ShapeType: "Flow"})
	edges = append(edges, edge{Source: "2_right", Target: "3_left", SourceNode: 2, TargetNode: 3, Arrow: true, Type: "endpoint", ArrowPosition: 0.5})
	edges = append(edges, edge{Source: "3_right", Target: "4_left", SourceNode: 3, TargetNode: 4, Arrow: true, Type: "endpoint", ArrowPosition: 0.5})
	edges = append(edges, edge{Source: "4_right", Target: "5_left", SourceNode: 4, TargetNode: 5, Arrow: true, Type: "endpoint", ArrowPosition: 0.5})
	edges = append(edges, edge{Source: "5_right", Target: "6_left", SourceNode: 5, TargetNode: 6, Arrow: true, Type: "endpoint", ArrowPosition: 0.5})
	edges = append(edges, edge{Source: "6_right", Target: "7_left", SourceNode: 6, TargetNode: 7, Arrow: true, Type: "endpoint", ArrowPosition: 0.5})

	var groups []group
	groups = append(groups, group{Id: "group", Draggable: false, Top: 0, Left: 130, Width: 1100, Height: 225, Resize: false, Options: groupOptions{Title: "E123456789_zqf0304_dev -> E123456789_20210304"}})



	var endpoint8_nodes []endpoint
	var endpoint9_nodes []endpoint
	var endpoint10_nodes []endpoint

	endpoint8_nodes = append(endpoint8_nodes, endpoint{Id: "1_right", Orientation: []int{1,0}, Pos: []float64{0, 0.5}})
	endpoint9_nodes = append(endpoint9_nodes, endpoint{Id: "2_left", Orientation: []int{-1,0}, Pos: []float64{0, 0.5}})
	endpoint9_nodes = append(endpoint9_nodes, endpoint{Id: "2_right", Orientation: []int{1,0}, Pos: []float64{0, 0.5}})
	endpoint10_nodes = append(endpoint10_nodes, endpoint{Id: "3_left", Orientation: []int{-1,0}, Pos: []float64{0, 0.5}})

	var nodes2 []node
	var edges2 []edge

	nodes2 = append(nodes2, node{StageIdExecId: "8_1", StageId: 8, ExecId: 1, Label: "机器变更", ClassName: "icon-background-color", IconType: "icon-kaifa", Top: 55, Left: 50, Group: "group", Endpoints: endpoint8_nodes, Id: 1})
	nodes2 = append(nodes2, node{StageIdExecId: "9_1", StageId: 9, ExecId: 1, Label: "镜像构建", ClassName: "icon-background-color", IconType: "icon-kaifa", Top: 55, Left: 300, Group: "group", Endpoints: endpoint9_nodes, Id: 2})
	nodes2 = append(nodes2, node{StageIdExecId: "10_1", StageId: 10, ExecId: 1, Label: "发布", ClassName: "icon-background-color", IconType: "icon-kaifa", Top: 55, Left: 550, Group: "group", Endpoints: endpoint10_nodes, Id: 3})

	edges2 = append(edges2, edge{Source: "1_right", Target: "2_left", SourceNode: 1, TargetNode: 2, Arrow: true, Type: "endpoint", ArrowPosition: 0.5})
	edges2 = append(edges2, edge{Source: "2_right", Target: "3_left", SourceNode: 2, TargetNode: 3, Arrow: true, Type: "endpoint", ArrowPosition: 0.5})

	var groups2 []group
	groups2 = append(groups2, group{Id: "group", Draggable: false, Top: 0, Left: 130, Width: 750, Height: 150, Resize: false, Options: groupOptions{Title: "张启帆 申请了服务器 E987654321 (3天前)"}})


	basicMR := pipelineInfo{Nodes: nodes, Edges: edges, Groups: groups, PipelineId: 2, ActionInfo: "张启帆 给MR：#999999 的源分支提交代码触发了Pipeline #10000000 开发环境", AvatarSrc: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png", ExtInfo: "this is extInfo"}
	serverApply := pipelineInfo{Nodes: nodes2, Edges: edges2, Groups: groups2, PipelineId: 1, ActionInfo: "张启帆 申请了服务器 E987654321 (3天前)", AvatarSrc: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png", ExtInfo: "this is extInfo"}



	if iterationId == 1 {
		if envType == "dev" {
			data, _ := json.Marshal([]pipelineInfo{serverApply, basicMR})
			return data
		} else if envType == "pre" {
			basicMR.PipelineId=3
			basicMR.ActionInfo="张启帆 给MR：#999999 的源分支提交代码触发了Pipeline #10000000 预发环境"
			basicMR.Groups[0].Options.Title = "E123456789_zqf0304_pre -> E123456789_20210304"
			data, _ := json.Marshal([]pipelineInfo{basicMR})
			return data
		}
	}
	data,_ := json.Marshal([]pipelineInfo{})
	return data
}

func StageInfo(c *context.Context) []byte{
	stageId := c.ParamsInt64(":stage")

	var code_review []stageInfo
	var conflict_detect []stageInfo
	var code_scan []stageInfo
	var pre_compile []stageInfo
	var merge []stageInfo
	var compile []stageInfo
	var quality_detect []stageInfo
	var server_change []stageInfo
	var image_build []stageInfo
	var release []stageInfo

	code_review = append(code_review, stageInfo{Title: "孙武", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png"})
	code_review = append(code_review, stageInfo{Title: "孔子", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png"})

	conflict_detect = append(conflict_detect, stageInfo{Title: "代码预合并", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png"})

	code_scan = append(code_scan, stageInfo{Title: "静态扫描", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png"})
	code_scan = append(code_scan, stageInfo{Title: "PMD", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png"})

	pre_compile = append(pre_compile, stageInfo{Title: "mvn compile", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png"})

	merge = append(merge, stageInfo{Title: "代码合并", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png"})

	compile = append(compile, stageInfo{Title: "mvn compile", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png"})

	quality_detect = append(quality_detect, stageInfo{Title: "单元测试", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png"})
	quality_detect = append(quality_detect, stageInfo{Title: "集成测试", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png"})

	server_change = append(server_change, stageInfo{Title: "服务器安装", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png"})

	image_build = append(image_build, stageInfo{Title: "环境安装", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png"})
	image_build = append(image_build, stageInfo{Title: "服务构建", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png"})

	release = append(release, stageInfo{Title: "服务部署", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png"})
	release = append(release, stageInfo{Title: "服务测试", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png"})


	switch stageId {
	case 1:
		data, _ := json.Marshal(code_review)
		return data
	case 2:
		data, _ := json.Marshal(conflict_detect)
		return data
	case 3:
		data, _ := json.Marshal(code_scan)
		return data
	case 4:
		data, _ := json.Marshal(pre_compile)
		return data
	case 5:
		data, _ := json.Marshal(merge)
		return data
	case 6:
		data, _ := json.Marshal(compile)
		return data
	case 7:
		data, _ := json.Marshal(quality_detect)
		return data
	case 8:
		data,_ := json.Marshal(server_change)
		return data
	case 9:
		data,_ := json.Marshal(image_build)
		return data
	case 10:
		data,_ := json.Marshal(release)
		return data
	default:
		return nil
	}
}

func StageExecInfo(c *context.Context) {
	//stageId := c.ParamsInt64(":stage")
	//execId := c.ParamsInt64(":exec")
}

type stepStatus struct {
	Type  string `json:"type"`
	Color string `json:"color"`
}


func Ping(c *context.Context) []byte{
	p.mu.Lock()
	msg := stepStatus{Type: p.m[int64(int(p.counter)%len(p.m))], Color: p.colorM[int64(int(p.counter)%len(p.colorM))]}
	p.counter++
	p.mu.Unlock()
	data, _ := json.Marshal(msg)
	return data
}

var ieah = &iterEnvActionHolder{m: map[string][]iterEnvAction{"dev": {{ "完成开发阶段"}, { "提交MR"}, {"Jar包管理"}, {"配置变更"}, {"触发pipeline"}, {"申请服务器"}, {"新建联调环境"}},
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