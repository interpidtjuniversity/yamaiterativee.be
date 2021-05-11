package resource

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateDataBaseInGlobalDataBase(t *testing.T) {
	initGlobalMysql("172.16.1.2")
	err := CreateDataBaseInGlobalDataBase("SPRING_MYSQL", "interpidtjuniversity_miniselfop_dev")
	assert.Nil(t, err)
}
