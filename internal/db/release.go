package db

import "xorm.io/builder"

type Release struct {
	ID          int64  `xorm:"id autoincr pk"`
	AppOwner    string `xorm:"app_owner"`
	AppName     string `xorm:"app_name"`
	CommitId    string `xorm:"commit_id"`
	CommitLink  string `xorm:"commit_link"`
	Time        string `xorm:"time"`
	IterationId int64  `xorm:"iter_id"`
	URL         string `xorm:"url"`
}

func InsertRelease(release *Release) error {
	_, err := x.Table("release").Insert(release)
	return err
}

func GetReleaseByAppOwnerAndAppName(appOwner, appName string) ([]*Release, error){
	var releases []*Release
	err := x.Table("release").Where(builder.Eq{"app_owner": appOwner, "app_name": appName}).Find(&releases)
	return releases, err
}

func GetReleaseById(id int64) *Release {
	release := new(Release)
	exist, err := x.Table("release").Where(builder.Eq{"id": id}).Get(release)
	if err!=nil || !exist{
		return nil
	}
	return release
}