package network

import (
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/deploy/container/yamapouch/command"
)

const (
	OWNER = "SYSTEM"
	NAME = "GLOBAL_RESOURCE"
)

func InitNetwork() {
	network,_ := db.GetApplicationNetWorkByName(NAME)
	if network==nil {
		network, err := db.CreateApplicationNetWork(NAME, OWNER, NAME)
		if err != nil {
			panic("allocate global resource network IPRange error")
		}
		err = command.CreateNetWork(NAME, network.IPRange, "bridge")
		if err != nil {
			panic("init global resource network error")
		}
	}
}
