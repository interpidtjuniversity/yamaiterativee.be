package db

import "xorm.io/builder"

type BlackBranch struct {
	Branch   string `xorm:"branch"`
	AppOwner string `xorm:"app_owner"`
	AppName  string `xorm:"app_name"`
}

func GetAppAllBlackBranch(appOwner, appName string) []string {
	var bbs []*BlackBranch
	var branches []string
	err := x.Table("black_branch").Cols("branch").Where(builder.Eq{"app_owner":appOwner, "app_name":appName}).Find(&bbs)
	if err != nil {
		return nil
	}
	for _, bb := range bbs{
		branches = append(branches, bb.Branch)
	}
	branches = append(branches, "master")
	return branches
}