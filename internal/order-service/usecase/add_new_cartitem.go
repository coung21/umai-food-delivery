package usecase

import (
	"context"
	"log"
)

func (c *orderUC) AddItemToCart(ctx context.Context, uid int, mid string, amount int) bool {
	existed := c.OrderRepo.FindCartItem(ctx, uid, mid)
	if existed {
		_, err := c.OrderRepo.IncrCartItem(ctx, uid, mid, amount)
		if err != nil {
			log.Fatal(err)
			return false
		}
		return true
	} else {
		ok := c.OrderRepo.InsertCartItem(ctx, uid, mid)
		if ok {
			return true
		}
		return false
	}
}
