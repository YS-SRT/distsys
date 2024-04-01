package utils

import (
	"context"
	"distsys/global"
	"errors"
	"time"

	"go.uber.org/zap"
)

func GetByRedisSet(key string) (any, error) {
	if value, err := global.GRedisCli.Get(context.TODO(), key).Result(); err != nil {
		return nil, err

	} else {
		return value, nil
	}
}

func AddByRedisSet(key string, value string, expiredTime time.Time) error {

	return InvokeRedisSet(key, value, time.Until(expiredTime))
}

func InvokeRedisSet(key string, value any, expiration time.Duration) error {

	if result, err := global.GRedisCli.Set(context.TODO(), key, value, expiration).Result(); err != nil {

		zap.L().Error("set object in Redis failed", zap.Error(err))
		return err

	} else {

		if result == "0" {

			return errors.New("no object set")
		}
	}

	return nil
}

func AddByRedisHSet(key string, vals map[string]string) error {
	ctx := context.TODO()
	for k, v := range vals {

		if err := global.GRedisCli.HSet(ctx, key, k, v).Err(); err != nil {
			zap.L().Error("HSet object failed", zap.Error(err))
			return err
		}
	}

	return nil
}

func GetByRedisHSet(key string) map[string]string {

	return global.GRedisCli.HGetAll(context.TODO(), key).Val()
}
