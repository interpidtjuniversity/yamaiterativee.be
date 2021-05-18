package invokerImpl

import (
	"context"
	"time"
	"yama.io/yamaIterativeE/internal/grpc/invoke"
)

func InvokeRegisterMergeRequestService(ownerName, repoName, sourceBranch, targetBranch string,
	actionId, stageId, stepId int64, reviewers []string) (int64, string) {
	conn := invoke.GetConnection()
	defer invoke.Return(conn)
	client := invoke.NewYaMaHubBranchServiceClient(conn)
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, _ := client.RegisterMergeRequest(context.Background(), &invoke.RegisterMRRequest{
		OwnerName: ownerName,
		RepoName: repoName,
		SourceBranch: sourceBranch,
		TargetBranch: targetBranch,
		ActionId: actionId,
		StagedId: stageId,
		StepId: stepId,
		Reviewers: reviewers,
	})

	return response.MRId, response.ShowDiffUri
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

	if response!=nil {
		return response.CommitId, err
	}
	return "", err
}

func InvokeQueryAppAllBranchesService(ownerName, repoName string) ([]string, error) {
	conn := invoke.GetConnection()
	defer invoke.Return(conn)
	client := invoke.NewYaMaHubBranchServiceClient(conn)
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.QueryAppAllBranch(context.Background(), &invoke.QueryAppAllBranchRequest{
		AppName: repoName,
		AppOwner: ownerName,
	})

	if response != nil {
		return response.AppBranches, err
	}
	return nil, err
}

func InvokeMerge2BranchService(userName, repoName, source, target, mergeInfo string) (bool, error){
	conn := invoke.GetConnection()
	defer invoke.Return(conn)
	client := invoke.NewYaMaHubBranchServiceClient(conn)
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.Merge2Branch(context.Background(), &invoke.Merge2BranchRequest{
		UserName: userName,
		Repository: repoName,
		SourceBranch: source,
		TargetBranch: target,
		MergeInfo: mergeInfo,
	})
	if response != nil {
		return response.Success, err
	}
	return false, err
}