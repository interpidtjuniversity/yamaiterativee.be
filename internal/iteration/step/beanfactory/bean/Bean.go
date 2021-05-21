package bean


type Bean interface {
	Execute([]string, *map[string]interface{}) error
}
