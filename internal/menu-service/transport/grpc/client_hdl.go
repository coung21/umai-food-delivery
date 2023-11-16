package grpc

import (
	"context"
	"log"
	"menu-service/transport/grpc/grpcPb"
)

func GetIdentityHdl(c grpcPb.MenuAuthServiceClient, uid int) (int, string) {
	resp, err := c.GetIdentity(context.Background(), &grpcPb.IdentityReq{
		UserID: int32(uid),
	})
	if err != nil {
		log.Fatalf("error while call %v", err)
	}

	return int(resp.GetUserID()), resp.GetRole()
}
