package main

import (
	"go-mall/client/mysql"
	"go-mall/client/redis"
	"go-mall/conf"
)

func main() {
	// 初始化配置文件，路径以当前路径为依托
	conf.InitConf("conf/conf.ini")

	// 初始化 mysql
	mysql.InitMysql()

	// 初始化 redis
	redis.InitRedis()

	// 路由注册
	r := Register()
	if err := r.Run(conf.HttpPort); err != nil {
		panic("路由没成功")
	}
}
