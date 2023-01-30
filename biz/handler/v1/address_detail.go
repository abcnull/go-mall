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
)

func AddressDetail(c *gin.Context) {
	claim, err := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		service.MakeResp(c, http.StatusOK, status.AccessErr, err.Error())
		return
	}

	// 绑定入参
	req := new(model.AddressDetailRequest)
	if err := c.ShouldBind(req); err != nil {
		service.MakeResp(c, http.StatusOK, status.AccessErr, err.Error())
		return
	}

	// 展示地址详细信息
	addrId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		service.MakeResp(c, http.StatusOK, status.Error, err.Error())
		return
	}
	resp, sta, err := addressDetailService(c, req, uint(addrId), claim.ID)
	service.MakeResp(c, http.StatusOK, sta, resp)
}

func addressDetailService(c *gin.Context, req *model.AddressDetailRequest, id uint, userId uint) (*model.AddressDetailResponse, *status.Status, error) {
	mysqlCli := mysql.Client(c)

	// 拿到详细地址信息
	addr := new(entity.Address)
	if err := mysqlCli.Model(&entity.Address{}).Where("id = ? and user_id = ?", id, userId).Take(addr).Error; err != nil {
		return nil, status.Error, err
	}

	return &model.AddressDetailResponse{Address: model.Address{
		Id:      addr.ID,
		Name:    addr.Name,
		Phone:   addr.Phone,
		Address: addr.Address,
	}}, status.Success, nil
}
