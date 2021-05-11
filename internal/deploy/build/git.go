package build

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

var (
	BUILD_PATH = "/var/run/yama/yamaIterativeE/repo/%s"
	BUILD_LOG_PATH = "/var/run/yama/yamaIterativeE/repo/%s/build.log"
	BUILD_RESOURCE_PATH = "/var/run/yama/yamaIterativeE/repo/%s/%s/target/"
	APPLICATION_PROPERTIES_PATH = "/var/run/yama/yamaIterativeE/repo/%s/%s/src/main/resources/application.properties"
)

func Clone(repoPath, branchName, tmpDir, appName string) ([]byte, string, error) {
	bufOut := new(bytes.Buffer)
	bufErr := new(bytes.Buffer)
	workDir := fmt.Sprintf(BUILD_PATH, tmpDir)
	if err := os.MkdirAll(workDir, os.ModePerm); err != nil {
		return nil, "", err
	}
	cmd := exec.Command("git", "clone", "-b", branchName, repoPath)
	cmd.Dir = workDir
	cmd.Stdout = bufOut
	cmd.Stderr = bufErr
	err := cmd.Run()

	if err != nil {
		return nil, "", err
	}
	bufOut.Write(bufErr.Bytes())
	return bufOut.Bytes(), fmt.Sprintf("%s/%s",workDir, appName),nil
}
