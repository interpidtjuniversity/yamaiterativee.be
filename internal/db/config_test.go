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


func Test_GetJavaSpringConfig(t *testing.T) {
	NewEngine()
	config := GetJavaSpringConfig()
	assert.Equal(t, int64(1),config.ID)
}