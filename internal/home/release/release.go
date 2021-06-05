package release

import (
	"encoding/json"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/db"
)

type ReleaseRecord struct {
	OwnerName   string `json:"ownerName"`
	RepoName    string `json:"repoName"`
	CommitId    string `json:"commitId"`
	CommitLink  string `json:"commitLink"`
	Time        string `json:"time"`
	IterationId int64  `json:"iterationId"`
	URL         string `json:"url"`
	ID          int64  `json:"id"`
}

func QueryReleaseHistory(c *context.Context) ([]byte, error) {

	appOwner := c.Query("ownerName")
	appName := c.Query("repoName")

	records := make([]ReleaseRecord, 0)
	releases, _ := db.GetReleaseByAppOwnerAndAppName(appOwner, appName)
	for _, r := range releases {
		records = append(records, ReleaseRecord{
			OwnerName: r.AppOwner, RepoName: r.AppName, CommitId: r.CommitId, Time: r.Time, IterationId: r.IterationId, URL: r.URL, ID: r.ID, CommitLink: r.CommitLink,
		})
	}
	data, err := json.Marshal(records)
	return data, err
}


