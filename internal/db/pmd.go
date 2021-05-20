package db

import "xorm.io/builder"

type PMD struct {
	ID       int64  `xorm:"id autoincr pk"`
	RuleName string `xorm:"rule_name"`
	Type     string `xorm:"type"`
}

func GetPMDByRuleName(names []string) map[string]string {
	var pmds []*PMD
	m := make(map[string]string)
	err := x.Table("pmd").Cols("rule_name","type").Where(builder.In("rule_name",names)).Find(&pmds)
	if err!=nil {
		return m
	}
	for i:=0; i<len(pmds); i++ {
		m[pmds[i].RuleName] = pmds[i].Type
	}
	return m
}