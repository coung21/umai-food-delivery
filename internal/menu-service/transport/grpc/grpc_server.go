package grpc

import (
	"log"
	menu "menu-service/interfaces"
	"menu-service/transport/grpc/grpcPb"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	// grpcPb.MenuAuthServiceServer
	grpcPb.MenuItemServiceServer
	menuRepo menu.Repository
}

func RunGrpcServer(menuRepo menu.Repository) {
	lis, err := net.Listen("tcp", "localhost:50869")

	if err != nil {
		log.Fatalf("error while create listen %v", err)
	}

	s := grpc.NewServer()

	// grpcPb.RegisterMenuAuthServiceServer(s, &GrpcServer{})
	// log.Println("MenuAuth Grpc Server is running!")
	grpcPb.RegisterMenuItemServiceServer(s, &GrpcServer{menuRepo: menuRepo})

	if err := s.Serve(lis); err != nil {
		log.Printf("Error while serve grpc server %v", err)
	}
}
