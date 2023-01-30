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

func DeleteInFavor(c *gin.Context) {
	claim, err := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		service.MakeResp(c, http.StatusOK, status.AccessErr, err.Error())
		return
	}

	// 绑定入参
	req := new(model.DeleteInFavorRequest)
	if err := c.ShouldBind(req); err != nil {
		service.MakeResp(c, http.StatusOK, status.InvalidParam, err.Error())
		return
	}

	// 删除收藏夹中的商品
	resp, sta, err := deleteInFavorService(c, req, claim.ID)
	if err != nil {
		service.MakeResp(c, http.StatusOK, sta, err.Error())
		return
	}

	service.MakeResp(c, http.StatusOK, sta, resp)
}

func deleteInFavorService(c *gin.Context, req *model.DeleteInFavorRequest, userId uint) (*model.DeleteInFavorResponse, *status.Status, error) {
	mysqlCli := mysql.Client(c)

	// 删除收藏夹中的商品
	favor := new(entity.Favorite)
	if err := mysqlCli.Model(&entity.Favorite{}).Where("product_id = ? and user_id = ?", req.ProductId, userId).Delete(favor).Error; err != nil {
		return nil, status.Error, err
	}

	return &model.DeleteInFavorResponse{}, status.Success, nil
}
