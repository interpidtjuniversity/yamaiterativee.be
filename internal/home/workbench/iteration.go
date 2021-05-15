package workbench

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/grpc/invoke/invokerImpl"
	"yama.io/yamaIterativeE/internal/util"
)

func GetAllOwners(c *context.Context) []byte {
	names, err := invokerImpl.InvokeQueryApplicationOwners()
	if err != nil {
		return nil
	}
	data, _ := json.Marshal(names)
	return data
}

func GetOwnerApplications(c *context.Context) []byte {
	ownerName := c.Params(":ownerName")
	names, err := invokerImpl.InvokeQueryApplications(ownerName)
	if err != nil {
		return []byte("[]")
	}

	data, _ := json.Marshal(names)
	return data
}

func NewIteration(c *context.Context) []byte{
	creator := c.Query("creator")
	appOwner := c.Query("appOwner")
	appName := c.Query("appName")
	admins := strings.Split(c.Query("admins"),",")
	iterType := c.Query("iterType")
	content := c.Query("content")
	baseCommit,_ := invokerImpl.InvokeQueryMasterLatestCommitIdService(appOwner, appName)
	app,_ := db.GetApplicationConfigByOwnerAndRepo(appOwner, appName)


	var newIterationFunc = func(session *db.ProxySession) (map[string]interface{}, error){
		iteration := db.Iteration{IterCreator: creator, IterType: iterType, IterAdmin: admins, OwnerName: appOwner, RepoName: appName,
			IterBranch: generateIterBranch(), IterState: -1, BaseCommit: baseCommit, Content: content, Title: util.GenerateRandomStringWithPrefix(10,"E"),
			DevConfig: app.DevConfig, StableConfig: app.StableConfig, TestConfig: app.TestConfig, PreConfig: app.PreConfig, ProdConfig: app.ProdConfig,
		}
		// insert iteration
		_, err := session.Session.Insert(&iteration)
		if err != nil {
			return nil, err
		}
		iag := db.IterationActGroup{TargetBranch: iteration.IterBranch, IterationId: iteration.ID, GroupType: "dev"}
		// insert iterationactgroup
		_,err = session.Session.Insert(&iag)
		if err != nil {
			return nil, err
		}
		// update iteration
		iteration.IterDevActGroup = iag.ID
		_, err = session.Session.ID(iteration.ID).Update(&iteration)
		if err != nil {
			return nil, err
		}
		blackBranch := db.BlackBranch{Branch: iteration.IterBranch,AppOwner: iteration.OwnerName, AppName: iteration.RepoName}
		_, err = session.Session.Table("black_branch").Insert(&blackBranch)
		if err != nil {
			return nil, err
		}

		return map[string]interface{}{
			"iterBranch": iteration.IterBranch,
			"ownerName": iteration.OwnerName,
			"repoName": iteration.RepoName,
			"protected": true,
			"needMR": true,
		} ,nil
	}
	data, err := db.Proxy.TransactionExecute(newIterationFunc)
	if err != nil {
		return []byte(err.Error())
	}
	// invoke rpc
	err = invokerImpl.InvokeCreateBranchService(data)
	if err != nil {
		return []byte(err.Error())
	}
	return []byte("success")
}


func generateIterBranch() string {
	y,m,d := time.Now().Date()
	return util.GenerateRandomStringWithSuffix(10,fmt.Sprintf("%d_%d_%d", y,m,d))
}
