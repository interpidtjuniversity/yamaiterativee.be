package server

import (
	"fmt"
	"yama.io/yamaIterativeE/internal/deploy/container/yamapouch/command"
	"yama.io/yamaIterativeE/internal/util"
)

const (
	// root-miniselfop-cl5664dwed-dev
	// system-mysql-zcnelw05-dev
	// userName-appName-id-env
	SERVER_NAME = "%s-%s-%s-%s"
)

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
