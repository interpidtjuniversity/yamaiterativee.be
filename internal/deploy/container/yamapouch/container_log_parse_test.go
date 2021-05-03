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
	err := command.CreateNetWork("bridasdg","10.0.9.0/24","bridge")
	assert.Nil(t, err)
}

func Test_CreateContainer(t *testing.T) {
	ip, err := command.CreateContainer("zipkin-dev", "test", "", "ZipkinImage", "top -b")
	assert.Nil(t, err)
	fmt.Print(ip)
}

func Test_EnhanceContainer(t *testing.T) {
	err := command.EnhanceContainer("spring-prod10000", "/root", "/root/demo.jar")
	assert.Nil(t, err)
}

func Test_DeployAppInContainer(t *testing.T) {
	err := command.DeployAppInContainer("spring-prod10000", "/jdk1.8.0_281/bin/java","-jar /root/demo.jar")
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

func Test_DeployConsulInContainer(t *testing.T) {
	//nohup consul agent -server -ui -bootstrap-expect=1 -data-dir=/data/consul -node=consul-dev -client=0.0.0.0 -bind=192.168.200.2 -datacenter=consul-dev &
	err := command.ExecuteBatchCommandOnceInContainer("consul-stable", "/bin/chmod 777 standalone-consul.sh#./standalone-consul.sh consul-stable 192.168.10.8 consul-stable")
	assert.Nil(t, err)
}

func Test_DeployZipkinInContainer(t *testing.T) {
	// nohup /jdk1.8.0_281/bin/java -jar zipkin.jar >> zipkin.log 2>&1 &
	err := command.ExecuteBatchCommandOnceInContainer("zipkin-dev", "/bin/chmod 777 standalone-zipkin.sh#./standalone-zipkin.sh")
	assert.Nil(t, err)
}

func Test_DeployMysqlBatchCommandOnceInContainer(t *testing.T) {
	err := command.ExecuteBatchCommandOnceInContainer("mysql-dev", "/bin/chmod 666 /dev/null#/usr/sbin/service mysql start")
	assert.Nil(t, err)
}


func Test_StopContainer(t *testing.T) {
	err := command.StopContainer("spring-prod10000")
	assert.Nil(t, err)
}

func Test_RemoveContainer(t *testing.T) {
	err := command.RemoveContainer("spring-prod10000")
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

// for consul
// ./pouch run -d --name consul-itg01 --net consul -p 8501:8500 ConsulImage top -b
// ./pouch run -d --name consul-dev --net consul -p 8500:8500 ConsulImage top -b
// consul agent -server -ui -bootstrap-expect=1 -data-dir=/data/consul -node=consul-dev -client=0.0.0.0 -bind=192.168.200.2 -datacenter=consul-dev

// for zipkin
// /jdk1.8.0_281/bin/java -jar zipkin.jar