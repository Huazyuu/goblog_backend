package core

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"time"
)

func InitRedis() *redis.Client {
	return InitRedisDB(global.Config.Redis.DB)
}
func InitRedisDB(db int) *redis.Client {
	redisConf := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.GetAddr(),
		Password: redisConf.Password,
		DB:       db,
		PoolSize: redisConf.PoolSize,
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		global.Logger.Errorf("连接redis失败%s", redisConf.GetAddr())
		return nil
	}
	logrus.Infof(fmt.Sprintf("[%s] redis connects success", redisConf.GetAddr()))
	return rdb
}
