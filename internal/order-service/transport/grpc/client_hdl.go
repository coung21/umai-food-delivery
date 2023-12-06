package grpc

import (
	"context"
	"errors"
	"log"
	"order-service/transport/grpc/grpcPb"

	"google.golang.org/grpc/status"
)

func GetMenuItemHdl(c grpcPb.MenuItemServiceClient, id string) (string, error) {
	resp, err := c.GetMenuItem(context.Background(), &grpcPb.GetMenuItemReq{
		Id: id,
	})
	if err != nil {
		if errStatus, ok := status.FromError(err); ok {
			log.Println(errStatus.Message())
			log.Println(errStatus.Code())
			return "", errors.New(errStatus.Message())
		}
	}

	return resp.GetData(), nil
}

func GetUserIdentityHdl(c grpcPb.IdentityServiceClient, id int) (*int, error) {
	resp, err := c.GetUserIdentity(context.Background(), &grpcPb.IdentityReq{
		UserID: int32(id),
	})
	if err != nil {
		if errStatus, ok := status.FromError(err); ok {
			log.Println(errStatus.Message())
			log.Println(errStatus.Code())
			return nil, errors.New(errStatus.Message())
		}
	}

	uid := int(resp.GetUserID())

	return &uid, nil
}
