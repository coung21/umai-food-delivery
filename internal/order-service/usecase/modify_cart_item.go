package usecase

import (
	"context"
)

func (c *orderUC) ModifyCart(ctx context.Context, uid int, mid string, amount int) int {
	val := c.OrderRepo.InsertCartItem(ctx, uid, mid, amount)
	return val
}
