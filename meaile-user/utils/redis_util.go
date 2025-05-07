package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

type RedisClient struct {
	RedisClient *redis.Client
}

func (r *RedisClient) Set(ctx context.Context, key, value string) bool {
	result, err := r.RedisClient.Set(ctx, key, value, 0).Result()
	if err != nil {
		zap.S().Error("----Redis Set Failed----" + key + ":" + value)
		return false
	}
	return result == "OK"
}

func (r *RedisClient) SetEx(ctx context.Context, key, value string, ex time.Duration) bool {
	result, err := r.RedisClient.Set(ctx, key, value, ex).Result()
	if err != nil {
		zap.S().Error("----Redis Set Failed----" + key + ":" + value)
		return false
	}
	return result == "OK"
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, bool) {
	result, err := r.RedisClient.Get(ctx, key).Result()
	if err != nil {
		zap.S().Error("----Redis Get Failed----" + key)
		return "", false
	}
	return result, true
}
func (r *RedisClient) GetSet(ctx context.Context, key, value string) (string, bool) {
	oldValue, err := r.RedisClient.GetSet(ctx, key, value).Result()
	if err != nil {
		zap.S().Error("----Redis GetSet Failed----" + key)
		return "", false
	}
	return oldValue, true
}
func (r *RedisClient) Incr(ctx context.Context, key string) (int64, bool) {
	result, err := r.RedisClient.Incr(ctx, key).Result()
	if err != nil {
		zap.S().Error("----Redis Incr Failed----" + key)
		return result, false
	}
	return result, true
}
func (r *RedisClient) IncrBy(ctx context.Context, key string, incr int64) (int64, bool) {
	result, err := r.RedisClient.IncrBy(ctx, key, incr).Result()
	if err != nil {
		zap.S().Error("----Redis IncrBy Failed----" + key)
		return result, false
	}
	return result, true
}
func (r *RedisClient) IncrByFloat(ctx context.Context, key string, incrFloat float64) (float64, bool) {
	result, err := r.RedisClient.IncrByFloat(ctx, key, incrFloat).Result()
	if err != nil {
		zap.S().Error("----Redis IncrByFloat Failed----" + key)
		return result, false
	}
	return result, true
}
