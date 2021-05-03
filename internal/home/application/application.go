package application

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/form"
	"yama.io/yamaIterativeE/internal/grpc/invoke/invokerImpl"
	"yama.io/yamaIterativeE/internal/grpc/invoke/invokerarg"
)

func GetAllUsers(c *context.Context) []byte{
	names, _ := db.GetAllUser()
	data, _ := json.Marshal(names)
	return data
}

func NewApplication(c *context.Context, form form.Application) []byte {
	userName := c.Query("appOwner")
	RepoName := c.Query("applicationName")
	members := c.Query("authMembers")
	description := c.Query("appDescription")
	appAuth := c.Query("appAuthScope")
	newApplication := db.Application{}
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
		AutoInit: false,
		RepoName: RepoName,
	})
	if err!=nil {
		return []byte(fmt.Errorf("failed to create application:%s, %s",RepoName, err).Error())
	}
	newApplication.Owner = appCreateRes.Owner
	newApplication.RepoUrl = appCreateRes.CloneUrl
	newApplication.AppName = appCreateRes.RepoName
	// 2. set up network,
	network,err := db.CreateApplicationNetWork(fmt.Sprintf("%s_%s",userName,RepoName),userName, RepoName)
	if err != nil {
		return []byte(fmt.Errorf("failed to create application network:%s, %s, %s",userName, RepoName, err).Error())
	}
	newApplication.NetWork = network.IPRange
	// 3. set up domain
		/**
			default dev domain is instanceName.dev.appName.userName.manchestercity.ren
		    default stable domain is instanceName.stable.appName.userName.manchestercity.ren
			default test domain is instanceName.test.appName.userName.manchestercity.ren
			default pre domain is instanceName.pre.appName.userName.manchestercity.ren
			default prod domain is instanceName.prod.appName.userName.manchestercity.ren
		*/
		// to enable domain please click enable domain options after application is built
	// 4. set up resource
		/**
			although global resource had been initialized, by we still have to do following things
			1. create database in GLOBAL_MYSQL: app_dev, app_stable, app_test, app_pre, app_prod
		*/
	// 5. set up config

	newApplication.Users = strings.Split(members,",")

	return nil
}

func SetNewApplicationIcon(c *context.Context) []byte {
	image, handle, err :=c.Req.FormFile("image")
	if err != nil {
		return []byte(err.Error())
	}
	// 检查图片后缀
	ext := strings.ToLower(path.Ext(handle.Filename))
	if ext != ".jpg" && ext != ".png" {
		return []byte("只支持jpg/png图片上传")
	}
	// 保存图片
	os.Mkdir("./uploaded/", 0777)
	saveFile, err := os.OpenFile("./uploaded/" + handle.Filename, os.O_WRONLY|os.O_CREATE, 0666);
	io.Copy(saveFile, image)

	defer image.Close()
	defer saveFile.Close()
	return nil
}
