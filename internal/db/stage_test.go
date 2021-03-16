package db

import (
	"fmt"
	"testing"
)

func Test_StageInsert(t *testing.T) {
	NewEngine()

	var stages []Stage

	stages = append(stages,Stage{Name: "代码评审",CreatorId: 0, CreatorName: "interpidtjuniversity", IsPublic: true, Steps: []int64{}})
	stages = append(stages,Stage{Name: "冲突检测",CreatorId: 0, CreatorName: "interpidtjuniversity", IsPublic: true, Steps: []int64{1}})
	stages = append(stages,Stage{Name: "代码扫描",CreatorId: 0, CreatorName: "interpidtjuniversity", IsPublic: true, Steps: []int64{2,3,4}})
	stages = append(stages,Stage{Name: "预编译",CreatorId: 0, CreatorName: "interpidtjuniversity", IsPublic: true, Steps: []int64{5}})
	stages = append(stages,Stage{Name: "代码合并",CreatorId: 0, CreatorName: "interpidtjuniversity", IsPublic: true, Steps: []int64{6}})
	stages = append(stages,Stage{Name: "合并后编译",CreatorId: 0, CreatorName: "interpidtjuniversity", IsPublic: true, Steps: []int64{7}})
	stages = append(stages,Stage{Name: "质量检测",CreatorId: 0, CreatorName: "interpidtjuniversity", IsPublic: true, Steps: []int64{8,9}})

	for _, stage := range stages {
		if _, err := x.Insert(stage); err != nil {
			fmt.Print(err)
		}
	}
}
