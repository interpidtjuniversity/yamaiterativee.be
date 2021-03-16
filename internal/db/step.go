package db

type Step struct {
	ID          int64    `xorm:"id autoincr pk"`
	Name        string   `xorm:"name"`
	CreatorId   int64    `xorm:"creator_id"`
	CreatorName string   `xorm:"creator_name"`
	Requires    []string `xorm:"requires" json:"-"`
	Command     string   `xorm:"command"`
	Args        []string `xorm:"args" json:"-"`
	IsPublic    bool     `xorm:"is_public"`
}
