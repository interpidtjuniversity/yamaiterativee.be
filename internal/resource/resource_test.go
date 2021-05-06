package resource

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateDataBaseInGlobalDataBase(t *testing.T) {
	initGlobalMysql("172.16.1.4")
	err := CreateDataBaseInGlobalDataBase("Mysql", "root_app01_prod")
	assert.Nil(t, err)
}
