package models

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var redisClient cache.Cache
var enableRedis, _ = beego.AppConfig.Bool("enableRedis")
var redisTime, _ = beego.AppConfig.Int("redisTime")
var YzmClient cache.Cache
var logger = logs.GetBeeLogger()

func init() {
	if enableRedis {
		redisKey := beego.AppConfig.String("redisKey")
		redisConn := beego.AppConfig.String("redisConn")
		redisDbNum := beego.AppConfig.String("redisDbNum")
		redisPwd := beego.AppConfig.String("redisPwd")
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

func (c cacheDb) Set(ctx context.Context, key string, value interface{}) {
	if enableRedis {
		bytes, _ := json.Marshal(value)
		redisClient.Put(key, string(bytes), time.Second*time.Duration(redisTime))
	}
}

func (c cacheDb) Get(ctx context.Context, key string, obj interface{}) bool {
	if enableRedis {
		if redisStr := redisClient.Get(key); err != nil {
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
