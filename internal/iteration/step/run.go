package step

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"yama.io/yamaIterativeE/internal/iteration/step/beanfactory"
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

func RunCodeStep(name string, args ...string) error {
	bean := beanfactory.GetBean(name)
	if bean == nil {
		return fmt.Errorf("no such bean: %s", name)
	}
	return bean.Execute(args, nil)
}

func RunCodeStepWithResult(name string, env *map[string]interface{}, args ...string) error {
	bean := beanfactory.GetBean(name)
	if bean == nil {
		return fmt.Errorf("no such bean: %s", name)
	}
	return bean.Execute(args, env)
}