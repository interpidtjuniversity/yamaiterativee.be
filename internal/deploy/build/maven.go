package build

import (
	"bytes"
	"fmt"
	"os/exec"
)

func MavenCompile(path string) ([]byte, error) {
	bufOut := new(bytes.Buffer)
	bufErr := new(bytes.Buffer)
	cmd := exec.Command("mvn", "compile")
	cmd.Dir = path
	cmd.Stdout = bufOut
	cmd.Stderr = bufErr
	err := cmd.Run()

	if err != nil {
		return nil, err
	}
	bufOut.Write(bufErr.Bytes())
	return bufOut.Bytes(), nil
}

func MavenInstall(path string) ([]byte, error) {
	bufOut := new(bytes.Buffer)
	bufErr := new(bytes.Buffer)
	cmd := exec.Command("mvn", "install")
	cmd.Dir = path
	cmd.Stdout = bufOut
	cmd.Stderr = bufErr
	err := cmd.Run()

	if err != nil {
		return nil, err
	}
	if bufErr.Len() > 0{
		return nil, fmt.Errorf("error while execute git clone, err: %s", bufErr.String())
	}
	return bufOut.Bytes(), nil
}
