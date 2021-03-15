package serviceImpl

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"yama.io/yamaIterativeE/internal/grpc/service"
)

type PileLineService struct{}

func (p *PileLineService) StartYaMaPipeLine(ctx context.Context, request *service.StartYaMaPipeLineRequest) (*service.StartYaMaPipeLineResponse, error) {
	return &service.StartYaMaPipeLineResponse{
		Success: true,
	}, nil
}

func StartPipeLineService() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//建立 gPRC 服务器，并注册服务
	s := grpc.NewServer()
	service.RegisterYaMaPipeLineServiceServer(s, &PileLineService{})

	log.Println("Server run ...")
	//启动服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("fail to serve: %v", err)
	}
}


