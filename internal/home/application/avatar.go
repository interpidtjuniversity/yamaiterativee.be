package application

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"yama.io/yamaIterativeE/internal/resource"
)

var (
	AliYunOssImageClient *oss.Client = nil
)

const (
	OSS_ENDPOINT = "https://oss-cn-hangzhou.aliyuncs.com"
	BUCKET = "3levelimage"
)

func InitAliYunOSS() {
	client, err := oss.New(OSS_ENDPOINT, resource.GLOBAL_ALIYUN_ACCESSKEY,resource.GLOBAL_ALIYUN_ACCESSKEY_SECRET)
	if err != nil {
		panic("error while init oss resource")
	}
	AliYunOssImageClient = client
}

func PutImage(name string, img []byte) (bool, error){

	bucket, err := AliYunOssImageClient.Bucket(BUCKET)
	if err != nil {
		return false, err
	}
	err = bucket.PutObject(fmt.Sprintf("%s.png", name), bytes.NewReader(img))
	if err != nil {
		return false, err
	}
	return true, nil

}
