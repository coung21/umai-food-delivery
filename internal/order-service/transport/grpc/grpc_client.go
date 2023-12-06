package grpc

import (
	"log"
	"order-service/transport/grpc/grpcPb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	MenuC grpcPb.MenuItemServiceClient
	AuthC grpcPb.IdentityServiceClient
}

func RunGrpcClient() *GrpcClient {
	menuConn, err := grpc.Dial("localhost:50869", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("error while dial to menu-service: %v", err)
	}
	authConn, err := grpc.Dial("localhost:50868", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("error while dial to auth-service: %v", err)
	}

	// defer conn.Close()

	menuc := grpcPb.NewMenuItemServiceClient(menuConn)
	authc := grpcPb.NewIdentityServiceClient(authConn)

	return &GrpcClient{MenuC: menuc, AuthC: authc}
}
