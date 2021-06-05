package bean

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/deploy/container/yamapouch/command"
	"yama.io/yamaIterativeE/internal/util"
)

const (
	ServerDeploymentConstantsIP = "spring.cloud.consul.discovery.hostname"
	ServerDeploymentConstantsName = "spring.cloud.consul.discovery.instance-id"
	ServerDeploymentConstantsTags = "spring.cloud.consul.discovery.tags"
	ServerDeploymentConstantsConsulHost = "spring.cloud.consul.host"
)

type DeployBean struct {
	Bean
}

// appName, execPath, serverName, serverIP
func (depb *DeployBean) Execute(stringArgs []string, env *map[string]interface{}) error{
	return depb.deploy(stringArgs[0], stringArgs[1], stringArgs[2], stringArgs[3])
}

func (depb *DeployBean) prepareSourcePath(execPath, appName string) string{
	return fmt.Sprintf(GenerateSourcePath, execPath, appName)
}

func (depb *DeployBean) deploy(appName, execPath, serverName, serverIp string) error {
	generateSourceDir := depb.prepareSourcePath(execPath, appName)
	return depb.doDeploy(generateSourceDir, serverName, serverIp)
}

func (*DeployBean) doDeploy (generateSourceDir, serverName, serverIp string) error {
	sources, _ := ioutil.ReadDir(generateSourceDir)
	executable := ""
	for _, source := range sources{
		if strings.HasSuffix(source.Name(), "jar") {
			executable = source.Name()
			command.EnhanceContainer(serverName, "/root", generateSourceDir+executable)
			break
		}
	}

	if executable != "" {
		server := db.GetServerStateAndTypeByName(serverName)
		if server.State == db.RUNNING {
			command.StopContainer(serverName)
			command.StartAppInContainer(serverName, "top -b")
		}
		command.DeployAppInContainer(serverName, "/jdk1.8.0_281/bin/java", fmt.Sprintf("-jar /root/%s --%s=%s --%s=%s --%s=%s --%s=%s", executable, ServerDeploymentConstantsIP, serverIp, ServerDeploymentConstantsName, fmt.Sprintf("instance-%s",serverName), ServerDeploymentConstantsTags, server.Type.ToString(), ServerDeploymentConstantsConsulHost, util.GetLocalIPv4Address()))
		db.UpdateServerAfterDeploy(serverName)
	}
	var counter int
	var success bool
	for true {
		if counter > 100{
			break
		}
		if isDeployFinish(serverIp) {
			success = true
			break
		}
		time.Sleep(time.Duration(5)*time.Second)
		counter++
	}
	if success{
		db.UpdateServerDeployId(serverName, "", db.RUNNING)
	} else {

	}


	return nil
}

// ip:8088/actuator/health
func isDeployFinish(ip string) bool {
	resp, err := http.Get(fmt.Sprintf("http://%s:8088/actuator/health", ip))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return true
	}
	return false
}
