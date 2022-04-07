package bmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BmtActivityRouter struct {
}

// InitBmtActivityRouter 初始化 BmtActivity 路由信息
func (s *BmtActivityRouter) InitBmtActivityRouter(Router *gin.RouterGroup) {
	bmtActivityRouter := Router.Group("bmtActivity").Use(middleware.OperationRecord())
	bmtActivityRouterWithoutRecord := Router.Group("bmtActivity")
	var bmtActivityApi = v1.ApiGroupApp.BmtApiGroup.BmtActivityApi
	{
		bmtActivityRouter.POST("createBmtActivity", bmtActivityApi.CreateBmtActivity)             // 新建BmtActivity
		bmtActivityRouter.DELETE("deleteBmtActivity", bmtActivityApi.DeleteBmtActivity)           // 删除BmtActivity
		bmtActivityRouter.DELETE("deleteBmtActivityByIds", bmtActivityApi.DeleteBmtActivityByIds) // 批量删除BmtActivity
		bmtActivityRouter.PUT("updateBmtActivity", bmtActivityApi.UpdateBmtActivity)              // 更新BmtActivity
	}
	{
		bmtActivityRouterWithoutRecord.GET("findBmtActivity", bmtActivityApi.FindBmtActivity)       // 根据ID获取BmtActivity
		bmtActivityRouterWithoutRecord.GET("getBmtActivityList", bmtActivityApi.GetBmtActivityList) // 获取BmtActivity列表
	}
}
