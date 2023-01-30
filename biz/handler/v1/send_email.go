package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall/biz/model"
	"go-mall/biz/service"
	"go-mall/pkg/status"
	"go-mall/pkg/util"
	"net/http"
)

func SendEmail(c *gin.Context) {
	// 解析 token
	token := c.GetHeader("Authorization")
	claims, err := util.ParseToken(token)
	if err != nil || claims == nil {
		service.MakeResp(c, http.StatusOK, status.AccessErr, err.Error())
		return
	}

	// 绑定入参
	sendEmailReq := new(model.SendEmailRequest)
	if err := c.ShouldBind(sendEmailReq); err != nil {
		service.MakeResp(c, http.StatusOK, status.InvalidParam, err.Error())
		return
	}

	// 后端校验格式正确
	if sta, err := service.SendEmailCheck(c, sendEmailReq); err != nil {
		service.MakeResp(c, http.StatusOK, sta, err.Error())
		return
	}

	// 发送邮件
	sta, err := service.SendEmailService(c, sendEmailReq, claims.ID)
	if err != nil {
		service.MakeResp(c, http.StatusOK, sta, err.Error())
		return
	}

	service.MakeResp(c, http.StatusOK, status.Success, nil)
}
