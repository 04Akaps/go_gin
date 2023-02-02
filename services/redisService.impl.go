package services

import (
	"context"
	"errors"
	"time"
	redis "github.com/redis/go-redis/v9"
)

type RedisServiceImpl struct {
	redis	*redis.Client
	ctx	context.Context
}

func NewRedisService(r *redis.Client, ctx context.Context) RedisService {
	return &RedisServiceImpl {
		redis : r,
		ctx : ctx,
	}
}

func (r *RedisServiceImpl) SetRedis(key string, value string) error {
	err := r.redis.Set(r.ctx, key, value, time.Hour).Err()

	if err != nil {
		return errors.New("redis Set Error")
	}

	return nil
}

func (r *RedisServiceImpl) GetRedis(key string) (string, error) {

	val, err := r.redis.Get(r.ctx, key).Result()

	if err != nil{
		return "", err
	}

	return val, nil
}

func (r *RedisServiceImpl) DeleteRedis(key string) error {
	err := r.redis.Del(r.ctx, key)

	if err != nil {
		return err
	}

	return nil
}



