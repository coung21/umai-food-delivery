package rest

import (
	jwt "order-service/component"
	order "order-service/interfaces"
	"order-service/transport/grpc"
)

type orderHandler struct {
	orderUc  order.Usecase
	tokenPro jwt.TokenProvider
	grpcC    *grpc.GrpcClient
}

func NewOrderHandler(orderUc order.Usecase, tokenPro jwt.TokenProvider, grpcC *grpc.GrpcClient) *orderHandler {
	return &orderHandler{orderUc: orderUc, tokenPro: tokenPro, grpcC: grpcC}
}
