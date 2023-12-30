package repository

import (
	"common"
	"context"
	"encoding/json"
	"fmt"
	"menu-service/model"
	"time"

	"strconv"

	"github.com/redis/go-redis/v9"
)

type cacheMenuRepo struct {
	cdb *redis.Client
}

func NewCacheRepo(addr, password string, db int) *cacheMenuRepo {
	cdb := redis.NewClient(&redis.Options{Addr: addr, Password: password, DB: db})
	return &cacheMenuRepo{cdb: cdb}
}

func (c *cacheMenuRepo) Set(ctx context.Context, id int, mitem *model.MenuItem, ttl time.Duration) error {
	mitemBytes, err := json.Marshal(mitem)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("menu:%d", id)
	return c.cdb.Set(ctx, key, mitemBytes, ttl).Err()
}

func (c *cacheMenuRepo) Get(ctx context.Context, id int) (*model.MenuItem, error) {
	key := fmt.Sprintf("menu:%d", id)

	val, err := c.cdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, common.ErrMissCache
		} else {
			return nil, err
		}
	}
	var mitem model.MenuItem
	err = json.Unmarshal([]byte(val), &mitem)

	return &mitem, nil
}

func (c *cacheMenuRepo) Del(ctx context.Context, id int) error {
	key := fmt.Sprintf("menu:%d", id)
	return c.cdb.Del(ctx, key).Err()
}

func (c *cacheMenuRepo) SetFavorite(ctx context.Context, uid int, mid int) error {
	key := fmt.Sprintf("favorites:%d", uid)
	return c.cdb.SAdd(ctx, key, mid).Err()
}

func (c *cacheMenuRepo) ListFavorites(ctx context.Context, uid int) ([]int, error) {
	key := fmt.Sprintf("favorites:%d", uid)

	mids, err := c.cdb.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var result []int
	for _, mid := range mids {
		midInt, err := strconv.Atoi(mid)
		if err != nil {
			return nil, err
		}
		result = append(result, midInt)
	}

	return result, nil
}

func (c *cacheMenuRepo) DelFavorite(ctx context.Context, uid int, mid int) error {
	key := fmt.Sprintf("favorites:%d", uid)
	return c.cdb.SRem(ctx, key, mid).Err()
}

func (c *cacheMenuRepo) GetFavorite(ctx context.Context, uid int, mid int) (bool, error) {
	key := fmt.Sprintf("favorites:%d", uid)
	return c.cdb.SIsMember(ctx, key, mid).Result()
}
