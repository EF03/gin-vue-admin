package bmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BmtPeriodRouter struct {
}

// InitBmtPeriodRouter 初始化 BmtPeriod 路由信息
func (s *BmtPeriodRouter) InitBmtPeriodRouter(Router *gin.RouterGroup) {
	bmtPeriodRouter := Router.Group("bmtPeriod").Use(middleware.OperationRecord())
	bmtPeriodRouterWithoutRecord := Router.Group("bmtPeriod")
	var bmtPeriodApi = v1.ApiGroupApp.BmtApiGroup.BmtPeriodApi
	{
		bmtPeriodRouter.POST("createBmtPeriod", bmtPeriodApi.CreateBmtPeriod)             // 新建BmtPeriod
		bmtPeriodRouter.DELETE("deleteBmtPeriod", bmtPeriodApi.DeleteBmtPeriod)           // 删除BmtPeriod
		bmtPeriodRouter.DELETE("deleteBmtPeriodByIds", bmtPeriodApi.DeleteBmtPeriodByIds) // 批量删除BmtPeriod
		bmtPeriodRouter.PUT("updateBmtPeriod", bmtPeriodApi.UpdateBmtPeriod)              // 更新BmtPeriod
	}
	{
		bmtPeriodRouterWithoutRecord.GET("findBmtPeriod", bmtPeriodApi.FindBmtPeriod)       // 根据ID获取BmtPeriod
		bmtPeriodRouterWithoutRecord.GET("getBmtPeriodList", bmtPeriodApi.GetBmtPeriodList) // 获取BmtPeriod列表
	}
}
