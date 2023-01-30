package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall/biz/model"
	"go-mall/biz/service"
	"go-mall/pkg/status"
	"go-mall/pkg/util"
	"net/http"
)

// CreateProduct 创建商品
func CreateProduct(c *gin.Context) {
	// 拿到 claims
	claims, err := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		service.MakeResp(c, http.StatusOK, status.AccessErr, err.Error())
		return
	}

	// 拿到文件 todo: 这里没太明白是什么操作
	form, err := c.MultipartForm()
	if err != nil {
		service.MakeResp(c, http.StatusOK, status.InvalidParam, err.Error())
		return
	}
	files := form.File["file"]

	// 拿到入参
	createProductReq := new(model.CreateProductRequest)
	if err := c.ShouldBind(createProductReq); err != nil {
		service.MakeResp(c, http.StatusOK, status.InvalidParam, err.Error())
		return
	}

	// 判断入参是否正常
	service.CreateProductServiceCheck(c, createProductReq)

	// 创建商品
	resp, sta, err := service.CreateProductService(c, claims.ID, createProductReq, files)
	if err != nil {
		service.MakeResp(c, http.StatusOK, sta, err.Error())
		return
	}

	service.MakeResp(c, http.StatusOK, status.Success, resp)
}
