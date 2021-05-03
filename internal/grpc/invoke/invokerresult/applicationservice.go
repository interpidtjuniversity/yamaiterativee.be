package invokerresult


type CreateApplicationResult struct {
	RepoId        int64
	Owner         string
	RepoName      string
	FullRepoName  string
	Description   string
	Private       bool
	HtmlUrl       string
	SshUrl        string
	CloneUrl      string
	WebSite       string
	DefaultBranch string
	Success       bool
}
