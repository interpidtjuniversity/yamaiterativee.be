package stage

import (
	"encoding/json"
	"fmt"
	"sync"
	"yama.io/yamaIterativeE/internal/context"
)

type stageInfo struct {
	Title  string `json:"title"`
	Img    string `json:"img"`
	Index  int64  `json:"index"`
	// if is running, then axios pre 5 seconds task to query it until is failed||succeed
	Status string `json:"status"`
}

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

var p *ping = &ping{counter: 0, m: map[int64]string{0: "loading", 1:"warning", 2:"success", 3:"error"}, colorM: map[int64]string{0: "1F49E0", 1:"#FFA003", 2:"#1DC11D", 3:"#FF3333"}}
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

	code_review = append(code_review, stageInfo{Title: "孙武", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png", Index: 0})
	code_review = append(code_review, stageInfo{Title: "孔子", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png", Index: 1})

	conflict_detect = append(conflict_detect, stageInfo{Title: "代码预合并", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png", Index: 0})

	code_scan = append(code_scan, stageInfo{Title: "静态扫描", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png",Index: 0})
	code_scan = append(code_scan, stageInfo{Title: "PMD", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png",Index: 1})

	pre_compile = append(pre_compile, stageInfo{Title: "mvn compile", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png",Index: 0})

	merge = append(merge, stageInfo{Title: "代码合并", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png",Index: 0})

	compile = append(compile, stageInfo{Title: "mvn compile", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png",Index: 0})

	quality_detect = append(quality_detect, stageInfo{Title: "单元测试", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png",Index: 0})
	quality_detect = append(quality_detect, stageInfo{Title: "集成测试", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png",Index: 1})

	server_change = append(server_change, stageInfo{Title: "服务器安装", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png",Index: 0})

	image_build = append(image_build, stageInfo{Title: "环境安装", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png",Index: 0})
	image_build = append(image_build, stageInfo{Title: "服务构建", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png",Index: 1})

	release = append(release, stageInfo{Title: "服务部署", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png",Index: 0})
	release = append(release, stageInfo{Title: "服务测试", Img: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png",Index: 1})


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