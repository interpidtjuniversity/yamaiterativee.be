package form

type Iteration struct {
	Owner   string   `binding:"Required"`
	Repo    string   `binding:"Required"`
	Creator string   `binding:"Required"`
	Admin   []string `binding:"Required"`
	Type    string   `binding:"Required"`
}
