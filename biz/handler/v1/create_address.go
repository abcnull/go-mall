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

func CreateAddress(c *gin.Context) {
	claim, err := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		service.MakeResp(c, http.StatusOK, status.AccessErr, err.Error())
		return
	}

	// 绑定参数
	req := new(model.CreateAddressRequest)
	if err := c.ShouldBind(req); err != nil {
		service.MakeResp(c, http.StatusOK, status.InvalidParam, err.Error())
		return
	}

	// 创建地址
	resp, sta, err := createAddressService(c, req, claim.ID)
	if err != nil {
		service.MakeResp(c, http.StatusOK, sta, err.Error())
		return
	}

	service.MakeResp(c, http.StatusOK, status.Success, resp)
}

func createAddressService(c *gin.Context, req *model.CreateAddressRequest, userId uint) (*model.CreateAddressResponse, *status.Status, error) {
	mysqlCli := mysql.Client(c)

	// 创建地址
	addr := &entity.Address{
		UserID:  userId,
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
	}
	if err := mysqlCli.Model(&entity.Address{}).Create(addr).Error; err != nil { // todo: 这里可以加上地址是否默认当前
		return nil, status.Error, err
	}

	return &model.CreateAddressResponse{Address: model.Address{
		Id:      addr.ID,
		Name:    addr.Name,
		Phone:   addr.Name,
		Address: addr.Address,
	}}, status.Success, nil
}
