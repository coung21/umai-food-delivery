package rest

import (
	jwt "menu-service/component"
	menu "menu-service/interfaces"
	"menu-service/transport/grpc"
)

type menuHandler struct {
	menuUC        menu.Usecase
	grpcC         *grpc.GrpcClient
	tokenProvider jwt.TokenProvider
}

func NewMenuHandler(menuUC menu.Usecase, grpcC *grpc.GrpcClient, tokenProvider jwt.TokenProvider) *menuHandler {
	return &menuHandler{menuUC: menuUC, grpcC: grpcC, tokenProvider: tokenProvider}
}
