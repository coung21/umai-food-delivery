package grpc

import (
	"common"
	"context"
	"errors"
	"order-service/transport/grpc/grpcPb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetMenuItemHdl(c grpcPb.MenuItemServiceClient, id string) (*string, error) {
	resp, err := c.GetMenuItem(context.Background(), &grpcPb.GetMenuItemReq{
		Id: id,
	})
	if err != nil {
		if errStatus, ok := status.FromError(err); ok {
			if errStatus.Code() == codes.NotFound {
				return nil, common.NotFound
			} else if errStatus.Code() == codes.InvalidArgument {
				return nil, common.BadQueryParams
			}
			return nil, errors.New(errStatus.Message())
		}
	}
	data := resp.GetData()
	return &data, nil
}

func GetUserIdentityHdl(c grpcPb.IdentityServiceClient, id int) (*int, error) {
	resp, err := c.GetUserIdentity(context.Background(), &grpcPb.IdentityReq{
		UserID: int32(id),
	})
	if err != nil {
		if errStatus, ok := status.FromError(err); ok {
			if errStatus.Code() == codes.Unauthenticated {
				return nil, common.Unauthorized
			}
			return nil, errors.New(errStatus.Message())
		}
	}

	uid := int(resp.GetUserID())

	return &uid, nil
}
