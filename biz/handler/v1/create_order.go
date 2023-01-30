package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mall/biz/dal/entity"
	"go-mall/biz/model"
	"go-mall/biz/service"
	"go-mall/client/mysql"
	"go-mall/pkg/status"
	"go-mall/pkg/util"
	"math/rand"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	claim, err := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		service.MakeResp(c, http.StatusOK, status.AccessErr, err.Error())
		return
	}

	// 绑定入参
	req := new(model.CreateOrderRequest)
	if err := c.ShouldBind(req); err != nil {
		service.MakeResp(c, http.StatusOK, status.InvalidParam, err.Error())
		return
	}

	// 检验地址是否存在
	mysqlCli := mysql.Client(c)
	count := new(int64)
	if err := mysqlCli.Model(&entity.Address{}).Where("id = ?", req.AddressId).Find(&entity.Address{}).Count(count).Error; err != nil {
		service.MakeResp(c, http.StatusOK, status.Error, err.Error())
		return
	}
	if *count == 0 { // 如果地址不存在
		service.MakeResp(c, http.StatusOK, status.InvalidParam, errors.New("地址不存在"))
		return
	}

	// 生成订单号
	num := createOrderNumber(c)

	// 创建订单
	order := &entity.Order{
		UserID:    claim.ID,
		ProductID: req.ProductId,
		BossID:    req.BossId,
		AddressID: req.AddressId,
		Num:       1,
		OrderNum:  int(num),
		Type:      uint(req.Type),
		Money:     req.Money,
	}
	if err := mysqlCli.Model(&entity.Order{}).Create(order).Error; err != nil {
		service.MakeResp(c, http.StatusOK, status.Error, err.Error())
		return
	}

	// 返回
	service.MakeResp(c, http.StatusOK, status.Success, resp)
}

// 生成订单编号
func createOrderNumber(c *gin.Context) uint64 { // todo: 待做完
	randNum := fmt.Sprintf("%s", rand.New())
}
