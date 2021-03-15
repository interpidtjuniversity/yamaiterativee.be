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
