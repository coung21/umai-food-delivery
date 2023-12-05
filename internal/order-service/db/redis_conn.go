package db

import "github.com/redis/go-redis/v9"

func RedisConn(adrr, passowrd string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{Addr: adrr, Password: passowrd, DB: db})
}
