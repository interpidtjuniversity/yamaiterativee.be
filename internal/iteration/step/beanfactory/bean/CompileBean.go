package bean

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"yama.io/yamaIterativeE/internal/db"
	config2 "yama.io/yamaIterativeE/internal/home/config"
)

const (
	PropertiesPath = "%s/%s/src/main/resources/application.properties"
	GenerateSourcePath = "%s/%s/target/"
)

type CompileBean struct {
	Bean
}
//   0        1     2      3      4      5      6         7         8            9
//appOwner appName repo branch execPath env serverName serverIP iterId(string) logPath

func (cb *CompileBean) Execute(stringArgs []string, env *map[string]interface{}) error {
	// 1.parse arg
	if len(stringArgs) != 10{
		return fmt.Errorf("arguement error")
	}
	iterId,err := strconv.Atoi(stringArgs[8])
	if err!=nil {
		return err
	}
	err = cb.clone(stringArgs[2], stringArgs[3], stringArgs[4], stringArgs[9])
	if err!=nil {
		return err
	}
	err = cb.flushConfig(stringArgs[1], stringArgs[4], stringArgs[5], stringArgs[6], stringArgs[7], int64(iterId))
	if err!=nil {
		return err
	}
	err = cb.mavenInstall(stringArgs[1], stringArgs[4], stringArgs[9])

	return err
}

func (cb *CompileBean) clone(repoPath, branchName, execPath, logPath string) error {
	if err := os.MkdirAll(execPath, os.ModePerm); err != nil {
		return fmt.Errorf("error while execute git clone, err: %s", err)
	}

	cmd := exec.Command("git", "clone", "-b", branchName, repoPath)
	cmd.Dir = execPath
	if logPath != "" {
		log, _ := os.Open(logPath)
		defer log.Close()
		cmd.Stdout = log
		cmd.Stderr = log
	}
	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("error while execute git clone, err: %s", err)
	}
	return nil
}

func (cb *CompileBean) flushConfig(appName, execPath, serverEnv, serverName, serverIp string, iterId int64) error {
	var configString string
	config := new(db.Config)
	switch serverEnv {
	case "dev":
		configString = db.GetIterationDevConfig(iterId)
		json.Unmarshal([]byte(configString), config)
		config.SetConfigItem(config2.JAVA_SPRING_DYNAMIC_CONFIG.HOST_TAGS, "dev")
		break
	case "stable":
		configString = db.GetIterationStableConfig(iterId)
		json.Unmarshal([]byte(configString), config)
		config.SetConfigItem(config2.JAVA_SPRING_DYNAMIC_CONFIG.HOST_TAGS, "stable")
		break
	case "test":
		configString = db.GetIterationTestConfig(iterId)
		json.Unmarshal([]byte(configString), config)
		config.SetConfigItem(config2.JAVA_SPRING_DYNAMIC_CONFIG.HOST_TAGS, "test")
		break
	case "pre":
		configString = db.GetIterationPreConfig(iterId)
		json.Unmarshal([]byte(configString), config)
		config.SetConfigItem(config2.JAVA_SPRING_DYNAMIC_CONFIG.HOST_TAGS, "pre")
		break
	case "prod":
		configString = db.GetIterationProdConfig(iterId)
		json.Unmarshal([]byte(configString), config)
		config.SetConfigItem(config2.JAVA_SPRING_DYNAMIC_CONFIG.HOST_TAGS, "prod")
	}
	config.SetConfigItem(config2.JAVA_SPRING_DYNAMIC_CONFIG.INSTANCE_ID, fmt.Sprintf("instance-%s",serverName))
	config.SetConfigItem(config2.JAVA_SPRING_DYNAMIC_CONFIG.HOST_NAME, serverIp)

	// write properties
	properties ,err := os.OpenFile(fmt.Sprintf(PropertiesPath, execPath, appName), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err!=nil {
		return err
	}
	defer properties.Close()
	write := bufio.NewWriter(properties)
	for _, item := range config.ConfigItems {
		write.WriteString(fmt.Sprintf("%s=%v\n", item.Key, item.Value))
	}
	return write.Flush()
	return nil
}

func (cb *CompileBean) mavenInstall(appName, execPath, logPath string) error {

	cmd := exec.Command("mvn", "install")
	cmd.Dir = fmt.Sprintf("%s/%s", execPath, appName)
	if logPath != "" {
		log, _ := os.Open(logPath)
		defer log.Close()
		cmd.Stdout = log
		cmd.Stderr = log
	}
	err := cmd.Run()

	if err != nil {
		return err
	}
	return nil
}
