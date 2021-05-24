package application

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/form"
	"yama.io/yamaIterativeE/internal/grpc/invoke/invokerImpl"
	"yama.io/yamaIterativeE/internal/grpc/invoke/invokerarg"
	"yama.io/yamaIterativeE/internal/home/config"
	"yama.io/yamaIterativeE/internal/resource"
	"yama.io/yamaIterativeE/internal/util"
)

type ApplicationData struct {
	Id              int64        `json:"applicationId"`
	Name            string       `json:"applicationName"`
	Repo            string       `json:"repository"`
	Members         []MemberData `json:"members"`
	Owner           string       `json:"owner"`
	Icon            string       `json:"icon"`
	LatestIteration int64        `json:"latestIteration"`
}

type MemberData struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetUserAllApplication(c *context.Context) []byte {
	userName := c.Params("username")
	apps := db.GetApplicationByUser(userName)
	// agg user
	userQueryMap := make(map[string]bool)
	userMap := make(map[string]*db.User)
	var userNames []string
	for _,app := range apps{
		for _, user := range app.Users {
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
	var applicationDatas []ApplicationData
	for _, app := range apps {
		applicationData := ApplicationData{
			Id: app.ID,
			Name: app.AppName,
			Repo: app.RepoUrl,
			Owner: app.Owner,
			Icon: app.AvatarURL,
			LatestIteration: app.LatestIteration,
		}
		var members []MemberData
		for _, user := range app.Users {
			members = append(members, MemberData{user, userMap[user].Avatar})
		}
		applicationData.Members = members
		applicationDatas = append(applicationDatas, applicationData)
	}

	data, _ := json.Marshal(applicationDatas)
	return data
}

func GetAllUsers(c *context.Context) []byte{
	names, _ := db.GetAllUser()
	data, _ := json.Marshal(names)
	return data
}

func GetAppAllBranch(c *context.Context) []byte {
	appOwner := c.Query("appOwner")
	appName := c.Query("appName")

	branches, err := invokerImpl.InvokeQueryAppAllBranchesService(appOwner, appName)
	if err!=nil {
		return []byte(fmt.Sprintf("[]"))
	}
	data, _ := json.Marshal(branches)
	return data
}

func GetAppAllWhiteBranch(c *context.Context) []byte {
	appOwner := c.Query("appOwner")
	appName := c.Query("appName")

	branches, err := invokerImpl.InvokeQueryAppAllBranchesService(appOwner, appName)
	if err!=nil {
		return []byte(fmt.Sprintf("[]"))
	}
	blackBranches := db.GetAppAllBlackBranch(appOwner, appName)
	var whiteBranches []string
	for _, ab := range branches {
		contain := false
		for _, bb := range blackBranches {
			if ab==bb{
				contain = true
			}
		}
		if !contain{
			whiteBranches = append(whiteBranches, ab)
		}
	}
	date, _ := json.Marshal(whiteBranches)
	return date
}


func NewApplication(c *context.Context, form form.Application) []byte {
	userName := c.Query("appOwner")
	repoName := c.Query("applicationName")
	members := c.Query("authMembers")
	description := c.Query("appDescription")
	appAuth := c.Query("appAuthScope")
	configBytes := []byte(c.Query("configs"))

	appImage := c.Query("appImage")
	appRegistry := c.Query("appRegistry")
	appTrace := c.Query("appTraceImage")
	appDataBase := c.Query("appDataBase")
	appBusinessDomain := c.Query("appBusinessDomain")
	imgBase64 := c.Query("image")[22:]
	appImportRepoUrl := c.Query("yamaHubUrl")

	if db.ApplicationIsExist(userName, repoName) {
		return []byte(fmt.Sprintf("user: %s already has application: %s", userName, repoName))
	}
	newApplication := db.Application{Owner: userName, AppName: repoName, Description: description, ApplicationAuth: appAuth, ApplicationImage: appImage,
		ApplicationRegistry: appRegistry, ApplicationTrace: appTrace, ApplicationDataBase: appDataBase, ApplicationBusinessDomain: appBusinessDomain,
		ApplicationDomainName: fmt.Sprintf("%s.%s", userName, repoName),Users:  strings.Split(members,","),
	}
	if err := buildUpApplicationIcon(&newApplication, imgBase64); err!=nil {
		return []byte(err.Error())
	}

	// 1. set up git repo(invoke grpc in YamaHub)
	if appImportRepoUrl=="" {
		appCreateRes, err := invokerImpl.InvokeCreateApplicationService(invokerarg.CreateApplicationOptions{
			Description: description,
			UserName:    userName,
			/**
			1:公司内部
			2:团队内部
			3:个人
			*/
			IsPrivate: appAuth == "3",
			AutoInit:  false,
			RepoName:  repoName,
		})
		if err != nil {
			return []byte(fmt.Errorf("failed to create application:%s, %s", repoName, err).Error())
		}
		newApplication.RepoUrl = appCreateRes.CloneUrl
	} else {
		newApplication.RepoUrl = appImportRepoUrl
	}
	// 4. set up database
	if err := buildUpApplicationDataBase(&newApplication); err!=nil {
		return []byte(fmt.Errorf("error while buildup application database, err %s", err).Error())
	}
	// 3. set up domain
	if err := buildUpApplicationDomain(&newApplication); err!=nil {
		return []byte(fmt.Errorf("error while buildup application domain, err %s", err).Error())
	}
	// 2. set up network,
	network,err := db.CreateApplicationNetWork(util.GenerateRandomStringWithSuffix(15,""),userName, repoName)
	if err != nil {
		return []byte(fmt.Errorf("error while create application network:%s, %s, %s",userName, repoName, err).Error())
	}
	newApplication.NetWorkIP = network.IPRange
	newApplication.NetWorkName = network.NetWorkName
	// 5. set up config
	if err := buildUpApplicationConfigWithDataBaseType(&newApplication, configBytes, appDataBase); err!=nil {
		return []byte(fmt.Errorf("error while buildup application config, err %s", err).Error())
	}
	// 6. db insert
	if err := db.InsertApplication(&newApplication); err!=nil {
		return []byte(fmt.Errorf("error while insert application domain, err %s", err).Error())
	}
	return nil
}

func buildUpApplicationIcon(application *db.Application, imgSrc string) error {
	if img,err := base64.StdEncoding.DecodeString(imgSrc); err != nil {
		return err
	} else {
		name := fmt.Sprintf("%s-%s",application.Owner, application.AppName)
		success,err := PutImage(name, img)
		if !success || err!=nil {
			return err
		}
		application.AvatarURL = fmt.Sprintf("https://3levelimage.oss-cn-hangzhou.aliyuncs.com/%s.png", name)
		return nil
	}
}

func buildUpApplicationDomain(application *db.Application) error {
	if err := ApplyApplicationDomain("dev", application.ApplicationDomainName); err != nil {
		return err
	}
	if err := ApplyApplicationDomain("stable", application.ApplicationDomainName); err != nil {
		return err
	}
	if err := ApplyApplicationDomain("test", application.ApplicationDomainName); err != nil {
		return err
	}
	if err := ApplyApplicationDomain("pre", application.ApplicationDomainName); err != nil {
		return err
	}
	if err := ApplyApplicationDomain("", application.ApplicationDomainName); err != nil {
		return err
	}
	return nil
}

func buildUpApplicationDataBase(application *db.Application) error {
	scheme := "%s_%s_%s"
	if err := resource.CreateDataBaseInGlobalDataBase(application.ApplicationDataBase, fmt.Sprintf(scheme,application.Owner, application.AppName, "dev")); err!=nil {
		return err
	}
	if err := resource.CreateDataBaseInGlobalDataBase(application.ApplicationDataBase, fmt.Sprintf(scheme,application.Owner, application.AppName, "stable")); err!=nil {
		return err
	}
	if err := resource.CreateDataBaseInGlobalDataBase(application.ApplicationDataBase, fmt.Sprintf(scheme,application.Owner, application.AppName, "test")); err!=nil {
		return err
	}
	if err := resource.CreateDataBaseInGlobalDataBase(application.ApplicationDataBase, fmt.Sprintf(scheme,application.Owner, application.AppName, "pre")); err!=nil {
		return err
	}
	if err := resource.CreateDataBaseInGlobalDataBase(application.ApplicationDataBase, fmt.Sprintf(scheme,application.Owner, application.AppName, "prod")); err!=nil {
		return err
	}
	return nil
}



func buildUpApplicationConfigWithDataBaseType(application *db.Application, configs []byte, appDataBase string) error{
	var configItems []db.ConfigItem
	if err := json.Unmarshal(configs, &configItems); err!=nil {
		return err
	}
	cfg := db.Config{ConfigItems: configItems}
	pattern := (&cfg).GetConfigItem(config.JAVA_SPRING_DYNAMIC_CONFIG.DATABASE_URL).(string)

	dataBaseIP := ""
	dataBaseType := ""
	if strings.Contains(strings.ToLower(appDataBase), "mysql") {
		dataBaseIP = resource.GLOBAL_MYSQL_IP
		dataBaseType = "mysql"
	} else if appDataBase == "OceanBase" {
	} else if appDataBase == "Oracle" {
	} else if appDataBase == "Sqlite3" {
	}
	cfg.SetConfigItem(config.JAVA_SPRING_DYNAMIC_CONFIG.ZIPKIN_URL, fmt.Sprintf("http://%s:9411", resource.GLOBAL_ZIPKIN_IP))

	cfg.SetConfigItem(config.JAVA_SPRING_DYNAMIC_CONFIG.DATABASE_URL, fmt.Sprintf(pattern, dataBaseType, dataBaseIP, fmt.Sprintf("%s_%s_dev", application.Owner, application.AppName)))
	data, _:= json.Marshal(cfg)
	application.DevConfig = string(data)

	cfg.SetConfigItem(config.JAVA_SPRING_DYNAMIC_CONFIG.DATABASE_URL, fmt.Sprintf(pattern, dataBaseType, dataBaseIP, fmt.Sprintf("%s_%s_test", application.Owner, application.AppName)))
	data, _= json.Marshal(cfg)
	application.TestConfig = string(data)

	cfg.SetConfigItem(config.JAVA_SPRING_DYNAMIC_CONFIG.DATABASE_URL, fmt.Sprintf(pattern, dataBaseType, dataBaseIP, fmt.Sprintf("%s_%s_stable", application.Owner, application.AppName)))
	data, _= json.Marshal(cfg)
	application.StableConfig = string(data)

	cfg.SetConfigItem(config.JAVA_SPRING_DYNAMIC_CONFIG.DATABASE_URL, fmt.Sprintf(pattern, dataBaseType, dataBaseIP, fmt.Sprintf("%s_%s_pre", application.Owner, application.AppName)))
	data, _= json.Marshal(cfg)
	application.PreConfig = string(data)

	cfg.SetConfigItem(config.JAVA_SPRING_DYNAMIC_CONFIG.DATABASE_URL, fmt.Sprintf(pattern, dataBaseType, dataBaseIP, fmt.Sprintf("%s_%s_prod", application.Owner, application.AppName)))
	data, _= json.Marshal(cfg)
	application.ProdConfig = string(data)

	return nil
}