package bean

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strconv"
)

type NexusUploadBean struct {
	CompileBean
}

//   0        1     2      3      4      5      6         7         8            9
//appOwner appName repo branch execPath env serverName serverIP iterId(string) logPath
func (nub NexusUploadBean) Execute(stringArgs []string, env *map[string]interface{}) error {
	// 1.parse arg
	if len(stringArgs) != 10{
		return fmt.Errorf("arguement error")
	}
	iterId,err := strconv.Atoi(stringArgs[8])
	if err!=nil {
		return err
	}
	err = nub.clone(stringArgs[2], stringArgs[3], stringArgs[4], stringArgs[9])
	if err!=nil {
		return err
	}
	err = nub.flushConfig(stringArgs[1], stringArgs[4], stringArgs[5], stringArgs[6], stringArgs[7], int64(iterId))
	if err!=nil {
		return err
	}
	// through pom.xml, a jar will be upload to nexus
	err = nub.mavenDeploy(stringArgs[1], stringArgs[4], env)

	return err
}

func (nub NexusUploadBean) mavenDeploy(appName, execPath string, env *map[string]interface{}) error{
	bufOut := new(bytes.Buffer)
	bufErr := new(bytes.Buffer)

	cmd := exec.Command("mvn", "deploy")
	cmd.Dir = fmt.Sprintf("%s/%s", execPath, appName)
	cmd.Stdout = bufOut
	cmd.Stderr = bufErr
	err := cmd.Run()

	if err != nil {
		return err
	}
	if bufErr.Len() > 0 {
		return fmt.Errorf("error while execute maven deploy, err: %s", bufErr.String())
	}
	var targetLine []byte
	reader := bufio.NewReader(bufOut)
	for{
		line, _, e := reader.ReadLine()
		if e == io.EOF{
			break
		}
		if bytes.HasPrefix(line, []byte("Uploading to releases: ")) && bytes.HasSuffix(line, []byte(".jar")){
			targetLine = line
			break
		}
	}
	if targetLine != nil {
		(*env)["jarNexusURL"] = string(targetLine[23:])
	}

 	return nil
}