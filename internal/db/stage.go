package db

type Stage struct {
	ID          int64   `xorm:"id autoincr pk"`
	Name        string  `xorm:"name"`
	CreatorId   int64   `xorm:"creator_id"`
	CreatorName string  `xorm:"creator_name"`
	IsPublic    bool    `xorm:"is_public"`
	Steps       []int64 `xorm:"steps" json:"-"`
}
