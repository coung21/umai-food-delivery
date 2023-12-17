package order

import (
	"context"
	"order-service/model"
)

type CacheRepository interface {
	ListCart(ctx context.Context, id int) ([]model.CartItem, error)
	FindCartItem(ctx context.Context, uid int, mid string) *int
	InsertCartItem(ctx context.Context, uid int, mid string) bool
	IncrCartItem(ctx context.Context, uid int, mid string, amount int) (int, error)
	DelCartItem(ctx context.Context, uid int, mid string) bool
}
