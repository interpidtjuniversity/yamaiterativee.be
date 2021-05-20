package step

import (
	"fmt"
	"math/rand"
	"time"
	"yama.io/yamaIterativeE/internal/iteration/step/log"
	"yama.io/yamaIterativeE/internal/iteration/step/report"
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
	Link       string `json:"link"`
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

func GetLog(path, execPath string, logTypeNum int, appName string) []byte {
	logType := log.Type(logTypeNum)
	switch logType {
	case log.PMDCollapse:
		return log.ConstructCollapseLog(path, execPath)
	case log.Text:
		return log.ConstructTextLog(path)
	case log.TestReport:
		return log.ConstructTestReport(execPath, appName)
	default:
		return nil
	}
}

func GetFileCovered(execPath, appName, packageName, fileName string) []byte {
	return report.GetFileCovered(execPath, appName, packageName, fileName)
}