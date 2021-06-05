package util

import "time"

func GetDefaultCurrentTime() string {
	return time.Now().Format("2006/02/01-15:04:05")
}