package db

type User struct{
	ID int64             `xorm:"id autoincr pk"`
	Name string          `xorm:"name"`
	LowerName string     `xorm:"lower_name"`
	Email string         `xorm:"email"`
	Passwd string        `xorm:"passwd"`
	IsActive int         `xorm:"is_active"`
	Avatar string        `xorm:"avatar"`
	ExtInfo string       `xorm:"ext_info"`
}
