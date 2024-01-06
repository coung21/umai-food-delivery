package usecase

import (
	"common"
	"context"
	"umai-auth-service/model"
)

func (u *authUC) UpdateShipper(ctx context.Context, id int, upd *model.ShipperUpdate) (*model.Shipper, error) {
	// check if shipper exist
	olddata, err := u.authRepo.FindShipperByID(ctx, id)
	if err != nil {
		if err == common.NotFound {
			return nil, common.NotFound
		}
		return nil, err
	}

	// update shipper
	shipper, err := u.authRepo.UpdateShipper(ctx, olddata, upd)
	if err != nil {
		return nil, err
	}

	return shipper, nil
}
