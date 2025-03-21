package utils

import (
	"context"
	"go.uber.org/zap"
	"meaile-server/meaile-user/global"
	"time"
)

type RedisUtil struct {
}

func (r *RedisUtil) Set(ctx context.Context, key, value string) bool {
	result, err := global.RedisClient.Set(ctx, key, value, 0).Result()
	if err != nil {
		zap.S().Error("----Redis Set Failed----" + key + ":" + value)
		return false
	}
	return result == "OK"
}

func (r *RedisUtil) SetEx(ctx context.Context, key, value string, ex time.Duration) bool {
	result, err := global.RedisClient.Set(ctx, key, value, ex).Result()
	if err != nil {
		zap.S().Error("----Redis Set Failed----" + key + ":" + value)
		return false
	}
	return result == "OK"
}

func (r *RedisUtil) Get(ctx context.Context, key string) (string, bool) {
	result, err := global.RedisClient.Get(ctx, key).Result()
	if err != nil {
		zap.S().Error("----Redis Get Failed----" + key)
		return "", false
	}
	return result, true
}
func (r *RedisUtil) GetSet(ctx context.Context, key, value string) (string, bool) {
	oldValue, err := global.RedisClient.GetSet(ctx, key, value).Result()
	if err != nil {
		zap.S().Error("----Redis GetSet Failed----" + key)
		return "", false
	}
	return oldValue, true
}
func (r *RedisUtil) Incr(ctx context.Context, key string) (int64, bool) {
	result, err := global.RedisClient.Incr(ctx, key).Result()
	if err != nil {
		zap.S().Error("----Redis Incr Failed----" + key)
		return result, false
	}
	return result, true
}
func (r *RedisUtil) IncrBy(ctx context.Context, key string, incr int64) (int64, bool) {
	result, err := global.RedisClient.IncrBy(ctx, key, incr).Result()
	if err != nil {
		zap.S().Error("----Redis IncrBy Failed----" + key)
		return result, false
	}
	return result, true
}
func (r *RedisUtil) IncrByFloat(ctx context.Context, key string, incrFloat float64) (float64, bool) {
	result, err := global.RedisClient.IncrByFloat(ctx, key, incrFloat).Result()
	if err != nil {
		zap.S().Error("----Redis IncrByFloat Failed----" + key)
		return result, false
	}
	return result, true
}
