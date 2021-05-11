package server

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/deploy/build"
	"yama.io/yamaIterativeE/internal/deploy/container/yamapouch/command"
	"yama.io/yamaIterativeE/internal/util"
)

func DeployAppInServer(c *context.Context) []byte {
	serverName := c.Query("serverName")
	serverIp := c.Query("serverIP")
	serverEnv := c.Query("serverEnv")
	appOwner := c.Query("appOwner")
	appName := c.Query("appName")
	deployBranch := c.Query("deployBranch")

	//0. check if can deploy
	last, _ := db.GetDeployIdByServerName(serverName)
	if last != "" {
		return []byte(fmt.Sprintf("error while execute deploy, server: %s is already in deploying", serverName))
	}

	//1. update server state
	deployId := util.GenerateRandomStringWithSuffix(10,"")
	tmpDir := util.GenerateRandomStringWithPrefix(15,"")
	logPath := fmt.Sprintf(build.BUILD_LOG_PATH, tmpDir)
	generateSourceDir := fmt.Sprintf(build.BUILD_RESOURCE_PATH, tmpDir, appName)
	db.UpdateServerDeployId(serverName, deployId, db.DEPLOYING)
	db.InsertDeploy(&db.Deploy{ServerName: serverName, DeployId: deployId, DeployLogPath: logPath})

	go func() {
		//2. git clone -b xxx xxx
		repoPath := db.GetApplicationRepoByOwnerAndRepo(appOwner, appName)
		out, mvnDir, _ := build.Clone(repoPath, deployBranch, tmpDir, appName)
		writeLog(out, logPath)
		//3. update yml

		//4. mvn install
		out, _ = build.MavenCompile(mvnDir)
		writeLog(out, logPath)
		//5. copy generate source to container
		sources, _ := ioutil.ReadDir(generateSourceDir)
		executeable := ""
		for _, source := range sources{
			if strings.HasSuffix(source.Name(), "jar") {
				executeable = source.Name()
				command.EnhanceContainer(serverName, "/root", executeable)
				break
			}
		}
		//6. execute
		if executeable != "" {
			command.DeployAppInContainer(serverName, "/jdk1.8.0_281/bin/java", fmt.Sprintf("-jar /root/%s", executeable))
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
			if ping {
				db.UpdateServerDeployId(serverName, "", db.IDLE)
			} else {
				db.UpdateServerDeployId(serverName, "", db.STOPPED)
			}
		}

	}()

	return nil
}

func writeLog(data []byte, path string) error{
	log ,err := os.OpenFile(path,os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err!=nil {
		return err
	}
	defer log.Close()
	write := bufio.NewWriter(log)
	write.Write(data)
	return write.Flush()
}

// ip:8088/actuator/health
func isDeployFinish(ip string) bool {
	resp, err := http.Get(fmt.Sprintf("%s:8088/actuator/health", ip))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return true
	}
	return false
}