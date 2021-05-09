package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InitJavaSpringConfig(t *testing.T) {
	NewEngine()
	err := InsertJavaSpringConfig()
	assert.Nil(t, err)
}
func TestInsertSpringConsulConfig(t *testing.T) {
	NewEngine()
	err := InsertSpringConsulConfig()
	assert.Nil(t, err)
}
func TestInsertSpringGRPCConfig(t *testing.T) {
	NewEngine()
	err := InsertSpringGRPCConfig()
	assert.Nil(t, err)
}
func TestInsertSpringMysqlConfig(t *testing.T) {
	NewEngine()
	err := InsertSpringMysqlConfig()
	assert.Nil(t, err)
}
func TestInsertSpringActuatorConfig(t *testing.T) {
	NewEngine()
	err := InsertSpringActuatorConfig()
	assert.Nil(t, err)
}
func TestInsertSpringZipkinConfig(t *testing.T) {
	NewEngine()
	err := InsertSpringZipkinConfig()
	assert.Nil(t, err)
}


func Test_GetJavaSpringConfig(t *testing.T) {
	NewEngine()
	config := GetJavaSpringConfig()
	assert.Equal(t, int64(1),config.ID)
}