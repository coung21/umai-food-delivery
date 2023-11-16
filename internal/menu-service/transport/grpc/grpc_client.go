package grpc

import (
	"log"
	"menu-service/transport/grpc/grpcPb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	Client grpcPb.MenuAuthServiceClient
}

func RunGrpcClient() *GrpcClient {
	conn, err := grpc.Dial("localhost:50868", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("error while dial %v", err)
	}

	// defer conn.Close()

	client := grpcPb.NewMenuAuthServiceClient(conn)

	return &GrpcClient{Client: client}
}
