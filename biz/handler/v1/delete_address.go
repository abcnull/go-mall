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

func DeleteAddress(c *gin.Context) {
	claim, err := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		service.MakeResp(c, http.StatusOK, status.AccessErr, err.Error())
		return
	}

	// 绑定入参
	req := new(model.DeleteAddressRequest)
	if err := c.ShouldBind(req); err != nil {
		service.MakeResp(c, http.StatusOK, status.InvalidParam, err.Error())
		return
	}

	// 删除地址
	resp, sta, err := deleteAddressService(c, req, claim.ID)
	if err != nil {
		service.MakeResp(c, http.StatusOK, sta, err.Error())
		return
	}

	service.MakeResp(c, http.StatusOK, status.Success, resp)
}

func deleteAddressService(c *gin.Context, req *model.DeleteAddressRequest, userId uint) (*model.DeleteAddressResponse, *status.Status, error) {
	mysqlCli := mysql.Client(c)

	addrIds := make([]uint, 0)
	for _, v := range req.Address {
		addrIds = append(addrIds, v.Id)
	}

	// 删除地址，可能存在批量删除功能
	if err := mysqlCli.Model(&entity.Address{}).Where("id in ? and user_id = ?", addrIds, userId).Delete(&entity.Address{}).Error; err != nil {
		return nil, status.Error, err
	}

	return &model.DeleteAddressResponse{}, status.Success, nil
}
