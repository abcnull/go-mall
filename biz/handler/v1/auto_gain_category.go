package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall/biz/model"
	"go-mall/biz/service"
	"go-mall/pkg/status"
	"net/http"
)

// AutoGainCategory 能够即时获取商品种类，入参是 title，info，img_path，查询种类表，返回种类 id
func AutoGainCategory(c *gin.Context) {
	// 绑定入参
	req := new(model.AutoGainCategoryRequest)
	if err := c.ShouldBind(req); err != nil {
		service.MakeResp(c, http.StatusOK, status.InvalidParam, err.Error())
		return
	}

	// 获取种类 id
	cateId, cateName := gainCategoryCallRPC(req)

	service.MakeResp(c, http.StatusOK, status.Success,
		model.AutoGainCategoryResponse{
			Id:   cateId,
			Name: cateName,
		},
	)
}

func gainCategoryCallRPC(req *model.AutoGainCategoryRequest) (uint, string) {
	// todo: 这里调用下游 rpc 服务，通过实时入参信息判断种类 id 是什么并且返回过来（底层有查数据库的操作）
	return 1, "衣服"
}
