package redis

import (
	"github.com/go-redis/redis"
	"go-mall/conf"
	"strconv"
)

var (
	RedisCli *redis.Client

	RedisAddr   string
	RedisPwd    string
	RedisDb     string
	RedisDbName string
)

func InitRedis() {
	RedisAddr = conf.RedisAddr
	RedisPwd = conf.RedisPwd
	RedisDb = conf.RedisDb
	RedisDbName = conf.RedisDbName

	db, _ := strconv.ParseUint(RedisDbName, 10, 64) // todo: 为什么要做这样的转化？
	RedisCli = redis.NewClient(&redis.Options{
		Addr: RedisAddr,

		// 有密码写密码

		DB: int(db),
	})

	_, err := RedisCli.Ping().Result() // 心跳检测 // todo: 意义是什么？
	if err != nil {
		panic(err)
	}
}
