package db

type Release struct {
	ID       int64  `xorm:"id autoincr pk"`
	Message  string `xorm:"message"`
	AppOwner string `xorm:"app_owner"`
	AppName  string `xorm:"app_name"`
	CommitID string `xorm:"commit_id"`
}
