package grpc

import (
	"log"
	"menu-service/transport/grpc/grpcPb"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	grpcPb.MenuAuthServiceServer
}

func RunGrpcServer() {
	lis, err := net.Listen("tcp", "localhost:50061")

	if err != nil {
		log.Fatalf("error while create listen %v", err)
	}

	s := grpc.NewServer()

	grpcPb.RegisterMenuAuthServiceServer(s, &GrpcServer{})
	log.Println("MenuAuth Grpc Server is running!")

	if err := s.Serve(lis); err != nil {
		log.Printf("Error while serve grpc server %v", err)
	}
}
