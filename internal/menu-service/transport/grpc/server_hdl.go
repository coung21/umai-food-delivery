package grpc

import (
	"context"
	"encoding/json"
	"menu-service/transport/grpc/grpcPb"
)

func (g *GrpcServer) GetMenuItem(ctx context.Context, req *grpcPb.GetMenuItemReq) (*grpcPb.GetMenuItemRes, error) {
	var resp *grpcPb.GetMenuItemRes

	mitem, err := g.menuRepo.FindMenuItemByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	jsonMitem, err := json.Marshal(mitem)
	if err != nil {
		return nil, err
	}

	resp = &grpcPb.GetMenuItemRes{
		Data: string(jsonMitem),
	}
	return resp, nil
}
