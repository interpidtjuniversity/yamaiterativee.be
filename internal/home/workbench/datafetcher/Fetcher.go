package datafetcher

import "strings"

var fetcherMap = make(map[string]Fetcher)

func init() {
	fetcherMap["default"] = &DefaultFetcher{}
	fetcherMap["exec_execMergeRequest"] = &ExecMergeRequestFetcher{resultMap: map[string]string{"Finish":"已合并", "Canceled": "已取消", "Failure":"不通过"}}
	fetcherMap["latestIteration"] = &LatestIterationFetcher{}
}

type Fetcher interface {
	Fetch(userName string, limit int) ([]byte, error)
}

func GetFetcher(keys ...string) Fetcher {
	key := strings.Join(keys, "_")
	if _, ok := fetcherMap[key]; ok {
		return fetcherMap[key]
	}
	return fetcherMap["default"]
}
