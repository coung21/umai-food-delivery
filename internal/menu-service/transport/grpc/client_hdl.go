package grpc

import (
	"common"
	"context"
	"errors"
	"menu-service/transport/grpc/grpcPb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetResIdentityHdl(c grpcPb.IdentityServiceClient, uid int) (map[string]interface{}, error) {
	resp, err := c.GetResIdentity(context.Background(), &grpcPb.IdentityResReq{
		UserID: int32(uid),
	})
	if err != nil {
		if statusErr, ok := status.FromError(err); ok {
			if statusErr.Code() == codes.Unauthenticated {
				return nil, common.Unauthorized
			}
			return nil, errors.New(statusErr.Message())
		}
	}

	indentity := map[string]interface{}{
		"user_id":       int(resp.GetUserID()),
		"role":          resp.GetRole(),
		"restaurant_id": int(resp.GetRestaurantID()),
	}

	return indentity, nil
}
