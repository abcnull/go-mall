package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall/biz/model"
	"go-mall/biz/service"
	"go-mall/pkg/status"
	"go-mall/pkg/util"
	"net/http"
)

func ShowMoney(c *gin.Context) {
	// 能到 claims
	claims, err := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		service.MakeResp(c, http.StatusOK, status.AccessErr, err.Error())
		return
	}

	// 绑定请求参数
	showMoneyReq := new(model.ShowMoneyRequest)
	if err := c.ShouldBind(showMoneyReq); err != nil {
		service.MakeResp(c, http.StatusOK, status.InvalidParam, err.Error())
		return
	}

	// 实际 show money
	sta, err, resp := service.ShowMoneyService(c, showMoneyReq, claims.ID)
	if err != nil {
		service.MakeResp(c, http.StatusOK, status.Error, err.Error())
		return
	}

	service.MakeResp(c, http.StatusOK, sta, resp)
}
