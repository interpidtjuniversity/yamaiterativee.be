package datafetcher

import (
	"encoding/json"
	"yama.io/yamaIterativeE/internal/db"
)

type LatestIterationFetcher struct {
	Fetcher
}

type LatestIterationInfo struct {
	IterTitle   string       `json:"iterTitle"`
	IterContent string       `json:"iterContent"`
	IterId      int64        `json:"iterId"`
	Members     []MemberData `json:"members"`
}

type MemberData struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func (f LatestIterationFetcher) Fetch(userName string, limit int) ([]byte, error) {
	iterations, _ := db.GetIterationByAdmin(userName, limit)
	// agg user
	userQueryMap := make(map[string]bool)
	userMap := make(map[string]*db.User)
	var userNames []string
	for _,iter := range iterations{
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
	var latestIterationInfos []LatestIterationInfo
	for _, iteration := range iterations {
		info := LatestIterationInfo{
			IterId: iteration.ID,
			IterTitle: iteration.Title,
			IterContent: iteration.Content,
		}
		var members []MemberData
		for _, member := range iteration.IterAdmin {
			members = append(members, MemberData{member, userMap[member].Avatar})
		}
		info.Members = members
		latestIterationInfos = append(latestIterationInfos, info)
	}
	data, err := json.Marshal(latestIterationInfos)
	return data, err
}

