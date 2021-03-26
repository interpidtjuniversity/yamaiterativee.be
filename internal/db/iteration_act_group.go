package db

type IterationActGroup struct {
	ID          int64  `xorm:"id autoincr pk"`
	GroupType   string `xorm:"group_type"`
	IterationId int64  `xorm:"iter_id"`
}
