package datafetcher

import (
	"encoding/json"
	"os"
	"runtime"
	"syscall"
	"time"
	"yama.io/yamaIterativeE/internal/db"
)

type ExecMergeRequestFetcher struct {
	Fetcher
	resultMap map[string]string
}

type UserMergeRequestInfo struct {
	Result     string `json:"result"`
	UserName   string `json:"userName"`
	Time       string `json:"time"`
	ActionInfo string `json:"actionInfo"`
}

func (f ExecMergeRequestFetcher) Fetch(userName string, limit int) ([]byte, error) {
	actions, _:= db.GetIterActionByUserNameAndPipelineId(userName, 1, limit)
	var infos []UserMergeRequestInfo
	for _, action := range actions {
		infos = append(infos, UserMergeRequestInfo{
			UserName: userName,
			ActionInfo: action.ActionInfo,
			Time: f.actionTime(action.ExecPath),
			Result: f.resultMap[action.State],
		})
	}
	data, _ := json.Marshal(infos)

	return data, nil
}

func (f ExecMergeRequestFetcher) actionTime(path string) string {
	var longTime int64
	osType := runtime.GOOS
	fileInfo, _ := os.Stat(path)
	if fileInfo == nil {
		longTime = time.Now().Unix()
	} else {
		if osType == "linux" {
			statT := fileInfo.Sys().(*syscall.Stat_t)
			longTime = statT.Ctim.Sec
		}
	}
	tm := time.Unix(longTime, 0)
	return tm.Format("2006-01-02 15:04:05")
}
