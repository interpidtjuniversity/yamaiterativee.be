package route

import "yama.io/yamaIterativeE/internal/db"

// customConf yamaIterative.ini
func GlobalInit(customConf string) error {
	return db.NewEngine()
}
