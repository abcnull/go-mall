package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall/biz/model"
	"go-mall/biz/service"
	"go-mall/pkg/status"
	"net/http"
)

func Login(c *gin.Context) {
	loginReq := new(model.LoginRequest)
	if err := c.ShouldBind(loginReq); err != nil {
		service.MakeResp(c, http.StatusOK, status.InvalidParam, err.Error())
		return
	}

	// 登录后后端基础的文本校验
	if sta, err := service.LoginCheck(c, loginReq); err != nil {
		service.MakeResp(c, http.StatusOK, sta, err.Error())
		return
	}

	// 登录
	loginResp, sta, err := service.LoginService(c, loginReq)
	if err != nil {
		service.MakeResp(c, http.StatusOK, sta, err.Error())
		return
	}

	// 返回成功信息
	service.MakeResp(c, http.StatusOK, status.Success, loginResp)
}
