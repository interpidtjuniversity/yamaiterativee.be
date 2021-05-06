package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"yama.io/yamaIterativeE/internal/deploy/container/yamapouch/model"
)

func Run0(command string, args ...string) (error) {
	pouchDir,_ := os.Getwd()
	cmd := exec.Command(command, args...)
	cmd.Dir = fmt.Sprintf("%s/internal/deploy/container/yamapouch/command", pouchDir)
	err := cmd.Run()

	return err
}

func Run(command string, args ...string) (*bytes.Buffer, *bytes.Buffer, error) {
	bufOut := new(bytes.Buffer)
	bufErr := new(bytes.Buffer)
	pouchDir,_ := os.Getwd()
	cmd := exec.Command(command, args...)
	cmd.Dir = fmt.Sprintf("%s/internal/deploy/container/yamapouch/command", pouchDir)
	cmd.Stdout = bufOut
	cmd.Stderr = bufErr
	err := cmd.Run()

	return bufOut, bufErr, err
}
// network
// ./pouch network create --driver bridge --subnet 192.168.90.1/24 basebridge
func CreateNetWork(name, ipRange, driver string) error{
	_, bufErr, err := Run("./pouch", "network", "create", "--driver", driver, "--subnet",ipRange, name)
	if err != nil {
		return err
	}
	if bufErr.Len() > 0 {
		return fmt.Errorf("error while execute pouch network create, error: %s", bufErr.String())
	}
	return nil
}

// run
// ./pouch run -d --name spring-prod3000 -net basebridge -p 8100:8080 JavaImage top -b
func CreateContainer(name, netWork, portMapping, imageName, initCommand string) (string, error){
	var bufOut, bufErr *bytes.Buffer
	var err error
	if portMapping != "" {
		bufOut, bufErr, err = Run("./pouch", "run", "-d", "--name", name, "--net", netWork, "-p", portMapping, imageName, initCommand)
	} else {
		bufOut, bufErr, err = Run("./pouch", "run", "-d", "--name", name, "--net", netWork, imageName, initCommand)
	}
	if err != nil {
		return "", err
	}
	if bufErr.Len() > 0{
		return "", fmt.Errorf("error while execute pouch run, error: %s", bufErr.String())
	}
	// parse ip
	for {
		slice, err := bufOut.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return "", err
			}
		}
		slice = slice[:len(slice)-1]
		var ccf model.ContainerCreateInfo
		json.Unmarshal(slice, &ccf)
		if strings.HasPrefix(ccf.Msg, "ip") {
			return ccf.Msg[3:], nil
		}
	}
	return "", fmt.Errorf("error while execute pouch run, error: %s", "ip is not generate")
}

// ps
// ./pouch ps
func ListContainer() ([]model.ContainerInfo, error){
	bufOut, bufErr, err := Run("./pouch","ps")

	if err != nil {
		return nil, err
	}
	if bufErr.Len() > 0 {
		return nil , fmt.Errorf("error while execute pouch ps, error: %s", bufErr.String())
	}
	var line int
	var containers []model.ContainerInfo
	for {
		slice, err := bufOut.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		slice = slice[:len(slice)-1]
		if line > 0 {
			var fields [][]byte
			data := bytes.Split(slice, []byte(" "))
			for _,b := range data {
				if len(b) != 0{
					fields = append(fields, b)
				}
			}
			var portMapping []string
			err := json.Unmarshal(fields[6], &portMapping)
			if err != nil {

			}
			containerInfo := model.ContainerInfo{
				Id: string(fields[0]),
				Name: string(fields[1]),
				Pid: string(fields[2]),
				Status: string(fields[3]),
				Command: string(fields[4]),
				CreatedTime: string(fields[5]),
				PortMapping: portMapping,
				Ip: string(fields[7]),
				NetWorkName: string(fields[8]),
			}

			containers = append(containers, containerInfo)
		}
		line++
	}

	return containers, nil
}

// logs
/** ./pouch logs spring-prod10000
     this is log for init command ignore
 */
func GetContainerLog() {
}

// applogs
// ./pouch applogs spring-prod10000 /root/miniselfop/start-up.log
func GetApplicationLog(name, logPath string) ([]byte, error) {
	var buffOut, buffErr *bytes.Buffer
	var err error
	// start-up log is app default log
	if logPath == "" {
		buffOut, buffErr, err = Run("./pouch", "applogs", name, "/root/logs/log.log")
	}else {
		buffOut, buffErr, err = Run("./pouch", "applogs", name, logPath)
	}
	if err != nil {
		return nil, err
	}
	if buffErr.Len() > 0 {
		return nil, fmt.Errorf("error while execute pouch applog, error: %s", buffErr.String())
	}
	return buffOut.Bytes(), nil
}

// stop
// ./pouch stop spring-prod10000
func StopContainer(name string) error {
	_, _, err := Run("./pouch", "stop", name)
	return err
}

// rm
// ./pouch rm spring-prod10000
func RemoveContainer(name string) error {
	_, _, err := Run("./pouch", "rm", name)
	return err
}

// deploy
// ./pouch deploy --name spring-prod10000 /jdk1.8.0_281/bin/java -jar /root/demo.jar
func DeployAppInContainer(name, languagePath, commandArgs string) error{
	err := Run0("nohup","./pouch", "deploy", "--name", name, languagePath, commandArgs, "> /dev/null 2>&1 &")
	return err
}

// deploy
// ./pouch deploy -kill --name spring-prod10000 /jdk1.8.0_281/bin/java -jar /root/demo.jar
func ReDeployAppInContainer(name, languagePath, commandArgs string) error{
	err := Run0("nohup","./pouch", "deploy", "-kill", "--name", name, languagePath, commandArgs, "> /dev/null 2>&1 &")
	return err
}

// execute
// ./pouch execute --name spring-prod10000 /bin/chmod 666 /dev/null
func ExecuteCommandInContainer(name, command string, args ...string) error{
	var all []string
	all = append(all, []string{"execute", "--name", name, command}...)
	all = append(all, args...)
	_,bufErr,err := Run("./pouch", all...)
	if err!= nil {
		return err
	}
	if bufErr.Len() > 0 {
		return fmt.Errorf("error while execute pouch execute, error: %s", bufErr.String())
	}
	return err
}

//executeonce
// ./pouch executeonce --name spring-prod10000 --commands /bin/chmod|666|/dev/null,/sbin/usr/service|mysql|start
func ExecuteBatchCommandOnceInContainer(name string, commandsArgs string) error{
	var all []string
	all = append(all, []string{"executeonce", "--name", name, commandsArgs}...)
	buffOut,buffErr,err := Run("./pouch", all...)
	if buffOut!=nil {
		fmt.Print(buffOut.String())
	}
	if buffErr!=nil {
		fmt.Print(buffErr.String())
	}

	return err
}

//start
// ./pouch start -d --name spring-prod10000 top -b
func StartAppInContainer(name, initCommand string) error{
	_, _, err := Run("./pouch", "start", "-d", "--name", name, initCommand)
	return err
}

// enhance
// ./pouch enhance --name spring-prod10000 --storepath /root --executable /root/demo.jar
// ./pouch enhance --name spring-prod10000 --storepath /root --executable https://nexus3/interpidtjuniversity/init/release/v20210421
func EnhanceContainer(name, storePath string, executable string) error{
	_, buffErr, err := Run("./pouch", "enhance", "--name", name, "--storepath", storePath, "--executable", executable)
	if err != nil {
		return err
	}
	if buffErr.Len() > 0 {
		return fmt.Errorf("error while execute pouch enhance, error: %s", buffErr.String())
	}
	return nil
}

// remove network
// ./pouch network remove basebridge
func RemoveNetWork(name string) error {
	_,_,err := Run("./pouch", "network", "remove", name)
	return err
}


// get current project's root path
// return path not contain the exec file
func GetProjectRoot() string {
	var (
		path string
		err  error
	)
	defer func() {
		if err != nil {
			panic(fmt.Sprintf("GetProjectRoot error :%+v", err))
		}
	}()
	path, err = filepath.Abs(filepath.Dir(os.Args[0]))
	return path
}