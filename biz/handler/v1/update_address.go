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

func UpdateAddress(c *gin.Context) {
	claim, err := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		service.MakeResp(c, http.StatusOK, status.AccessErr, err.Error())
		return
	}

	// 绑定入参
	req := new(model.UpdateAddressRequest)
	if err := c.ShouldBind(req); err != nil {
		service.MakeResp(c, http.StatusOK, status.InvalidParam, err.Error())
		return
	}

	// 编辑地址
	resp, sta, err := updateAddressService(c, req, claim.ID)
	if err != nil {
		service.MakeResp(c, http.StatusOK, sta, err.Error())
		return
	}

	service.MakeResp(c, http.StatusOK, status.Success, resp)
}

func updateAddressService(c *gin.Context, req *model.UpdateAddressRequest, userId uint) (*model.UpdateAddressResponse, *status.Status, error) {
	mysqlCli := mysql.Client(c)

	// 更新地址
	addr := &entity.Address{
		UserID:  req.Address.Id,
		Name:    req.Address.Name,
		Phone:   req.Address.Phone,
		Address: req.Address.Address,
	}
	if err := mysqlCli.Model(&entity.Address{}).Where("id = ? and user_id = ?", req.Address.Id, userId).Updates(addr).Error; err != nil {
		return nil, status.Error, err
	}

	return &model.UpdateAddressResponse{Address: model.Address{
		Id:      addr.ID,
		Name:    addr.Name,
		Phone:   addr.Phone,
		Address: addr.Address,
	}}, status.Success, nil
}
