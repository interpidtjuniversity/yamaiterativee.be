package db

type Step struct {
	ID int64
	Name string
	CreatorId int64
	CreatorName string
	Requires []string
	Command string
	Args []string
	IsPublic bool
}
