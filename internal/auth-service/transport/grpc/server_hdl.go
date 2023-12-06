package grpc

import (
	"context"
	"umai-auth-service/transport/grpc/grpcPb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g *GrpcServer) GetResIdentity(ctx context.Context, req *grpcPb.IdentityResReq) (*grpcPb.IdentityResResp, error) {
	var resp *grpcPb.IdentityResResp
	iuser, err := g.authRepo.FindUserByID(ctx, int(req.GetUserID()))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	ires, err := g.authRepo.FindRestaurantByUserID(ctx, iuser.ID)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	resp = &grpcPb.IdentityResResp{
		UserID:       int32(iuser.ID),
		Role:         iuser.Role,
		RestaurantID: int32(ires.ID),
	}
	return resp, nil
}

func (g *GrpcServer) GetUserIdentity(ctx context.Context, req *grpcPb.IdentityReq) (*grpcPb.IdentityResp, error) {
	var resp *grpcPb.IdentityResp
	user, err := g.authRepo.FindUserByID(ctx, int(req.GetUserID()))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	resp = &grpcPb.IdentityResp{
		UserID: int32(user.ID),
	}
	return resp, nil
}
