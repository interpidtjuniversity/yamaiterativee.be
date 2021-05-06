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
	"yama.io/yamaIterativeE/internal/resource"
	"yama.io/yamaIterativeE/internal/util"
)

func GetAllUsers(c *context.Context) []byte{
	names, _ := db.GetAllUser()
	data, _ := json.Marshal(names)
	return data
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
	newApplication := db.Application{Owner: userName, AppName: repoName, Description: description, ApplicationAuth: appAuth, ApplicationImage: appImage,
		ApplicationRegistry: appRegistry, ApplicationTrace: appTrace, ApplicationDataBase: appDataBase, ApplicationBusinessDomain: appBusinessDomain,
		ApplicationDomainName: fmt.Sprintf("%s.%s", userName, repoName),Users:  strings.Split(members,","),
	}
	if err := buildUpApplicationIcon(&newApplication, imgBase64); err!=nil {
		return []byte(err.Error())
	}

	// 1. set up git repo(invoke grpc in YamaHub)
	appCreateRes, err :=invokerImpl.InvokeCreateApplicationService(invokerarg.CreateApplicationOptions{
		Description: description ,
		UserName: userName,
		/**
			1:公司内部
			2:团队内部
			3:个人
		*/
		IsPrivate: appAuth=="3",
		AutoInit:  false,
		RepoName:  repoName,
	})
	if err!=nil {
		//return []byte(fmt.Errorf("failed to create application:%s, %s", repoName, err).Error())
	}
	newApplication.RepoUrl = appCreateRes.CloneUrl
	// 2. set up network,
	network,err := db.CreateApplicationNetWork(util.GenerateRandomString(15,""),userName, repoName)
	if err != nil {
		return []byte(fmt.Errorf("error while create application network:%s, %s, %s",userName, repoName, err).Error())
	}
	newApplication.NetWorkIP = network.IPRange
	newApplication.NetWorkName = network.NetWorkName
	// 3. set up domain
	if err := buildUpApplicationDomain(&newApplication); err!=nil {
		return []byte(fmt.Errorf("error while buildup application domain, err %s", err).Error())
	}
	// 4. set up database
	if err := buildUpApplicationDataBase(&newApplication); err!=nil {
		return []byte(fmt.Errorf("error while buildup application database, err %s", err).Error())
	}
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
		success,err := db.PutImage(name, img)
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
	config := db.Config{ConfigItems: configItems}
	pattern := (&config).GetConfigItem(JAVA_SPRING_DYNAMIC_CONFIG.DATABASE_URL).(string)

	dataBaseIP := ""
	dataBaseType := ""
	if appDataBase == "Mysql" {
		dataBaseIP = resource.GLOBAL_MYSQL_IP
		dataBaseType = "mysql"
	} else if appDataBase == "OceanBase" {
	} else if appDataBase == "Oracle" {
	} else if appDataBase == "Sqlite3" {
	}

	config.SetConfigItem(JAVA_SPRING_DYNAMIC_CONFIG.DATABASE_URL, fmt.Sprintf(pattern, dataBaseType, dataBaseIP, fmt.Sprintf("%s_%s_dev", application.Owner, application.AppName)))
	data, _:= json.Marshal(config)
	application.DevConfig = string(data)

	config.SetConfigItem(JAVA_SPRING_DYNAMIC_CONFIG.DATABASE_URL, fmt.Sprintf(pattern, dataBaseType, dataBaseIP, fmt.Sprintf("%s_%s_test", application.Owner, application.AppName)))
	data, _= json.Marshal(config)
	application.TestConfig = string(data)

	config.SetConfigItem(JAVA_SPRING_DYNAMIC_CONFIG.DATABASE_URL, fmt.Sprintf(pattern, dataBaseType, dataBaseIP, fmt.Sprintf("%s_%s_stable", application.Owner, application.AppName)))
	data, _= json.Marshal(config)
	application.StableConfig = string(data)

	config.SetConfigItem(JAVA_SPRING_DYNAMIC_CONFIG.DATABASE_URL, fmt.Sprintf(pattern, dataBaseType, dataBaseIP, fmt.Sprintf("%s_%s_pre", application.Owner, application.AppName)))
	data, _= json.Marshal(config)
	application.PreConfig = string(data)

	config.SetConfigItem(JAVA_SPRING_DYNAMIC_CONFIG.DATABASE_URL, fmt.Sprintf(pattern, dataBaseType, dataBaseIP, fmt.Sprintf("%s_%s_prod", application.Owner, application.AppName)))
	data, _= json.Marshal(config)
	application.ProdConfig = string(data)

	return nil
}