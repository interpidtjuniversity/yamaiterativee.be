package iteration

import (
	"encoding/json"
	"yama.io/yamaIterativeE/internal/context"
)

type stageInfo struct {
	Title string `json:"title"`
	Img   string `json:"img"`
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
	default:
		return nil
	}
}

func StageExecInfo(c *context.Context) {
	//stageId := c.ParamsInt64(":stage")
	//execId := c.ParamsInt64(":exec")
}
