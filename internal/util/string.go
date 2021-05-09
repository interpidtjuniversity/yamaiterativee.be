package util

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func GenerateRandomStringWithSuffix(len int, suffix string) string{
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

func GenerateRandomStringWithPrefix(len int, prefix string) string{
	id := strings.ReplaceAll(generateUUId(),"-","")
	if len<=32 {
		if prefix == "" {
			return id[:len]
		}else {
			return fmt.Sprintf("%s_%s", prefix, id[:len])
		}
	} else {
		if prefix == "" {
			return id
		}else {
			return fmt.Sprintf("%s_%s", prefix, id)
		}
	}
}

func generateUUId() string {
	uid, _ := uuid.NewUUID()
	return uid.String()
}