package grpc

import (
	"log"
	"menu-service/transport/grpc/grpcPb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	AuthC grpcPb.IdentityServiceClient
}

func RunGrpcClient() *GrpcClient {
	conn, err := grpc.Dial("localhost:50868", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("error while dial %v", err)
	}

	// defer conn.Close()

	authc := grpcPb.NewIdentityServiceClient(conn)

	return &GrpcClient{AuthC: authc}
}
