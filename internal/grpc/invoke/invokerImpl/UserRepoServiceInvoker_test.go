package invokerImpl

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"yama.io/yamaIterativeE/internal/grpc/invoke/invokerarg"
)

func Test_CreateApplication(t *testing.T) {
	res, err := InvokeCreateApplicationService(invokerarg.CreateApplicationOptions{
		UserId: 2,
		UserName: "root",
		RepoName: "miniselfop",
		AutoInit: false,
		IsPrivate: false,
		Description: "use miniselfop to test grpc create application",
	})
	assert.Nil(t, err)
	assert.True(t, res.Success)
	fmt.Print(res)
}
