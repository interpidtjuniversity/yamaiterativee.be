package invokerImpl

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	"yama.io/yamaIterativeE/internal/grpc/invoke"
)

func InvokeRegisterMergeRequestService() *invoke.RegisterMRResponse {
	//连接到gRPC服务端
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//建立客户端
	c := invoke.NewYaMaHubBranchServiceClient(conn)

	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用方法
	r, err := c.RegisterMergeRequest(context.Background(), &invoke.RegisterMRRequest{OwnerName: "interpidtjuniversity", RepoName: "init", SourceBranch: "dev_0311", TargetBranch: "master"})
	if err != nil {
		log.Fatalf("couldn not greet: %v", err)
		return nil
	}
	return r
}

func InvokeQueryRepoBranchCommitService(ownerName, repoName, branchName string) (string, string){

	conn := invoke.GetConnection()
	defer invoke.Return(conn)
	client := invoke.NewYaMaHubBranchServiceClient(conn)
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, _ := client.QueryRepoBranchCommit(context.Background(), &invoke.CommitQueryRequest{OwnerName: ownerName, RepoName: repoName, BranchName: branchName})
	if response != nil {
		return response.CommitId[:8], response.Url
	}

	return "",""
}

func InvokeCreateBranchService(data map[string]interface{}) error {
	conn := invoke.GetConnection()
	defer invoke.Return(conn)
	client := invoke.NewYaMaHubBranchServiceClient(conn)
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := client.CreateBranch(context.Background(), &invoke.CreateBranchRequest{
		UserName: data["ownerName"].(string),
		Repository: data["repoName"].(string),
		IsLock: data["protected"].(bool),
		Branch: data["iterBranch"].(string),
		NeedMr: data["needMR"].(bool),
	})

	return err
}

func InvokeQueryMasterLatestCommitIdService(ownerName, repoName string) (string, error){
	conn := invoke.GetConnection()
	defer invoke.Return(conn)
	client := invoke.NewYaMaHubBranchServiceClient(conn)
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.QueryMasterLatestCommit(context.Background(), &invoke.MasterLatestCommitQueryRequest{
		OwnerName: ownerName,
		RepoName: repoName,
	})

	return response.CommitId, err
}