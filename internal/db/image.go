package db

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const (
	OSS_ENDPOINT = "https://oss-cn-hangzhou.aliyuncs.com"
	ACCESS_KEY = "LTAI5tJ8otAjZGng83ksbpXX"
	ACCESS_KEY_SECRET = "OfcgDQjx8RyNQSKZ9vD3jt8HzTkc9X"
	BUCKET = "3levelimage"
)

func PutImage(name string, img []byte) (bool, error){
	client, err := oss.New(OSS_ENDPOINT, ACCESS_KEY,ACCESS_KEY_SECRET)
	if err != nil {
		return false, err
	}
	bucket, err := client.Bucket(BUCKET)
	if err != nil {
		return false, err
	}
	err = bucket.PutObject(fmt.Sprintf("%s.png", name), bytes.NewReader(img))
	if err != nil {
		return false, err
	}
	return true, nil

}