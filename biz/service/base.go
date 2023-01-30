package service

import (
	"github.com/gin-gonic/gin"
	"go-mall/pkg/status"
)

func MakeResp(c *gin.Context, httpStatus int, sta *status.Status, data interface{}) {
	c.JSON(httpStatus, status.StandardResponse{
		Sta:  sta,
		Data: data,
	})
}
