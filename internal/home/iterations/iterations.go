package iterations

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/grpc/invoke/invokerImpl"
	"yama.io/yamaIterativeE/internal/iteration/step"
	"yama.io/yamaIterativeE/internal/util"
)

var PIPELINE_EXEC_PATH = "/root/yamaIterativeE/yamaIterativeE-pipeline-exec/%s"

type IterationConfigGroup struct {
	DevConfig    []db.ConfigItem `json:"devConfig"`
	StableConfig []db.ConfigItem `json:"stableConfig"`
	TestConfig   []db.ConfigItem `json:"testConfig"`
	PreConfig    []db.ConfigItem `json:"preConfig"`
	ProdConfig   []db.ConfigItem `json:"prodConfig"`
}

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

type GrayStateData struct {
	State   bool   `json:"state"`
	Percent string `json:"percent"`
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

func GetIterationConfig(c *context.Context) []byte {
	iterId := c.ParamsInt64(":iterId")
	iteration, err := db.GetIterationConfigByIterId(iterId)
	if err != nil {
		return []byte(fmt.Sprintf("error while query iteration config, err: %v", err))
	}
	config := new(db.Config)

	json.Unmarshal([]byte(iteration.DevConfig), config)
	var devConfigItems []db.ConfigItem
	for _, v := range config.ConfigItems {
		if v.Displayable {
			v.Changeable = true
			devConfigItems = append(devConfigItems, v)
		}
	}

	json.Unmarshal([]byte(iteration.StableConfig), config)
	var stableConfigItems []db.ConfigItem
	for _, v := range config.ConfigItems {
		if v.Displayable {
			v.Changeable = true
			stableConfigItems = append(stableConfigItems, v)
		}
	}

	json.Unmarshal([]byte(iteration.TestConfig), config)
	var testConfigItems []db.ConfigItem
	for _, v := range config.ConfigItems {
		if v.Displayable {
			v.Changeable = true
			testConfigItems = append(testConfigItems, v)
		}
	}

	json.Unmarshal([]byte(iteration.PreConfig), config)
	var preConfigItems []db.ConfigItem
	for _, v := range config.ConfigItems {
		if v.Displayable {
			v.Changeable = true
			preConfigItems = append(preConfigItems, v)
		}
	}

	json.Unmarshal([]byte(iteration.ProdConfig), config)
	var prodConfigItems []db.ConfigItem
	for _, v := range config.ConfigItems {
		if v.Displayable {
			v.Changeable = true
			prodConfigItems = append(prodConfigItems, v)
		}
	}


	configs := IterationConfigGroup{
		DevConfig: devConfigItems,
		StableConfig: stableConfigItems,
		TestConfig: testConfigItems,
		PreConfig: preConfigItems,
		ProdConfig: prodConfigItems,
	}
	data, _ := json.Marshal(configs)
	return data
}


func GetIterationConfigByEnv(c *context.Context) []byte {
	iterId := c.ParamsInt64(":iterId")
	env := c.Params(":env")
	var cfg string
	config := new(db.Config)
	switch env {
	case "dev":
		cfg = db.GetIterationDevConfig(iterId)
		break
	case "stable":
		cfg = db.GetIterationStableConfig(iterId)
		break
	case "test":
		cfg = db.GetIterationTestConfig(iterId)
		break
	case "pre":
		cfg = db.GetIterationPreConfig(iterId)
		break
	case "prod":
		cfg = db.GetIterationProdConfig(iterId)
	}
	json.Unmarshal([]byte(cfg), config)

	var configItems []db.ConfigItem
	for _, v := range config.ConfigItems {
		if v.Displayable {
			v.Changeable = true
			configItems = append(configItems, v)
		}
	}
	data, _ := json.Marshal(configItems)
	return data
}

func ResetIterationConfig(c *context.Context) []byte {
	devCfg := c.Query("devConfig")
	stableCfg := c.Query("stableConfig")
	testCfg := c.Query("testConfig")
	preCfg := c.Query("preConfig")
	prodCfg := c.Query("prodConfig")
	iterId := c.QueryInt64("iterId")

	devItems := make([]db.ConfigItem,1)
	json.Unmarshal([]byte(devCfg), &devItems)
	devConfig := db.Config{
		ConfigItems: devItems,
	}
	devData, _ := json.Marshal(devConfig)

	stableItems := make([]db.ConfigItem,1)
	json.Unmarshal([]byte(stableCfg), &stableItems)
	stableConfig := db.Config{
		ConfigItems: stableItems,
	}
	stableData, _ := json.Marshal(stableConfig)

	testItems := make([]db.ConfigItem,1)
	json.Unmarshal([]byte(testCfg), &testItems)
	testConfig := db.Config{
		ConfigItems: testItems,
	}
	testData, _ := json.Marshal(testConfig)

	preItems := make([]db.ConfigItem,1)
	json.Unmarshal([]byte(preCfg), &preItems)
	preConfig := db.Config{
		ConfigItems: preItems,
	}
	preData, _ := json.Marshal(preConfig)

	prodItems := make([]db.ConfigItem,1)
	json.Unmarshal([]byte(prodCfg), &prodItems)
	prodConfig := db.Config{
		ConfigItems: prodItems,
	}
	prodData, _ := json.Marshal(prodConfig)

	iterationWithConfigs := &db.Iteration{
		DevConfig: string(devData),
		StableConfig: string(stableData),
		TestConfig: string(testData),
		PreConfig: string(preData),
		ProdConfig: string(prodData),
	}

	if _, err := db.UpdateIterationConfig(iterId, iterationWithConfigs); err!=nil {
		return []byte(fmt.Sprintf("error while reset iteration config, err: %v", err))
	}
	return nil
}

func GetIterationAllUsers(c *context.Context) []byte {
	iterId := c.ParamsInt64(":iterId")
	iteration, _ := db.GetIterationAllAdmins(iterId)
	if iteration == nil {
		return []byte("[]")
	}

	data, _ := json.Marshal(iteration.IterAdmin)
	return data
}

func AdvanceIteration(c *context.Context) []byte {
	iterId := c.QueryInt64("iterId")
	env := c.Params("env")
	curState := new(db.IterationState).FromString(env)
	nextState := curState.NextState()
	iteration,_ := db.GetIterationById(iterId)
	if nextState == db.PRE_STATE {
		err := sync(iteration, iteration.IterBranch, "master")
		if err != nil {
			return []byte("error")
		}
	} else if nextState == db.PROD_STATE {
		// 1. check if all prod server have deployed new version
		grayPercent := db.GetIterationGrayPercent(iterId)
		if grayPercent!="100" {
			return []byte("error")
		}
		// 2. get master commit add update prod server release_id
		masterCommitId, _ := invokerImpl.InvokeQueryMasterLatestCommitIdService(iteration.OwnerName, iteration.RepoName)
		db.UpdateProdServerReleaseId(masterCommitId)
		// 3. jar it to version repository
	}


	db.UpdateIterationState(iterId, nextState.ToString())
	actGroupId, _ := db.GetOrGenerateIterationActGroup(iterId, nextState.ToString())
	switch nextState {
	case db.ITG_STATE:
		db.UpdateIterationItgActGroup(iterId, actGroupId)
	case db.PRE_STATE:
		db.UpdateIterationPreActGroup(iterId, actGroupId)
	}
	return []byte("success")
}

func SyncMaster(c *context.Context) []byte {
	iterId := c.QueryInt64("iterId")
	iteration,_ := db.GetIterationById(iterId)
	err := sync(iteration, "master", iteration.IterBranch)
	if err!=nil {
		return []byte("error")
	}
	return []byte("success")
}

// source -> target
func sync(iteration *db.Iteration, sourceBranch, targetBranch string) error {
	// check if e_branch can merge to master
	repoURL := db.GetApplicationRepoByOwnerAndRepo(iteration.OwnerName, iteration.RepoName)
	execPath := fmt.Sprintf(PIPELINE_EXEC_PATH, util.GenerateRandomStringWithSuffix(20,""))
	os.MkdirAll(execPath, os.ModePerm)
	_,_,err := step.RunShellStep("/internal/iteration/step/command/pre-merge.sh", execPath, []string{repoURL,
		sourceBranch, targetBranch, iteration.RepoName}...)
	if err != nil {
		return err
	} else {
		_, _, err = step.RunShellStep("/internal/iteration/step/command/merge.sh", execPath, []string{
			iteration.OwnerName, iteration.RepoName, sourceBranch, targetBranch, "", "localhost:8000",
			"proto.YaMaHubBranchService/Merge2Branch", iteration.RepoName}...
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func AdvanceGray(c *context.Context) []byte {
	iterId := c.QueryInt64("iterId")
	iteration, _ := db.GetIterationById(iterId)
	if iteration.GrayPercent == "100" {
		return []byte("warning")
	}
	if iteration.GrayAdvanceState || iteration.GrayRollBackState{
		return []byte("running")
	}
	db.UpdateIterationAdvanceState(iterId, true)

	prodServer, _ := db.GetApplicationProdServer(iteration.OwnerName, iteration.RepoName)
	grayOrder := iteration.GrayOrder
	// prodServer -> grayOrder
	var newGray []string
	for i:=0; i<len(prodServer); i++ {
		var flag bool
		for j:=0; j<len(grayOrder); j++ {
			if prodServer[i].Name == grayOrder[j] {
				flag = true
				break
			}
		}
		if !flag {
			newGray = append(newGray, prodServer[i].Name)
		}
	}
	if len(newGray) == 0 {
		return []byte("warning")
	}
	// select one and deploy // select branch and deploy
	newGrayName := newGray[0]
	newGrayIP := db.GetServerIPByServerName(newGrayName)
	repo := db.GetApplicationRepoURLByOwnerAndRepo(iteration.OwnerName, iteration.RepoName)
	execPath := fmt.Sprintf(PIPELINE_EXEC_PATH, util.GenerateRandomStringWithSuffix(20,""))

	err := step.RunCodeStep("compileBean", iteration.OwnerName, iteration.RepoName, repo, "master", execPath,
		"prod", newGrayName, newGrayIP, strconv.Itoa(int(iterId)), "")
	if err!=nil {
		return []byte("error")
	}
	err = step.RunCodeStep("deployBean", iteration.RepoName, execPath, newGrayName, newGrayIP)

	// and then
	grayOrder = append(grayOrder, newGrayName)
	newPercent := float64(len(grayOrder))/float64(len(prodServer))
	db.UpdateIterationAdvanceGrayInfo(iterId, grayOrder, strconv.Itoa(int(newPercent*100)), false)
	if err!=nil {
		return []byte("error")
	}

	return []byte("success")
}

func RollBackGray(c *context.Context) []byte {
	iterId := c.QueryInt64("iterId")
	iteration, _ := db.GetIterationById(iterId)
	if iteration.GrayPercent == "" || iteration.GrayPercent == "0" || len(iteration.GrayOrder)==0 {
		return []byte("warning")
	}
	if iteration.GrayRollBackState || iteration.GrayAdvanceState{
		return []byte("running")
	}
	db.UpdateIterationRollBackState(iterId, true)
	prodServer, _ := db.GetApplicationProdServer(iteration.OwnerName, iteration.RepoName)
	grayServerNums := len(iteration.GrayOrder)
	newPercent := float64(grayServerNums-1)/float64(len(prodServer))
	rollBackServerName := iteration.GrayOrder[grayServerNums-1]
	rollBackServer,_ := db.GetServerByName(rollBackServerName)
	repo := db.GetApplicationRepoURLByOwnerAndRepo(iteration.OwnerName, iteration.RepoName)

	err := step.RunCodeStep("rollBackBean", iteration.OwnerName, iteration.RepoName, repo, "master", fmt.Sprintf(PIPELINE_EXEC_PATH, util.GenerateRandomStringWithSuffix(20,"")),
		"prod", rollBackServerName, rollBackServer.IP, strconv.Itoa(int(iterId)), rollBackServer.ReleaseId, "")

	db.UpdateIterationRollBackGrayInfo(iterId, iteration.GrayOrder[:grayServerNums-1], strconv.Itoa(int(newPercent*100)),false)
	if err!=nil {
		return []byte("error")
	}

	return nil

}

func GetIterationAdvanceGrayState(c *context.Context) []byte {
	iterId := c.ParamsInt64("iterId")
	percent, state := db.GetIterationAdvanceGrayState(iterId)
	if percent == "" {
		percent = "0"
	}
	gsd := GrayStateData{State: state, Percent: percent}
	data,_:=json.Marshal(gsd)
	return data
}

func GetIterationRollBackGrayState(c *context.Context) []byte {
	iterId := c.ParamsInt64("iterId")
	percent, state := db.GetIterationRollBackGrayState(iterId)
	if percent == "" {
		percent = "0"
	}
	gsd := GrayStateData{State: state, Percent: percent}
	data,_:=json.Marshal(gsd)
	return data
}