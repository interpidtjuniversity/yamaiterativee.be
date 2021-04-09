package invokerImpl

import (
	"context"
	"time"
	"yama.io/yamaIterativeE/internal/grpc/invoke"
)

func InvokeQueryApplicationOwners() ([]string, error){
	conn := invoke.GetConnection()
	defer invoke.Return(conn)

	client := invoke.NewYaMaHubApplicationServiceClient(conn)
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response,err := client.QueryApplicationOwners(context.Background(), &invoke.QueryOwnerRequest{})
	if err != nil {
		return nil, err
	}

	return response.OwnerNames, nil
}

func InvokeQueryApplications(ownerName string) ([]string, error){
	conn := invoke.GetConnection()
	defer invoke.Return(conn)

	client := invoke.NewYaMaHubApplicationServiceClient(conn)
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.QueryApplications(context.Background(), &invoke.QueryApplicationRequest{OwnerName: ownerName})
	if err != nil {
		return nil, err
	}

	return response.AppNames, nil
}
