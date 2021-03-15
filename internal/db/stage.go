package db

type Stage struct {
	ID int64
	Name string
	CreatorId int64
	CreatorName string
	IsPublic bool
	Steps []int64
}
