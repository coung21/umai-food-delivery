package order

import (
	"context"
	"order-service/model"
)

type CacheRepository interface {
	ListCart(ctx context.Context, id int) ([]model.CartItem, error)
	// FindCartItem(ctx context.Context, uid int, mid string) bool
	// InsertCartItem(ctx context.Context, uid int, mid string) int
	// IncrCartItem(ctx context.Context, uid int, mid string, amount int) int
	InsertCartItem(ctx context.Context, uid int, mid int, amount int) int
	DelCartItem(ctx context.Context, uid int, mid ...int) int
}
