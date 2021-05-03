package invokerImpl

import (
	"context"
	"time"
	"yama.io/yamaIterativeE/internal/grpc/invoke"
	"yama.io/yamaIterativeE/internal/grpc/invoke/invokerarg"
	"yama.io/yamaIterativeE/internal/grpc/invoke/invokerresult"
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

func InvokeCreateApplicationService(options invokerarg.CreateApplicationOptions) (*invokerresult.CreateApplicationResult, error){
	conn := invoke.GetConnection()
	defer invoke.Return(conn)
	client := invoke.NewYaMaHubApplicationServiceClient(conn)
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.CreateApplication(context.Background(), &invoke.CreateApplicationRequest{
		UserId: options.UserId,
		RepoName: options.RepoName,
		IsPrivate: options.IsPrivate,
		AutoInit: options.AutoInit,
		Description: options.Description,
		UserName: options.UserName,
	})
	baseResult := &invokerresult.CreateApplicationResult{Success: false}
	if err!=nil || !response.Success {
		return baseResult, err
	}
	baseResult.Description = response.Description
	baseResult.Success = response.Success
	baseResult.RepoName = response.RepoName
	baseResult.Private = response.Private
	baseResult.WebSite = response.WebSite
	baseResult.FullRepoName = response.FullRepoName
	baseResult.HtmlUrl = response.HtmlUrl
	baseResult.SshUrl = response.SshUrl
	baseResult.CloneUrl = response.CloneUrl
	baseResult.RepoId = response.RepoId
	baseResult.Owner = response.Owner
	baseResult.DefaultBranch = response.DefaultBranch
	return baseResult, nil

}
