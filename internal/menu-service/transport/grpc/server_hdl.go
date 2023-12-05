package grpc

import (
	"common"
	"context"
	"encoding/json"
	"menu-service/transport/grpc/grpcPb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g *GrpcServer) GetMenuItem(ctx context.Context, req *grpcPb.GetMenuItemReq) (*grpcPb.GetMenuItemRes, error) {
	var resp *grpcPb.GetMenuItemRes

	mitem, err := g.menuRepo.FindMenuItemByID(ctx, req.GetId())
	if err != nil {
		if err == common.NotFound {
			return nil, status.Error(codes.NotFound, "Given Id Not Found")
		} else if err == common.BadQueryParams {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid Agrument: %v", req.GetId())
		}
		return nil, status.Error(codes.Internal, "Internal")
	}

	jsonMitem, err := json.Marshal(mitem)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal")
	}

	resp = &grpcPb.GetMenuItemRes{
		Data: string(jsonMitem),
	}
	return resp, nil
}
