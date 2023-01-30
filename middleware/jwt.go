package middleware

import (
	"github.com/gin-gonic/gin"
	"go-mall/biz/service"
	"go-mall/pkg/status"
	"go-mall/pkg/util"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc { // todo: 函数为什么不直接这种形式，而是返回 gin.HandlerFunc 形式
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" { // token 为空
			service.MakeResp(c, http.StatusOK, status.AccessErr, nil)
			c.Abort()
			return
		} else {
			claims, err := util.ParseToken(token)
			if err != nil { // token 无权限
				service.MakeResp(c, http.StatusOK, status.AccessErr, nil)
				c.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt { // token 过期
				service.MakeResp(c, http.StatusOK, status.AccessErr, nil)
				c.Abort() // todo: abort 作用？
				return
			}
		}
		c.Next()
	}
}
