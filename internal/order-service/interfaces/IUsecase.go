package order

import (
	"context"
	"order-service/model"
)

type Usecase interface {
	ListCartItems(ctx context.Context, id int) ([]model.CartItem, error)
	ModifyCart(ctx context.Context, uid int, mid int, amount int) int
	DeleteItemFromCart(ctx context.Context, uid int, items []int) int
}
