package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-mall/biz/dal/entity"
	"go-mall/biz/model"
	"go-mall/biz/service"
	"go-mall/client/mysql"
	"go-mall/pkg/status"
	"net/http"
	"strconv"
)

func ProductDetail(c *gin.Context) {
	// 绑定入参
	productDetailReq := new(model.ProductDetailRequest)
	if err := c.ShouldBind(productDetailReq); err != nil {
		service.MakeResp(c, http.StatusOK, status.InvalidParam, err.Error())
		return
	}

	// 查询商品详细数据
	resp, sta, err := showProductDetail(c, productDetailReq, c.Param("id"))
	if err != nil {
		service.MakeResp(c, http.StatusOK, sta, err.Error())
		return
	}

	service.MakeResp(c, http.StatusOK, sta, resp)
}

// 展现商品详细信息
func showProductDetail(c *gin.Context, request *model.ProductDetailRequest, id string) (*model.ProductDetailResponse, *status.Status, error) {
	mysqlCli := mysql.Client(c)

	// 通过 product_id 获取指定商品详细信息
	product := new(entity.Product)
	if err := mysqlCli.Model(&entity.Product{}).Where("id = ?", id).Find(product).Error; err != nil {
		return nil, status.Error, err
	} else if product.ID == 0 {
		return nil, status.Success, errors.New("数据库无数据")
	}

	// 通过 product_id 获取指定商品的商品图片
	productImgSli := make([]entity.ProductImg, 0)
	if err := mysqlCli.Model(&entity.ProductImg{}).Where("product_id = ?", id).Find(&productImgSli).Error; err != nil {
		return nil, status.Error, err
	}
	productImgStrSli := make([]string, 0)
	for _, v := range productImgSli {
		productImgStrSli = append(productImgStrSli, v.ImgPath)
	}

	originPrice, err := strconv.ParseFloat(product.Price, 64)
	if err != nil {
		return nil, status.Error, err
	}
	discountPrice, err := strconv.ParseFloat(product.DiscountPrice, 64)
	if err != nil {
		return nil, status.Error, err
	}
	resp := &model.ProductDetailResponse{
		Product: model.Product{
			Id:                 product.ID,
			CategoryId:         product.CategoryId,
			Title:              product.Title,
			Info:               product.Info,
			CoverImgPath:       product.ImgPath,
			AllImgPath:         productImgStrSli,
			OriginPrice:        originPrice,
			DiscountPrice:      discountPrice,
			Locate:             product.Locate,
			Freight:            model.FreightType(product.Freight),
			BossId:             product.BossID,
			BossName:           product.BossName,
			BossAvatar:         product.BossAvatar,
			ExposureTimes:      product.ExposureTimes,
			ClickTimes:         product.ClickTimes,
			CommunicationTimes: product.CommunicationTimes,
		},
	}

	return resp, status.Success, nil
}
