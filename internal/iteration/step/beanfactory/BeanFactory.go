package beanfactory

import "yama.io/yamaIterativeE/internal/iteration/step/beanfactory/bean"

var beanMap = map[string]bean.Bean{}

func init() {
	beanMap["compileBean"] = &bean.CompileBean{}
	beanMap["deployBean"] = &bean.DeployBean{}
	beanMap["historyReleaseDeployBean"] = &bean.HistoryReleaseDeployBean{}
	beanMap["serverChangeBean"] = &bean.ServerChangeBean{}
	beanMap["serverImageBuildBean"] = &bean.ServerImageBuildBean{}
	beanMap["serverReleaseBean"] = &bean.ServerReleaseBean{}
	beanMap["branchDeployBean"] = &bean.BranchDeployBean{}
	beanMap["branchCompileBean"] = &bean.BranchCompileBean{}
	beanMap["rollBackBean"] = &bean.RollBackBean{}
	beanMap["nexusUploadBean"] = &bean.NexusUploadBean{}
}

func GetBean(beanName string) bean.Bean{
	return beanMap[beanName]
}