package application

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPutImage(t *testing.T) {
	client, err := oss.New("https://oss-cn-hangzhou.aliyuncs.com", "xxx","xxx")
	assert.Nil(t, err)
	bucket, err := client.Bucket("3levelimage")
	assert.Nil(t, err)
	err = bucket.PutObject(fmt.Sprintf("%s.txt", "tesddt"), bytes.NewReader([]byte("tessefs")))
	assert.Nil(t, err)
}
