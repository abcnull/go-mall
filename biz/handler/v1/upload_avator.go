package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall/biz/model"
	"go-mall/biz/service"
	"go-mall/pkg/status"
	"go-mall/pkg/util"
	"net/http"
)

func UploadAvatar(c *gin.Context) {
	// 从请求头拿到 token，并解析出 claims。后面为了获取其中的 userId
	token := c.GetHeader("Authorization")
	claims, err := util.ParseToken(token) // 解析 token
	if claims == nil || err != nil {
		service.MakeResp(c, http.StatusOK, status.AccessErr, nil)
		return
	}

	// 获取请求数据
	file, fileHeader, err := c.Request.FormFile("file") // todo: 不熟悉
	if err != nil {
		service.MakeResp(c, http.StatusOK, status.InvalidParam, nil)
		return
	}
	fileSize := fileHeader.Size
	uploadReq := new(model.UploadAvatarRequest)
	err = c.ShouldBind(uploadReq) // todo
	if err != nil {
		service.MakeResp(c, http.StatusOK, status.InvalidParam, nil)
		return
	}

	// 基本校验
	if sta, err := service.UploadAvatarCheck(c, uploadReq); err != nil {
		service.MakeResp(c, http.StatusOK, sta, err.Error())
		return
	}

	// 更新操作
	resp, sta, err := service.UploadAvatarService(c, claims.ID, file, fileSize)
	if err != nil {
		service.MakeResp(c, http.StatusOK, sta, err.Error())
		return
	}

	service.MakeResp(c, http.StatusOK, status.Success, resp)
}
