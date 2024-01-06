package usecase

import (
	"context"
	"umai-auth-service/model"
)

func (u *authUC) GetShipper(ctx context.Context, id int) (*model.Shipper, error) {
	// check if shipper exist
	shipper, err := u.authRepo.FindShipperByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return shipper, nil
}
