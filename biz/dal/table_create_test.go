package dal

import (
	"context"
	"go-mall/biz/dal/entity"
	"go-mall/client/mysql"
	"go-mall/conf"
	"testing"
)

// 通过 gorm 的 AutoMigration 函数依据 struct 字段建表，前提是 gorm 连接上数据库
func TestCreateTable(t *testing.T) {
	// 读取配置
	conf.InitConf("../../conf/conf.ini")
	// 初始化 mysql 连接数据库
	mysql.InitMysql()
	// 建表，mysql 数据库连接后主要使用了 AutoMigration 函数来自动依据 struct 建表
	err := mysql.Client(context.Background()).Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&entity.User{},
			&entity.Category{},
			&entity.Order{},
			&entity.Notice{},
			&entity.Product{},
			&entity.ProductImg{},
			&entity.Address{},
			&entity.Admin{},
			&entity.Carousel{},
			&entity.Cart{},
			&entity.Favorite{},
		)
	if err != nil {
		panic(err)
	}
}
