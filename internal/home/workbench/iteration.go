package workbench

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/form"
	"yama.io/yamaIterativeE/internal/grpc/invoke/invokerImpl"
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
		return nil
	}

	data, _ := json.Marshal(names)
	return data
}

func NewIteration(c *context.Context, f form.Iteration) []byte{

	var newIterationFunc = func(session *db.ProxySession) (map[string]interface{}, error){
		iteration := db.Iteration{IterCreatorId: f.Creator, IterType: f.Type, IterAdmin: f.Admin, OwnerName: f.Owner, RepoName: f.Repo,
			IterBranch: generateIterBranch(), IterState: 0}
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
		_, err = session.Session.Update(iteration)
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
	uid, _ := uuid.NewUUID()
	y,m,d := time.Now().Date()
	id := fmt.Sprintf("%s_%s", strings.ReplaceAll(uid.String(), "-",""), fmt.Sprintf("%d_%d_%d", y,m,d))
	return id
}
