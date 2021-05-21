package report

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"yama.io/yamaIterativeE/internal/util"
)

type JacocoFileCoverLine struct {
	Content string  `json:"content"`
	Type    string  `json:"type"`
}

func GetFileCovered(execPath, appName, packageName, fileName string) []byte {
	filePath := fmt.Sprintf("%s/%s/target/site/jacoco/%s/%s.java.html", execPath, appName, packageName, fileName)
	if !util.IsExist(filePath) {
		return nil
	}
	jacocoFile, _:= os.Open(filePath)
	defer jacocoFile.Close()
	buffer := bufio.NewReader(jacocoFile)
	index := 0
	var lines []string
	for {
		line, _, e := buffer.ReadLine()
		if e == io.EOF {
			break
		}
		lines = append(lines, string(line))
		index++
	}
	if lines!=nil {
		lines = lines[1:len(lines)-1]
	}
	var jacocoFileCoverLine []JacocoFileCoverLine
	for i:=0; i<len(lines); i++ {
		coverLine := JacocoFileCoverLine{}
		coverLine.Content = lines[i]
		jacocoFileCoverLine = append(jacocoFileCoverLine, coverLine)
	}

	data, _ := json.Marshal(jacocoFileCoverLine)

	return data
}
