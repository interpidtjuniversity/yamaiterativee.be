package db

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func Test_GetNetWorkByType(t *testing.T) {
	NewEngine()
	networks, err := GetNetWorkIPRangeByType("TYPE_C")
	assert.Nil(t, err)
	for _, v := range networks {
		assert.NotEqual(t, "", v.IPRange)
	}
}


func Test_AllocateNetWork(t *testing.T) {
	NewEngine()
	network := AllocateNetWork("TYPE_A")
	assert.NotEqual(t, network,"")
	fmt.Print(network)
}

func Test_Random(t *testing.T) {
	rand.Seed(time.Now().Unix())
	fmt.Print(rand.Intn(3))
}

func Test_GetNetWorkByName(t *testing.T) {
	NewEngine()
	network, err := GetApplicationNetWorkByName("testC")
	assert.Nil(t, err)
	assert.NotEqual(t, network.Id, 0)
}