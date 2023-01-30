package middleware

import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// todo: 跨域中间件
	}
}
