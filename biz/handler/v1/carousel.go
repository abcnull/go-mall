package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall/biz/model"
	"go-mall/biz/service"
	"go-mall/pkg/status"
	"net/http"
)

func Carousel(c *gin.Context) {
	// 绑定获取轮播头 todo: 因为没有入参，不绑定是否可行？
	carouselReq := new(model.CarouselRequest)
	if err := c.ShouldBind(carouselReq); err != nil {
		service.MakeResp(c, http.StatusOK, status.InvalidParam, err.Error())
		return
	}

	// 获取轮播图
	sta, err, resp := service.ListCarouselService(c)
	if err != nil {
		service.MakeResp(c, http.StatusOK, status.Error, err.Error())
		return
	}

	service.MakeResp(c, http.StatusOK, sta, resp)
}
