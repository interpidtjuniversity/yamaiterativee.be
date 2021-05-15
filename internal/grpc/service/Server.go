package service

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"yama.io/yamaIterativeE/internal/grpc/service/serviceImpl"
)

func Start() {
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//建立 gPRC 服务器，并注册服务
	s := grpc.NewServer()
	// start all grpc service
	serviceImpl.RegisterYaMaPipeLineServiceServer(s, &serviceImpl.PileLineService{})

	log.Println("Server run ...")
	//启动服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("fail to serve: %v", err)
	}
}
