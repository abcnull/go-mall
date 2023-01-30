package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall/biz/model"
	"go-mall/biz/service"
	"go-mall/pkg/status"
	"go-mall/pkg/util"
	"net/http"
)

func UpdateUser(c *gin.Context) {
	// 从请求头拿到 token，并解析出 claims。后面为了获取其中的 userId
	token := c.GetHeader("Authorization")
	claims, err := util.ParseToken(token) // 解析 token
	if claims == nil || err != nil {
		service.MakeResp(c, http.StatusOK, status.AccessErr, nil)
		return
	}

	updateUserReq := new(model.UpdateUserRequest)
	if err := c.ShouldBind(updateUserReq); err != nil { // todo: 绑定 form 入参和绑定 json 入参有什么区别吗
		service.MakeResp(c, http.StatusOK, status.InvalidParam, err.Error())
		return
	}

	// 后端基础的文本校验
	if sta, err := service.UpdateUserCheck(c, updateUserReq); err != nil {
		service.MakeResp(c, http.StatusOK, sta, err.Error())
		return
	}

	// 依据 token 解析出来的 claims 中的 userId 来更新用户
	sta, err := service.UpdateUserService(c, updateUserReq, claims.ID)
	if err != nil {
		service.MakeResp(c, http.StatusOK, sta, err.Error())
		return
	}

	// 返回成功信息
	service.MakeResp(c, http.StatusOK, status.Success, nil)
}
