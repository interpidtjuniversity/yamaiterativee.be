package server

import (
	"fmt"
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
	appType := c.Query("appType")
	iterBranch := c.Query("iterBranch")
	appOwner := c.Query("appOwner")
	appName := c.Query("appName")
	iterId := c.QueryInt64("iterId")
	owner := c.Query("owner")
	serverType := db.ServerType(c.QueryInt("serverType"))

	newServer := &db.Server{AppOwner: appOwner, AppName: appName, IterationId: iterId, Owner: owner, AppType: appType,
		Branch: iterBranch, Type: serverType,
	}
	// 1. query application network
	// 2. insert server instance
	db.InsertServer(newServer)
	// 3. invoke New...Server by serverType

	return nil
}

func NewJavaApplicationServer(userName, appName, network, env string) (string, error){
	containerName := fmt.Sprintf(SERVER_NAME, userName, appName, util.GenerateRandomString(10,""), env)
	ip, err := command.CreateContainer(containerName, network, "", "JavaImage", "top -b")
	return ip, err
}

func NewMysqlServer(containerName, network, env string) (string, error){
	if containerName == "" {
		containerName = fmt.Sprintf(SERVER_NAME, "system", "mysql", util.GenerateRandomString(10, ""), env)
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
		containerName = fmt.Sprintf(SERVER_NAME, "system", "consul", util.GenerateRandomString(10, ""), env)
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
		containerName = fmt.Sprintf(SERVER_NAME, "system", "zipkin", util.GenerateRandomString(10, ""), env)
	}
	ip, err := command.CreateContainer(containerName, network, "", "ZipkinImage", "top -b")
	if err != nil || ip == ""{
		return ip, err
	}
	err = command.ExecuteBatchCommandOnceInContainer(containerName, "/bin/chmod 777 standalone-zipkin.sh#./standalone-zipkin.sh")
	return ip, err
}
