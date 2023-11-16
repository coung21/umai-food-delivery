package grpc

import (
	"context"
	"log"
	"umai-auth-service/transport/grpc/grpcPb"
)

func (g *GrpcServer) GetIdentity(ctx context.Context, req *grpcPb.IdentityReq) (*grpcPb.IdentityRes, error) {
	iuser, err := g.authRepo.FindUserByID(ctx, int(req.GetUserID()))
	if err != nil {
		log.Fatal(err)
	}
	resp := &grpcPb.IdentityRes{
		UserID: int32(iuser.ID),
		Role:   iuser.Role,
	}
	return resp, nil
}
