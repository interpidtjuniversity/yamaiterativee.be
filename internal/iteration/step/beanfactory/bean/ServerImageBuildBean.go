package bean

import (
	"fmt"
	"time"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/deploy/container/yamapouch/command"
)

type ServerImageBuildBean struct {
	Bean
}

// appType logPath
func (sibb *ServerImageBuildBean) Execute(stringArgs []string, env *map[string]interface{}) error{
	if len(stringArgs)!=2 {
		return fmt.Errorf("arguement error")
	}
	sibb.build(stringArgs[0], env)
	return nil
}

func (sibb *ServerImageBuildBean) build(appType string, env *map[string]interface{}) error {
	var ip string
	switch appType {
	case db.JAVA_SPRING:
		ip,_ = newJavaApplicationServer((*env)["serverName"].(string), (*env)["networkName"].(string))
		break
	default:
		return fmt.Errorf("error while create server, error: unsupport appType %s", appType)
	}
	(*env)["serverIP"] = ip
	return nil
}

func newJavaApplicationServer(containerName, network string) (string, error){
	time.Sleep(time.Duration(5)*time.Second)
	ip, err := command.CreateContainer(containerName, network, "", "JavaImage", "top -b")
	return ip, err
}
