package invokerImpl

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PipeLIneService(t *testing.T) {
	//response, _ := InvokeRegisterMergeRequestService()
	//assert.NotEqual(t, response, "")
	//fmt.Print(response.ShowDiffUri)
	//fmt.Print("\n")
	//fmt.Print(response.MRId)
}

func TestInvokeQueryMasterLatestCommitIdService(t *testing.T) {
	id ,err := InvokeQueryMasterLatestCommitIdService("interpidtjuniversity","init")
	assert.Nil(t, err)
	fmt.Print(id)
}

func TestInvokeCreateBranchService(t *testing.T) {
	err := InvokeCreateBranchService(map[string]interface{}{
		"ownerName": "interpidtjuniversity",
		"repoName":"init",
		"iterBranch":"9c470110af_2021_5_7",
		"needMR":true,
		"protected":true,
	})
	assert.Nil(t, err)
}

func TestInvokeQueryAppAllBranchesService(t *testing.T) {
	branches, err := InvokeQueryAppAllBranchesService("interpidtjuniversity","init")
	assert.Nil(t, err)
	fmt.Print(branches)
}

func TestInvokeMerge2BranchService(t *testing.T) {
	response, err := InvokeMerge2BranchService("interpidtjuniversity","miniselfop","5defe44db1_2021_5_10_dev","5defe44db1_2021_5_10")
	assert.Nil(t, err)
	assert.True(t, response)
}