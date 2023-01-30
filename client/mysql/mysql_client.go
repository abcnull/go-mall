package mysql

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-mall/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

var dbClient *gorm.DB

func InitMysql() {
	// 日志
	// todo: 这里怎么用?
	var gormLogger logger.Interface
	// todo: 这里怎么用？
	if gin.Mode() == "debug" {
		gormLogger = logger.Default.LogMode(logger.Info)
	} else {
		gormLogger = logger.Default
	}

	// 创建 db 连接
	// todo: 两个gorm有什么区别
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: conf.ReadDSN, // 主数据库
	}), &gorm.Config{
		Logger: gormLogger, // 日志
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 列不加 s
		},
	})
	if err != nil {
		panic(err)
	}

	// 连接池配置
	sqlDB, _ := db.DB()
	// todo: 怎么设置？
	sqlDB.SetMaxIdleConns(20)  // 设置链接池
	sqlDB.SetMaxOpenConns(100) // 设置最大连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	dbClient = db

	// 主从配置
	// todo: 什么是主从配置？
	_ = dbClient.Use(dbresolver.Register(
		dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(conf.WriteDSN)},
			Replicas: []gorm.Dialector{mysql.Open(conf.ReadDSN), mysql.Open(conf.WriteDSN)}, // 复制
			Policy:   dbresolver.RandomPolicy{},                                             // 负载均衡的策略
		},
	))
}

func Client(ctx context.Context) *gorm.DB {
	return dbClient.WithContext(ctx)
}
