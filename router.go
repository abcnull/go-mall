package main

import (
	"github.com/gin-gonic/gin"
	"go-mall/biz/handler"
	v1 "go-mall/biz/handler/v1"
	"go-mall/middleware"
	"net/http"
)

// Register 模块
func Register() *gin.Engine {
	// gin Engine
	r := gin.Default()

	// 中间件-跨域
	// todo: 跨域问题？
	r.Use(middleware.Cors())

	// 加载静态文件路径
	// todo: ???
	r.StaticFS("/static", http.Dir("./static"))

	// ping
	r.GET("/ping", handler.Ping)

	// 路由组
	v1Group := r.Group("/api/v1")
	{
		// 注册
		v1Group.POST("/register", v1.Register)
		// 登录
		v1Group.POST("/login", v1.Login)

		// 获取轮播图
		v1Group.GET("/carousels", v1.Carousel)

		/* 用户模块 */
		authedGroup := v1Group.Group("/user")
		authedGroup.Use(middleware.JWT()) // token 鉴权先兜住没有没有权限的直接给返回了
		{
			/* 用户相关 */
			// 更新用户
			authedGroup.POST("/update", v1.UpdateUser) // todo
			// 上传头像
			authedGroup.POST("/uploadavatar", v1.UploadAvatar) // todo: 可以考虑添加背景图
			// 邮件发送
			authedGroup.POST("/send-email", v1.SendEmail) // todo: 邮箱相关没太看懂，需要好好看下
			// 验证邮箱
			authedGroup.POST("/validate-email", v1.ValidateEmail) // todo: 邮箱相关没太看懂，需要好好看下
			// 显示用户金额
			authedGroup.POST("/money", v1.ShowMoney)
		}

		/* 地址模块 */
		addressGroup := v1Group.Group("/address")
		addressGroup.Use(middleware.JWT())
		{
			// 查询地址列表
			addressGroup.GET("/list", v1.AddressList)
			// 查询具体地址
			addressGroup.GET("/:id", v1.AddressDetail)
			// 创建地址
			addressGroup.POST("/create", v1.CreateAddress)
			// 编辑地址
			addressGroup.POST("/update", v1.UpdateAddress)
			// 删除地址，可批量删除
			addressGroup.POST("/delete", v1.DeleteAddress)
		}

		/* 商品模块 */
		productGroup := v1Group.Group("/product")
		// 查询商品列表
		productGroup.GET("/list", v1.ProductList) // todo: 这里没有写在 service 里头了
		// 查询具体商品
		productGroup.GET("/:id", v1.ProductDetail)
		// 创建商品时依据商品查询种类
		productGroup.GET("/category", v1.AutoGainCategory) // todo: 发布商品信息时能即时性的获得商品的品类 id
		productGroup.Use(middleware.JWT())
		{
			/* 商品相关 */
			// 创建商品
			productGroup.POST("/create", v1.CreateProduct)
			// 商品更新
			productGroup.POST("/update", v1.UpdateProduct)
		}

		/* 收藏夹模块 */
		favoriteGroup := v1Group.Group("/favor")
		favoriteGroup.Use(middleware.JWT())
		{
			// 展示收藏夹
			favoriteGroup.GET("/list", v1.FavorList)
			// 新增收藏夹商品
			favoriteGroup.POST("/add", v1.AddFavor)
			// 删除收藏夹商品
			favoriteGroup.POST("/delete_in", v1.DeleteInFavor) // todo: 路由有下划线的写法吗?
			// 创建收藏夹
			favoriteGroup.POST("/create", v1.CreateFavor) // todo: 后续补充
			// 删除收藏夹
			favoriteGroup.POST("/delete", v1.DeleteFavor) // todo: 后续补充
		}

		/* 订单模块 */
		orderGroup := v1Group.Group("/order")
		orderGroup.Use(middleware.JWT())
		{
			// 创建订单
			orderGroup.POST("/create", v1.CreateOrder)
			//
		}
	}

	return r
}
