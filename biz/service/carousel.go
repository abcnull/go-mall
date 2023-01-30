package service

import (
	"github.com/gin-gonic/gin"
	"go-mall/biz/dal/entity"
	"go-mall/biz/model"
	"go-mall/client/mysql"
	"go-mall/pkg/status"
)

func ListCarouselService(c *gin.Context) (*status.Status, error, *model.CarouselResponse) {
	mysqlCli := mysql.Client(c)

	entityItems := new([]entity.Carousel)
	if err := mysqlCli.Model(&entity.Carousel{}).Find(entityItems).Error; err != nil {
		return status.Error, err, nil
	}

	modelItems := make([]model.CarouselItem, 0)
	for _, v := range *entityItems {
		modelItems = append(modelItems, model.CarouselItem{
			Id:        v.ID,
			ImgPath:   v.ImgPath,
			ProductId: v.ProductID,
			CreateAt:  v.CreatedAt.Unix(),
		})
	}
	return status.Success, nil, &model.CarouselResponse{
		Items: modelItems,
	}
}
