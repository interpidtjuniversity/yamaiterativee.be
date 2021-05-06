package db

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetApplicationByParticipant(t *testing.T) {
	NewEngine()
	applications,err:= GetApplicationByParticipant("1")
	assert.Nil(t, err)
	fmt.Print(applications)
}
