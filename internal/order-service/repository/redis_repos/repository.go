package repository

import "github.com/redis/go-redis/v9"

type CacheRepo struct {
	cdb *redis.Client
}

func NewCacheRepo(cdb *redis.Client) *CacheRepo {
	return &CacheRepo{cdb: cdb}
}

func (c *CacheRepo) AddCart()
