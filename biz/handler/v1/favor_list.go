package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall/biz/dal/entity"
	"go-mall/biz/model"
	"go-mall/biz/service"
	"go-mall/client/mysql"
	"go-mall/pkg/status"
	"go-mall/pkg/util"
	"net/http"
	"strconv"
	"strings"
)

func FavorList(c *gin.Context) {
	claim, err := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		service.MakeResp(c, http.StatusOK, status.AccessErr, err.Error())
		return
	}

	// 绑定入参
	req := new(model.FavorListRequest)
	if err := c.ShouldBind(req); err != nil {
		service.MakeResp(c, http.StatusOK, status.InvalidParam, err.Error())
		return
	}
	if req.PageSize == 0 {
		req.PageSize = 5
	}
	if req.PageNum == 0 {
		req.PageNum = 1
	}

	// 展示收藏夹内的商品数据
	resp, sta, err := showFavorList(c, req, claim.ID)

	service.MakeResp(c, http.StatusOK, sta, resp)
}

func showFavorList(c *gin.Context, req *model.FavorListRequest, userId uint) (*model.FavorListResponse, *status.Status, error) {
	// 一般数据量不大的情况下多表连接查询和单表多次查询效率差不多，
	// 如果数据量足够大，一定是单表多次查询效率更高，所以很多大公司都会禁用多表连接查询，
	// 一旦数据量足够大时，多表连接查询效率就会很慢，且不利于分库分表的查询优化
	mysqlCli := mysql.Client(c)

	// 表1：先查询收藏夹表中所有收藏的商品
	favorSli := make([]entity.Favorite, 0)
	if err := mysqlCli.Model(&entity.Favorite{}).Where("user_id = ?", userId).Order("created_at desc").Find(&favorSli).Error; err != nil {
		return nil, status.Error, err
	}

	// 构造 resp 结构
	productIdSli := make([]uint, 0)
	for _, v := range favorSli {
		productIdSli = append(productIdSli, v.ID)
	}
	productSli := make([]entity.Product, 0)
	// 表2：查询每个商品的商品表数据(keyword)
	if strings.TrimSpace(req.Keyword) == "" { // todo: 这里关键词查询可以做的更智能一些
		if err := mysqlCli.Model(&entity.Product{}).Where("id in ?", productIdSli).Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Find(&productSli).Error; err != nil { // todo: 时间顺序？
			return nil, status.Error, err
		}
	} else {
		if err := mysqlCli.Model(&entity.Product{}).Where("id in ? and (title = ? or info = ?)", productIdSli, "%"+req.Keyword+"%", "%"+req.Keyword+"%").Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Find(&productSli).Error; err != nil {
			return nil, status.Error, err
		}
	}
	// 表3：查询 product img
	productIdSli = make([]uint, 0)
	for _, v := range productSli {
		productIdSli = append(productIdSli, v.ID)
	}
	productImgSli := make([][]entity.ProductImg, 0)
	if err := mysqlCli.Model(&entity.ProductImg{}).Where("product_id in ?", productIdSli).Find(&productImgSli).Error; err != nil {
		return nil, status.Error, err
	}
	productImgMap := make(map[uint][]string)
	for _, v := range productImgSli {
		str := make([]string, 0)
		for _, vv := range v {
			str = append(str, vv.ImgPath)
		}
		productImgMap[v[0].ProductID] = str
	}

	// 构造结果
	respProductSli := make([]model.Product, 0)
	for _, product := range productSli {
		originPrice, err := strconv.ParseFloat(product.Price, 64)
		if err != nil {
			return nil, status.Error, err
		}
		discountPrice, err := strconv.ParseFloat(product.DiscountPrice, 64)
		if err != nil {
			return nil, status.Error, err
		}
		respProductSli = append(respProductSli, model.Product{
			Id:                 product.ID,
			CategoryId:         product.CategoryId,
			Title:              product.Title,
			Info:               product.Info,
			CoverImgPath:       product.ImgPath,
			AllImgPath:         productImgMap[product.ID],
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
		})
	}

	return &model.FavorListResponse{
		Products: respProductSli,
		Total:    len(respProductSli),
	}, status.Success, nil
}
