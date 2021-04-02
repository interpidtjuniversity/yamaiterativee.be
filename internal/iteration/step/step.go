package step

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	YAMALogPath = "/root/yamaIterativeE/yamaIterativeE-step-logs/%d%s.log"
	StrLength   = 10
	StrBase     = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

type InfoStep struct {
	StepId     int64  `json:"stepId"`
	Title      string `json:"title"`
	Image      string `json:"image"`
	Index      int    `json:"index"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func FormatLogPath() string{
	return fmt.Sprintf(YAMALogPath, time.Now().UnixNano(), GetRandomString(StrLength))
}

func GetRandomString(n int) string {
	bytes := []byte(StrBase)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
