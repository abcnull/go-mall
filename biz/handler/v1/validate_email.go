package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall/biz/model"
	"go-mall/biz/service"
	"go-mall/pkg/status"
	"go-mall/pkg/util"
	"net/http"
)

func ValidateEmail(c *gin.Context) {
	// 获取 claims
	_, err := util.ParseToken(c.GetHeader("Authorization")) // todo
	if err != nil {
		service.MakeResp(c, http.StatusOK, status.AccessErr, err.Error())
		return
	}

	// 绑定
	ValidateEmailReq := new(model.ValidateEmailRequest)
	if err := c.ShouldBind(ValidateEmailReq); err != nil {
		service.MakeResp(c, http.StatusOK, status.InvalidParam, err.Error())
		return
	}

	// 校验邮箱有效
	service.ValidateEmailService(c, ValidateEmailReq.EmailToken)

}
