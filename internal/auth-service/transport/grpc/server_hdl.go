package grpc

import (
	"context"
	"log"
	"umai-auth-service/transport/grpc/grpcPb"
)

func (g *GrpcServer) GetIdentity(ctx context.Context, req *grpcPb.IdentityReq) (*grpcPb.IdentityRes, error) {
	var resp *grpcPb.IdentityRes
	iuser, err := g.authRepo.FindUserByID(ctx, int(req.GetUserID()))
	if err != nil {
		log.Println(err)
		resp = &grpcPb.IdentityRes{
			UserID:       0,
			Role:         "",
			RestaurantID: 0,
		}
	}
	ires, err := g.authRepo.FindRestaurantByUserID(ctx, iuser.ID)
	if err != nil {
		log.Println(err)
		resp = &grpcPb.IdentityRes{
			UserID:       0,
			Role:         "",
			RestaurantID: 0,
		}
	}

	resp = &grpcPb.IdentityRes{
		UserID:       int32(iuser.ID),
		Role:         iuser.Role,
		RestaurantID: int32(ires.ID),
	}
	return resp, nil
}
