package grpc

import (
	"log"
	"net"
	auth "umai-auth-service/interfaces"
	"umai-auth-service/transport/grpc/grpcPb"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	grpcPb.MenuAuthServiceServer
	authRepo auth.Repository
}

// func NewGrpcServer(authRepo auth.Repository) *GrpcServer {
// 	return &GrpcServer{authRepo: authRepo}
// }

func RunGrpcServer(authRepo auth.Repository) {
	lis, err := net.Listen("tcp", "localhost:50868")

	if err != nil {
		log.Fatalf("error while create listen %v", err)
	}

	s := grpc.NewServer()

	grpcPb.RegisterMenuAuthServiceServer(s, &GrpcServer{authRepo: authRepo})
	log.Println("MenuAuth Grpc Server is running on", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Printf("Error while serve grpc server %v", err)
	}
}
