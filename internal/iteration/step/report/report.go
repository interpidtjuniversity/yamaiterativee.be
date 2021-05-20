package report

import "fmt"

func GetFileCovered(execPath, appName, packageName, fileName string) []byte {
	filePath := fmt.Sprintf("%s/%s/target/site/jacoco/%s/%s.java.html", execPath, appName, packageName, fileName)

}
