package bean

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"yama.io/yamaIterativeE/internal/deploy/container/yamapouch/command"
)

type RollBackBean struct {
	CompileBean
	DeployBean
}


// appOwner appName repo branch execPath env serverName serverIP iterId(string) commitId logPath
func (rb *RollBackBean) Execute(stringArgs []string, env *map[string]interface{}) error{
	// 1.parse arg
	if len(stringArgs) != 11{
		return fmt.Errorf("arguement error")
	}
	if stringArgs[9] != "" {
		iterId, err := strconv.Atoi(stringArgs[8])
		if err != nil {
			return err
		}
		err = rb.clone(stringArgs[1], stringArgs[2], stringArgs[3], stringArgs[4], stringArgs[9], stringArgs[10])
		if err != nil {
			return err
		}
		err = rb.flushConfig(stringArgs[1], stringArgs[4], stringArgs[5], stringArgs[6], stringArgs[7], int64(iterId))
		if err != nil {
			return err
		}
		err = rb.mavenInstall(stringArgs[1], stringArgs[4], stringArgs[10])
		if err != nil {
			return err
		}
		err = rb.deploy(stringArgs[1], stringArgs[4],stringArgs[6], stringArgs[7])
		return err
	} else if stringArgs[9] == ""{
		command.StopContainer(stringArgs[6])
		command.StartAppInContainer(stringArgs[6], "top -b")
	}
	return nil
}

func (rb *RollBackBean) clone(appName, repoPath, branchName, execPath, commitId, logPath string) error {
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

	cmd = exec.Command("git", "checkout", commitId)
	cmd.Dir = fmt.Sprintf("%s/%s", execPath, appName)
	if logPath != "" {
		log, _ := os.Open(logPath)
		cmd.Stdout = log
		cmd.Stderr = log
	}
	err = cmd.Run()

	return err
}

