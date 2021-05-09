package iterations

import (
	"encoding/json"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/db"
)

type IterationData struct {
	Id          int64        `json:"id"`
	Title       string       `json:"title"`
	Owner       string       `json:"owner"`
	Application string       `json:"application"`
	State       string       `json:"state"`
	Content     string       `json:"content"`
	Src         string       `json:"src"`
	Color       string       `json:"color"`
	Creator     string       `json:"creator"`
	Members     []MemberData `json:"members"`
}
type MemberData struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var iterationColorMap = map[string]string{
	"basic MR":"green",
}

func GetUserAllIterations(context *context.Context) []byte {
	username := context.Params(":username")
	// get this user all iterations
	iters := db.GetIterationByAdmin(username)
	// agg user
	userQueryMap := make(map[string]bool)
	userMap := make(map[string]*db.User)
	var userNames []string
	for _,iter := range iters{
		for _, user := range iter.IterAdmin {
			userQueryMap[user] = true
		}
	}
	for user, _ := range userQueryMap {
		userNames = append(userNames, user)
	}
	users,_ := db.BranchQueryUserByName(userNames)
	for _, user := range users{
		userMap[user.Name] = user
	}
	// build data
	var iterationDatas []IterationData
	for _, iter := range iters{
		iterData := IterationData{
			Title: iter.Title,
			Id: iter.ID,
			Owner: iter.OwnerName,
			Application: iter.RepoName,
			State: iter.IterState.ToString(),
			Content: iter.Content,
			Src: "https://gw.alipayobjects.com/zos/rmsportal/pbmKMSFpLurLALLNliUQ.svg",
			Color: iterationColorMap[iter.IterType],
			Creator: iter.IterCreator,
		}
		var members []MemberData
		for _, member := range iter.IterAdmin {
			members = append(members, MemberData{member, userMap[member].Avatar})
		}
		iterData.Members = members
		iterationDatas = append(iterationDatas, iterData)
	}

	data, _ := json.Marshal(iterationDatas)
	return data
}
