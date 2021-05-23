package step

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
)

func RunShellStep(name, execPath string, args ...string) (string, string, error) {
	basePath, _ := os.Getwd()
	// exec
	bufOut := new(bytes.Buffer)
	bufErr := new(bytes.Buffer)
	ctx := context.Background()
	commandStr := fmt.Sprintf("%s%s",basePath, name)
	commend := exec.CommandContext(ctx, commandStr, args...)
	commend.Dir = execPath
	commend.Stdout = bufOut
	commend.Stderr = bufErr
	err := commend.Run()
	return bufOut.String(), bufErr.String(), err
}
