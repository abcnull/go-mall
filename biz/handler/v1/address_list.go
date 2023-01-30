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

func AddressList(c *gin.Context) {
	claim, err := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		service.MakeResp(c, http.StatusOK, status.AccessErr, err.Error())
		return
	}

	// 绑定参数
	req := new(model.AddressListRequest)
	if err := c.ShouldBind(req); err != nil {
		service.MakeResp(c, http.StatusOK, status.InvalidParam, err.Error())
		return
	}

	// 展现地址列表
	resp, sta, err := addressListService(c, req, claim.ID)
	if err != nil {
		service.MakeResp(c, http.StatusOK, status.Error, err.Error())
		return
	}

	// 返回结构
	service.MakeResp(c, http.StatusOK, sta, resp)
}

func addressListService(c *gin.Context, req *model.AddressListRequest, userId uint) (*model.AddressListResponse, *status.Status, error) {
	mysqlCli := mysql.Client(c)

	// 查询所有的地址
	addrs := make([]entity.Address, 0)
	count := new(int64)
	if err := mysqlCli.Model(&entity.Address{}).Where("user_id = ?", userId).Count(count).Order("create_at desc").Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Find(&addrs).Error; err != nil {
		return nil, status.Error, err
	}

	// 封装结构体
	respAddrs := make([]model.Address, 0)
	for _, v := range addrs {
		respAddrs = append(respAddrs, model.Address{
			Name:    v.Name,
			Phone:   v.Phone,
			Address: v.Address,
		})
	}
	return &model.AddressListResponse{
		Addresses: respAddrs,
		Total:     *count,
	}, status.Success, nil
}
