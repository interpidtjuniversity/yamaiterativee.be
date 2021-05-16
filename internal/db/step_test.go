package db

import (
	"fmt"
	"testing"
)

// just gen data to database
func Test_StepInsert(t *testing.T) {
	NewEngine()

	var steps []Step

	//step := Step{Name: "代码预合并",CreatorId: 0, CreatorName: "interpidtjuniversity", IsPublic: true}
	//steps = append(steps,Step{Name: "静态扫描",CreatorId: 0, CreatorName: "interpidtjuniversity", IsPublic: true})
	//steps = append(steps,Step{Name: "PMD",CreatorId: 0, CreatorName: "interpidtjuniversity", IsPublic: true})
	//steps = append(steps,Step{Name: "YamaX",CreatorId: 0, CreatorName: "interpidtjuniversity", IsPublic: true})
	//steps = append(steps,Step{Name: "预编译",CreatorId: 0, CreatorName: "interpidtjuniversity", IsPublic: true, Requires: []string{"maven"}, Command: "mvn", Env: []string{"compile"}})
	//steps = append(steps,Step{Name: "代码合并",CreatorId: 0, CreatorName: "interpidtjuniversity", IsPublic: true})
	//steps = append(steps,Step{Name: "合并后编译",CreatorId: 0, CreatorName: "interpidtjuniversity", IsPublic: true, Requires: []string{"maven"}, Command: "mvn", Env: []string{"compile"}})
	//steps = append(steps,Step{Name: "单元测试",CreatorId: 0, CreatorName: "interpidtjuniversity", IsPublic: true})
	//steps = append(steps,Step{Name: "集成测试",CreatorId: 0, CreatorName: "interpidtjuniversity", IsPublic: true})
	//steps = append(steps,Step{Name: "集成测试",CreatorId: 0, CreatorName: "interpidtjuniversity", IsPublic: true})

	for _, step := range steps {
		if _, err := x.Insert(step); err != nil {
			fmt.Print(err)
		}
	}
}
