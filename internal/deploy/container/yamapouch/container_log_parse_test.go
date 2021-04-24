package yamapouch

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"yama.io/yamaIterativeE/internal/deploy/container/yamapouch/command"
)


func Test_ListContainer(t *testing.T) {
	containers, err := command.ListContainer()
	if err != nil {

	} else {
		fmt.Print(containers)
	}
}

func Test_CreateNetWork(t *testing.T) {
	err := command.CreateNetWork("basebridge","192.168.10.1/24","bridge")
	assert.Nil(t, err)
}

func Test_CreateContainer(t *testing.T) {
	ip, err := command.CreateContainer("spring-prod10000", "basebridge", "9000:8080", "JavaImage", "top -b")
	assert.Nil(t, err)
	fmt.Print(ip)
}

func Test_EnhanceContainer(t *testing.T) {
	err := command.EnhanceContainer("spring-dev", "/root", "/root/demo.jar")
	assert.Nil(t, err)
}

func Test_DeployAppInContainer(t *testing.T) {
	err := command.DeployAppInContainer("spring-dev", "/jdk1.8.0_281/bin/java","-jar /root/demo.jar")
	assert.Nil(t, err)
}

func Test_ReDeployAppInContainer(t *testing.T) {
	err := command.ReDeployAppInContainer("spring-dev", "/jdk1.8.0_281/bin/java","-jar /root/demo.jar")
	assert.Nil(t, err)
}

// if container type is MysqlImage, then these commands will start mysql-server in container
func Test_DeployMysqlInContainer(t *testing.T) {
	err := command.ExecuteCommandInContainer("mysql-pre", "/bin/chmod","666","/dev/null")
	assert.Nil(t, err)
	time.Sleep(time.Second*2)
	err = command.ExecuteCommandInContainer("mysql-pre", "/usr/sbin/service","mysql", "start")
	assert.Nil(t, err)
}

func Test_DeployMysqlBatchCommandOnceInContainer(t *testing.T) {
	err := command.ExecuteBatchCommandOnceInContainer("mysql-pre", "/bin/chmod 666 /dev/null#/usr/sbin/service mysql start")
	assert.Nil(t, err)
}


func Test_StopContainer(t *testing.T) {
	err := command.StopContainer("spring-dev")
	assert.Nil(t, err)
}

func Test_RemoveContainer(t *testing.T) {
	err := command.RemoveContainer("spring-dev")
	assert.Nil(t, err)
}

func Test_StartAppInContainer(t *testing.T) {
	err := command.StartAppInContainer("spring-prod10000", "top -b")
	assert.Nil(t, err)
}

func Test_ContainerAppLogs(t *testing.T) {
	out,err := command.GetApplicationLog("spring-prod10000", "/root/miniselfop/start-up.log")
	assert.Nil(t, err)

	fmt.Print(string(out))
}

func Test_ContainerAppDefaultLogs(t *testing.T) {
	out, err := command.GetApplicationLog("spring-prod10000","")
	assert.Nil(t, err)

	fmt.Print(string(out))
}

func Test_RemoveNetWork(t *testing.T) {
	err := command.RemoveNetWork("basebridge")
	assert.Nil(t, err)
}