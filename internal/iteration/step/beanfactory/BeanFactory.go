package beanfactory

import "yama.io/yamaIterativeE/internal/iteration/step/beanfactory/bean"

var beanMap = map[string]bean.Bean{}

func init() {
	beanMap["compileBean"] = &bean.CompileBean{}
	beanMap["deployBean"] = &bean.DeployBean{}
}

func GetBean(beanName string) bean.Bean{
	return beanMap[beanName]
}