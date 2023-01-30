package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall/biz/dal/entity"
	"go-mall/biz/model"
	"go-mall/biz/service"
	"go-mall/client/mysql"
	"go-mall/pkg/status"
	"net/http"
	"strconv"
)

func ProductList(c *gin.Context) {
	// 绑定入参
	productListReq := new(model.ProductListRequest)
	if err := c.ShouldBind(productListReq); err != nil {
		service.MakeResp(c, http.StatusOK, status.InvalidParam, err.Error())
		return
	}
	if productListReq.PageSize == 0 {
		productListReq.PageSize = 5 // 默认 page_size: 5
	}
	if productListReq.PageNum == 0 {
		productListReq.PageNum = 1 // 默认 page_num: 1
	}

	// 展示商品数据
	resp, sta, err := showProductList(c, productListReq)
	if err != nil {
		service.MakeResp(c, http.StatusOK, sta, err.Error())
		return
	}

	service.MakeResp(c, http.StatusOK, sta, resp)
}

func showProductList(c *gin.Context, productListReq *model.ProductListRequest) (*model.ProductListResponse, *status.Status, error) { // todo: 瀑布流怎么展现
	// 获取 mysql client
	mysqlCli := mysql.Client(c)

	products := make([]entity.Product, 0)

	// 如果 keyword 不为空，则需要查询
	if productListReq.Keyword != "" { // 这里需要除去空格和换行等 todo: 因为有很多商品，其实我们一般不用 mysql 进行查询，一般用 es 进行查询
		if err := mysqlCli.Model(&entity.Product{}).Where("title like ? or info like ?", "%"+productListReq.Keyword+"%", "%"+productListReq.Keyword+"%").Offset((productListReq.PageNum - 1) * productListReq.PageSize).Limit(productListReq.PageSize).Find(products).Error; err != nil {
			return nil, status.Error, err
		}
	} else { // 如果 keyword 为空，则直接查全量
		if err := mysqlCli.Model(&entity.Product{}).Offset((productListReq.PageNum - 1) * productListReq.PageSize).Limit(productListReq.PageSize).Find(&products).Error; err != nil {
			return nil, status.Error, err
		}
	}

	respProducts := make([]model.Product, 0)
	for _, v := range products {
		originPrice, err := strconv.ParseFloat(v.Price, 64)
		if err != nil {
			return nil, status.Error, err
		}
		discountPrice, err := strconv.ParseFloat(v.DiscountPrice, 64)
		if err != nil {
			return nil, status.Error, err
		}
		respProducts = append(respProducts, model.Product{
			Id:                 v.ID,
			CategoryId:         v.CategoryId,
			Title:              v.Title,
			Info:               v.Info,
			CoverImgPath:       v.ImgPath,
			OriginPrice:        originPrice,   // todo: 后续金额都采用 float 类型
			DiscountPrice:      discountPrice, // todo: 后续金额都采用 float 类型
			Locate:             v.Locate,
			Freight:            model.FreightType(v.Freight),
			BossId:             v.BossID,
			BossName:           v.BossName,
			BossAvatar:         v.BossAvatar,
			ExposureTimes:      v.ExposureTimes,
			ClickTimes:         v.ClickTimes,
			CommunicationTimes: v.CommunicationTimes,
		})
	}
	resp := &model.ProductListResponse{
		Products: respProducts,
		Total:    len(products),
	}
	return resp, status.Success, nil
}
