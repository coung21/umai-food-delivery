package repository

import (
	"context"
	"fmt"
	"order-service/model"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type CacheRepo struct {
	cdb *redis.Client
}

func NewCacheRepo(cdb *redis.Client) *CacheRepo {
	return &CacheRepo{cdb: cdb}
}

func (c *CacheRepo) ListCart(ctx context.Context, id int) ([]model.CartItem, error) {
	cartItems := []model.CartItem{}
	data, err := c.cdb.HGetAll(ctx, fmt.Sprintf("cart:%d", id)).Result()
	if err != nil {
		return nil, err
	}

	for key, value := range data {
		quantity, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}

		cartItems = append(cartItems, model.CartItem{ItemID: key, Quantity: quantity})
	}
	return cartItems, nil
}

// func (c *CacheRepo) FindCartItem(ctx context.Context, uid int, mid string) (string, error) {
// 	val, err := c.cdb.HGet(ctx, fmt.Sprintf("cart:%d", uid), mid).Result()
// 	if err == redis.Nil {
// 		return "", common.NotFound
// 	}

// 	return val, nil
// }

// func (c *CacheRepo) InsertCartItem(ctx context.Context, uid int, mid string) int {
// 	val := c.cdb.HSet(ctx, fmt.Sprintf("cart:%d", uid), mid, 1).Val()

// 	return int(val)
// }

func (c *CacheRepo) InsertCartItem(ctx context.Context, uid int, mid string, amount int) int {
	val := c.cdb.HIncrBy(ctx, fmt.Sprintf("cart:%d", uid), mid, int64(amount)).Val()

	if val <= 0 {
		c.DelCartItem(ctx, uid, mid)
		return 0
	}
	return int(val)
}

func (c *CacheRepo) DelCartItem(ctx context.Context, uid int, mid ...string) int {
	val := c.cdb.HDel(ctx, fmt.Sprintf("cart:%d", uid), mid...).Val()
	return int(val)
}
