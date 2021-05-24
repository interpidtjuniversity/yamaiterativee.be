package db

type Release struct {
	ID          int64  `xorm:"id autoincr pk"`
	AppOwner    string `xorm:"app_owner"`
	AppName     string `xorm:"app_name"`
	CommitId    string `xorm:"commit_id"`
	Time        string `xorm:"time"`
	IterationId int64  `xorm:"iter_id"`
	URL         string `xorm:"url"`
}
