package usecase

import (
	"context"
	"order-service/model"
)

func (u *orderUC) ListCartItems(ctx context.Context, id int) ([]model.CartItem, error) {
	cartitem, err := u.OrderRepo.ListCart(ctx, id)

	if err != nil {
		return nil, err
	}

	return cartitem, nil
}
