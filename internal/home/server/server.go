package server

import (
	"fmt"
	"time"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/deploy/container/yamapouch/command"
	"yama.io/yamaIterativeE/internal/util"
)

const (
	// root-miniselfop-cl5664dwed-dev
	// system-mysql-zcnelw05-dev
	// userName-appName-id-env
	SERVER_NAME = "%s-%s-%s-%s"
)


func NewServer(c *context.Context) []byte {
	appType := db.AppType(c.QueryInt("appType"))
	iterBranch := c.Query("iterBranch")
	appOwner := c.Query("appOwner")
	appName := c.Query("appName")
	iterId := c.QueryInt64("iterId")
	owner := c.Query("owner")
	serverType := db.ServerType(c.QueryInt("serverType"))

	newServer := &db.Server{AppOwner: appOwner, AppName: appName, IterationId: iterId, Owner: owner, AppType: appType,
		Branch: iterBranch, Type: serverType, State: db.APPLYING, CreatedTime: time.Now().Format("2006-01-01 15:04:05"),
		Name: util.GenerateRandomStringWithSuffix(20, fmt.Sprintf("%s.%s.%s", appOwner, appName, serverType.ToString())),
	}

	networkName, _ := db.GetApplicationNetworkByOwnerAndRepo(appOwner, appName)
	newServer.NetWork = networkName
	_, err := db.InsertServer(newServer)
	if err != nil {
		return []byte(fmt.Errorf("error while create server, error: %s", err).Error())
	}

	// 1. invoke New...Server by serverType
	var ip string
	switch appType {
	case db.JAVA_SPRING:
		ip,_ = NewJavaApplicationServer(newServer.Name, networkName)
	default:
		return []byte(fmt.Errorf("error while create server, error: unsupport appType %s", appType.ToString()).Error())
	}
	// 2. update server ip
	db.UpdateServerAfterApply(newServer.Name, ip)

	return nil
}

func NewJavaApplicationServer(containerName, network string) (string, error){
	time.Sleep(time.Duration(5)*time.Second)
	ip, err := command.CreateContainer(containerName, network, "", "JavaImage", "top -b")
	return ip, err
}

func NewMysqlServer(containerName, network, env string) (string, error){
	if containerName == "" {
		containerName = fmt.Sprintf(SERVER_NAME, "system", "mysql", util.GenerateRandomStringWithSuffix(10, ""), env)
	}
	ip, err := command.CreateContainer(containerName, network, "", "MysqlImage", "top -b")
	if err != nil || ip == ""{
		return ip, err
	}
	err = command.ExecuteBatchCommandOnceInContainer(containerName, "/bin/chmod 666 /dev/null#/usr/sbin/service mysql start")
	return ip, err
}

func NewConsulServer(containerName, network, env string) (string, error){
	if containerName == "" {
		containerName = fmt.Sprintf(SERVER_NAME, "system", "consul", util.GenerateRandomStringWithSuffix(10, ""), env)
	}
	ip, err := command.CreateContainer(containerName, network, "", "ConsulImage", "top -b")
	if err != nil || ip == ""{
		return ip, err
	}
	err = command.ExecuteBatchCommandOnceInContainer(containerName, fmt.Sprintf("/bin/chmod 777 standalone-consul.sh#./standalone-consul.sh %s %s %s", containerName, ip, containerName))
	return ip, err
}

func NewZipkinServer(containerName, network, env string) (string, error){
	if containerName == "" {
		containerName = fmt.Sprintf(SERVER_NAME, "system", "zipkin", util.GenerateRandomStringWithSuffix(10, ""), env)
	}
	ip, err := command.CreateContainer(containerName, network, "", "ZipkinImage", "top -b")
	if err != nil || ip == ""{
		return ip, err
	}
	err = command.ExecuteBatchCommandOnceInContainer(containerName, "/bin/chmod 777 standalone-zipkin.sh#./standalone-zipkin.sh")
	return ip, err
}
