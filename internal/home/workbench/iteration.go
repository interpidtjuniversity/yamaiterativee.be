package workbench

import (
	"encoding/json"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/grpc/invoke/invokerImpl"
)

func GetAllOwners(c *context.Context) []byte {
	names, err := invokerImpl.InvokeQueryApplicationOwners()
	if err != nil {
		return nil
	}
	data, _ := json.Marshal(names)
	return data
}

func GetOwnerApplications(c *context.Context) []byte {
	ownerName := c.Params(":ownerName")
	names, err := invokerImpl.InvokeQueryApplications(ownerName)
	if err != nil {
		return nil
	}

	data, _ := json.Marshal(names)
	return data
}
