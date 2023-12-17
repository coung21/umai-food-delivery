package order

import (
	"context"
	"order-service/model"
)

type Usecase interface {
	ListCartItems(ctx context.Context, id int) ([]model.CartItem, error)
}
