package db

import "xorm.io/builder"

type Deploy struct {
	ID            int64  `xorm:"id autoincr pk"`
	ServerName    string `xorm:"server_name"`
	DeployId      string `xorm:"deploy_id"`
	DeployLogPath string `xorm:"deploy_log_path"`
}

func InsertDeploy(deploy *Deploy) error {
	_,err := x.Table("deploy").Insert(deploy)
	return err
}

func GetDeployLogPathByDeployId(deployId string) (string, error) {
	deploy := new(Deploy)
	exist, err := x.Table("deploy").Cols("deploy_log_path").Where(builder.Eq{"deploy_id": deployId}).Get(deploy)
	if err!=nil || !exist {
		return "", err
	}
	return deploy.DeployLogPath, nil
}