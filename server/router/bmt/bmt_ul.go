package bmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BmtUlRouter struct {
}

// InitBmtUlRouter 初始化 BmtUl 路由信息
func (s *BmtUlRouter) InitBmtUlRouter(Router *gin.RouterGroup) {
	bmtUlRouter := Router.Group("bmtUl").Use(middleware.OperationRecord())
	bmtUlRouterWithoutRecord := Router.Group("bmtUl")
	var bmtUlApi = v1.ApiGroupApp.BmtApiGroup.BmtUlApi
	{
		bmtUlRouter.POST("createBmtUl", bmtUlApi.CreateBmtUl)             // 新建BmtUl
		bmtUlRouter.DELETE("deleteBmtUl", bmtUlApi.DeleteBmtUl)           // 删除BmtUl
		bmtUlRouter.DELETE("deleteBmtUlByIds", bmtUlApi.DeleteBmtUlByIds) // 批量删除BmtUl
		bmtUlRouter.PUT("updateBmtUl", bmtUlApi.UpdateBmtUl)              // 更新BmtUl
	}
	{
		bmtUlRouterWithoutRecord.GET("findBmtUl", bmtUlApi.FindBmtUl)       // 根据ID获取BmtUl
		bmtUlRouterWithoutRecord.GET("getBmtUlList", bmtUlApi.GetBmtUlList) // 获取BmtUl列表
	}
}
