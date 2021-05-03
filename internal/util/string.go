package util

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func GenerateRandomString(len int, suffix string) string{
	id := strings.ReplaceAll(generateUUId(),"-","")
	if len<=32 {
		if suffix == "" {
			return id[:len]
		}else {
			return fmt.Sprintf("%s_%s", id[:len], suffix)
		}
 	} else {
		if suffix == "" {
			return id
		}else {
			return fmt.Sprintf("%s_%s", id, suffix)
		}
	}
}

func generateUUId() string {
	uid, _ := uuid.NewUUID()
	return uid.String()
}