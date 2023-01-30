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
)

func AddFavor(c *gin.Context) {
	claim, err := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		service.MakeResp(c, http.StatusOK, status.AccessErr, err.Error())
		return
	}

	// 绑定参数
	req := new(model.AddFavorRequest)
	if err := c.ShouldBind(req); err != nil {
		service.MakeResp(c, http.StatusOK, status.AccessErr, err.Error())
		return
	}

	// 添加到收藏夹
	resp, sta, err := addFavorService(c, req, claim.ID)
	if err != nil {
		service.MakeResp(c, http.StatusOK, sta, err.Error())
		return
	}

	service.MakeResp(c, http.StatusOK, sta, resp)
}

func addFavorService(c *gin.Context, req *model.AddFavorRequest, userId uint) (*model.AddFavorResponse, *status.Status, error) {
	mysqlCli := mysql.Client(c)

	// 获取 product
	product := new(entity.Product)
	if err := mysqlCli.Model(&entity.Favorite{}).Where("product_id = ?", req.ProductId).Take(product).Error; err != nil {
		return nil, status.Error, err
	}

	// 存储到收藏夹表
	favor := &entity.Favorite{
		UserID:    userId,
		ProductID: product.ID,
		BossID:    product.BossID,
	}
	if err := mysqlCli.Model(&entity.Favorite{}).Create(favor).Error; err != nil {
		return nil, status.Error, err
	}

	return &model.AddFavorResponse{
		ProductId:  product.ID,
		CategoryId: product.CategoryId,
		BossId:     product.BossID,
	}, status.Success, nil
}
