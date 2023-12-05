package grpc

import (
	"context"
	"order-service/transport/grpc/grpcPb"
)

func GetMenuItemHdl(c grpcPb.MenuItemServiceClient, id string) (string, error) {
	resp, err := c.GetMenuItem(context.Background(), &grpcPb.GetMenuItemReq{
		Id: id,
	})
	if err != nil {
		return "", err
	}

	return resp.GetData(), nil
}
