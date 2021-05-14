package db

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"yama.io/yamaIterativeE/internal/util"
)

func TestInsertServer(t *testing.T) {
	NewEngine()
	newServer := &Server{AppOwner: "interpidtjuniversity", AppName: "init", IterationId: 1, Owner: "interpidtjuniversity", AppType: JAVA_SPRING,
		Branch: "master", Type: DEV, State: APPLYING, CreatedTime: time.Now().Format("2006-01-01 15:04:05"),
		Name: util.GenerateRandomStringWithSuffix(20, fmt.Sprintf("%s.%s.%s", "interpidtjuniversity", "init", DEV.ToString())),
	}
	newServer.NetWork = "test_network"

	_, err := InsertServer(newServer)
	assert.Nil(t, err)
}

func TestReleaseServer(t *testing.T) {
	NewEngine()
	_,err := ReleaseServer("d3358ad5b08311ebb22f_interpidtjuniversity.init.dev")
	assert.Nil(t, err)
}


func TestUpdateServerAfterApply(t *testing.T) {
	NewEngine()
	_, err := UpdateServerAfterApply("d3358ad5b08311ebb22f_interpidtjuniversity.init.dev", "192.168.255.6")
	assert.Nil(t, err)
}

func TestGetSameGroupServerByGroupIdAndServiceName(t *testing.T) {
	NewEngine()
	servers, err := GetSameGroupServerByGroupIdAndServiceName(0, "interpidtjuniversity-miniselfop")
	fmt.Print(len(servers))
	assert.Nil(t, err)
}

func TestGetDevServerByAppOwnerAndName(t *testing.T) {
	NewEngine()
	servers, err := GetDevServerByAppOwnerAndName("interpidtjuniversity","miniselfop")
	assert.Nil(t, err)
	fmt.Print(len(servers))
}

func TestBranchUpdateServerGroup(t *testing.T) {
	NewEngine()
	err := BranchUpdateServerGroup([]string{"0edef0abb19f11eb9274_interpidtjuniversity.miniselfop.dev","24391683b27111eba6c5_interpidtjuniversity.miniselfop.dev"},"")
	assert.Nil(t, err)
}