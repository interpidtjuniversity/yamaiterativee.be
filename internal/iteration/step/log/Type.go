package log

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"yama.io/yamaIterativeE/internal/db"
)

type Type int

const (
	PMDCollapse Type = iota
	Text
	TestReport
)

// agg by key to generate a collapse
// the title of collapse item is file:line and the content is code and info
type CollapseBaseLogItem struct {
	File string `json:"file"`
	Line string `json:"line"`
	Key  string `json:"key"`
	Type string `json:"type"`
	Info string `json:"info"`
	Code string `json:"code"`
}

type JacocoTestReport struct {
	Group              string `json:"group"`
	Package            string `json:"package"`
	Class              string `json:"class"`
	InstructionMissed  int    `json:"instructionMissed"`
	InstructionCovered int    `json:"instructionCovered"`
	BranchMissed       int    `json:"branchMissed"`
	BranchCovered      int    `json:"branchCovered"`
	LineMissed         int    `json:"lineMissed"`
	LineCovered        int    `json:"lineCovered"`
	ComplexityMissed   int    `json:"complexityMissed"`
	ComplexityCovered  int    `json:"complexityCovered"`
	MethodMissed       int    `json:"methodMissed"`
	MethodCovered      int    `json:"methodCovered"`
}

type Log struct {
	Type Type        `json:"type"`
	Data interface{} `json:"data"`
}

func readLog(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	return data
}

func ConstructTextLog(path string) []byte {
	data := readLog(path)
	log := Log{Type: Text, Data: string(data)}
	content, _ := json.Marshal(log)
	return content
}

func ConstructTestReport(execPath, appName string) []byte{
	jacocoPath := fmt.Sprintf("%s/%s/target/site/jacoco/jacoco.csv", execPath, appName)
	jacoco, _:= os.Open(jacocoPath)
	defer jacoco.Close()
	jacocoMap := make(map[string][]JacocoTestReport)

	buffer := bufio.NewReader(jacoco)
	// skip first line
	index := 0
	for{
		line, _, e := buffer.ReadLine()
		if e == io.EOF{
			break
		}
		if index > 0{
			array := strings.Split(string(line), ",")
			im, _ := strconv.Atoi(array[3])
			ic, _ := strconv.Atoi(array[4])
			bm, _ := strconv.Atoi(array[5])
			bc, _ := strconv.Atoi(array[6])
			lm, _ := strconv.Atoi(array[7])
			lc, _ := strconv.Atoi(array[8])
			cm, _ := strconv.Atoi(array[9])
			cc, _ := strconv.Atoi(array[10])
			mm, _ := strconv.Atoi(array[11])
			mc, _ := strconv.Atoi(array[12])
			jacocoMap[array[1]] = append(jacocoMap[array[1]], JacocoTestReport{
				array[0],
				array[1],
				array[2],
				im,
				ic,
				bm,
				bc,
				lm,
				lc,
				cm,
				cc,
				mm,
				mc,
			})
		}
		index++
	}
	log := Log{Type: TestReport, Data: jacocoMap}
	jacocoData, _ := json.Marshal(log)

	return jacocoData
}

func ConstructCollapseLog(path, execPath string) []byte {
	execPath = fmt.Sprintf("%s/", execPath)
	data := readLog(path)
	var log Log
	if data == nil {
		log = Log{Type: PMDCollapse, Data: map[string][]CollapseBaseLogItem{}}
	} else {
		dataMap := make(map[string][]CollapseBaseLogItem)
		items := make([]CollapseBaseLogItem, 1)
		json.Unmarshal(data, &items)
		var names []string
		for i := 0; i < len(items)-1; i++ {
			names = append(names, items[i].Key)
		}
		typeMap := db.GetPMDByRuleName(names)

		for i := 0; i < len(items)-1; i++ {
			items[i].File = strings.ReplaceAll(items[i].File, execPath, "")
			items[i].Type = typeMap[items[i].Key]
			dataMap[items[i].Key] = append(dataMap[items[i].Key], items[i])
		}
		log = Log{Type: PMDCollapse, Data: dataMap}
	}
	content, _ := json.Marshal(log)
	return content
}