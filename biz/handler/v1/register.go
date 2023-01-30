package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall/biz/model"
	"go-mall/biz/service"
	"go-mall/pkg/status"
	"net/http"
)

func Register(c *gin.Context) {
	var registerReq model.RegisterRequest
	// 请求绑定 struct req，一般用 should 表示响应码自己最后决定
	if err := c.ShouldBind(&registerReq); err != nil {
		service.MakeResp(c, http.StatusOK, status.InvalidParam, err.Error())
		return
	}

	// 注册时候后端基础文本校验
	if sta, err := service.RegisterCheck(c, registerReq); err != nil {
		service.MakeResp(c, http.StatusOK, sta, err.Error())
		return
	}

	// 注册
	if sta, err := service.RegisterService(c, registerReq); err != nil {
		service.MakeResp(c, http.StatusOK, sta, err.Error())
		return
	}

	// 返回成功信息
	service.MakeResp(c, http.StatusOK, status.Success, nil)
}
