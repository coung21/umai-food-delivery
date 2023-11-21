package repository

import (
	"common"
	"context"
	"encoding/json"
	"fmt"
	"time"
	"umai-auth-service/model"

	"github.com/redis/go-redis/v9"
)

type cacheAuthRepo struct {
	cdb *redis.Client
}

func NewCacheAuthRepo(cdb *redis.Client) *cacheAuthRepo {
	return &cacheAuthRepo{cdb: cdb}
}

func (c *cacheAuthRepo) Set(ctx context.Context, id int, res *model.Restaurant, ttl time.Duration) error {
	resBytes, err := json.Marshal(res)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("restaurant:%d", id)
	return c.cdb.Set(ctx, key, resBytes, ttl).Err()
}

func (c *cacheAuthRepo) Get(ctx context.Context, id int) (*model.Restaurant, error) {
	key := fmt.Sprintf("restaurant:%d", id)

	val, err := c.cdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, common.ErrMissCache
		} else {
			return nil, err
		}
	}
	var res model.Restaurant
	err = json.Unmarshal([]byte(val), &res)

	return &res, nil
}

func (c *cacheAuthRepo) Del(ctx context.Context, id int) error {
	key := fmt.Sprintf("restaurant:%d", id)
	return c.cdb.Del(ctx, key).Err()
}
