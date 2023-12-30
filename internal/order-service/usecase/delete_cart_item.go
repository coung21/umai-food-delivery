package usecase

import (
	"context"
)

func (u *orderUC) DeleteItemFromCart(ctx context.Context, uid int, items []int) int {
	result := u.OrderRepo.DelCartItem(ctx, uid, items...)
	return result
}
