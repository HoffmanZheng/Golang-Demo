package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/cache"
	_ "github.com/beego/beego/v2/client/cache/redis"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

var redisClient cache.Cache
var enableRedis, _ = web.AppConfig.Bool("enableRedis")
var redisTime, _ = web.AppConfig.Int("redisTime")
var YzmClient cache.Cache
var logger = logs.GetBeeLogger()

func init() {
	if enableRedis {
		redisKey, _ := web.AppConfig.String("redisKey")
		redisConn, _ := web.AppConfig.String("redisConn")
		redisDbNum, _ := web.AppConfig.String("redisDbNum")
		redisPwd, _ := web.AppConfig.String("redisPwd")
		config := map[string]string{
			"key":      redisKey,
			"conn":     redisConn,
			"dbNum":    redisDbNum,
			"password": redisPwd,
		}
		bytes, _ := json.Marshal(config)

		redisClient, err = cache.NewCache("redis", string(bytes))
		YzmClient, _ = cache.NewCache("redis", string(bytes))
		if err != nil {
			logger.Error("failed to connect Redis")
		} else {
			logger.Info("connect to Redis successfully!")
		}
	}
}

type cacheDb struct{}

var CacheDb = &cacheDb{}

func (c cacheDb) Set(key string, value interface{}) {
	if enableRedis {
		bytes, _ := json.Marshal(value)
		redisClient.Put(key, string(bytes), time.Second*time.Duration(redisTime))
	}
}

func (c cacheDb) Get(key string, obj interface{}) bool {
	if enableRedis {
		if redisStr := redisClient.Get(key); redisStr != nil {
			fmt.Println("read data from redis ...")
			redisValue, ok := redisStr.([]uint8)
			if !ok {
				fmt.Println("failed to get data from redis")
				return false
			}
			json.Unmarshal([]byte(redisValue), obj)
			return true
		}
		return false
	}
	return false
}
