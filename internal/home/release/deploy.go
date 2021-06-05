package release

import (
	"fmt"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/iteration/step"
	"yama.io/yamaIterativeE/internal/util"
)

var PIPELINE_EXEC_PATH = "/root/yamaIterativeE/yamaIterativeE-pipeline-exec/%s"

func DeployHistoryRelease(c *context.Context) ([]byte, error) {
	appOwner := c.Query("ownerName")
	appName := c.Query("repoName")
	releaseId := c.QueryInt64("releaseId")

	prodServer, _ := db.GetApplicationProdServer(appOwner, appName)
	if len(prodServer) == 0{
		return nil, fmt.Errorf("no prod server for appOwner: %s, appName: %s", appOwner, appName)
	}
	release := db.GetReleaseById(releaseId)
	if release == nil{
		return nil, fmt.Errorf("no such release, id: %v", releaseId)
	}
	execPath := fmt.Sprintf(PIPELINE_EXEC_PATH, util.GenerateRandomStringWithSuffix(20,""))
	for _, server := range prodServer {
		err := step.RunCodeStep("historyReleaseDeployBean", release.URL, execPath, server.Name, server.IP)
		if err != nil {
			return nil, err
		}
	}
	db.UpdateProdServerReleaseId(appOwner, appName, release.CommitId)

	return []byte("success"), nil
}
