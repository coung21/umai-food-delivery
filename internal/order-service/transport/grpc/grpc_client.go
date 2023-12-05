package grpc

import (
	"log"
	"order-service/transport/grpc/grpcPb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	MenuC grpcPb.MenuItemServiceClient
}

func RunGrpcClient() *GrpcClient {
	conn, err := grpc.Dial("localhost:50869", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("error while dial %v", err)
	}

	// defer conn.Close()

	menuc := grpcPb.NewMenuItemServiceClient(conn)

	return &GrpcClient{MenuC: menuc}
}
