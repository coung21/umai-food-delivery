package usecase

import (
	"context"
	"errors"
	"umai-auth-service/model"
)

func (u *authUC) ShipperRegister(ctx context.Context, shipper *model.Shipper) (int, error) {
	// check if user exist

	_, err := u.authRepo.FindShipperByUserID(ctx, shipper.UserID)
	if err == nil {
		return 0, errors.New("user already exist")
	}

	// insert shipper
	sid, err := u.authRepo.InsertShipper(ctx, shipper)
	if err != nil {
		return 0, err
	}

	return sid, nil
}
