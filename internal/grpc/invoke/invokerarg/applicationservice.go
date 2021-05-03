package invokerarg

type CreateApplicationOptions struct {
	UserId      int64
	UserName    string
	RepoName    string
	Description string
	IsPrivate   bool
	AutoInit    bool
}
