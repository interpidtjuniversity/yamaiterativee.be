package db

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func Test_ImagePut(t *testing.T) {
	data, err := ioutil.ReadFile("pandavpn.png")
	assert.Nil(t, err)
	success, err := PutImage("testagain",data)
	assert.Nil(t, err)
	assert.True(t, success)
}
