package bmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BmtUserRouter struct {
}

// InitBmtUserRouter 初始化 BmtUser 路由信息
func (s *BmtUserRouter) InitBmtUserRouter(Router *gin.RouterGroup) {
	bmtUserRouter := Router.Group("bmtUser").Use(middleware.OperationRecord())
	bmtUserRouterWithoutRecord := Router.Group("bmtUser")
	var bmtUserApi = v1.ApiGroupApp.BmtApiGroup.BmtUserApi
	{
		bmtUserRouter.POST("createBmtUser", bmtUserApi.CreateBmtUser)             // 新建BmtUser
		bmtUserRouter.DELETE("deleteBmtUser", bmtUserApi.DeleteBmtUser)           // 删除BmtUser
		bmtUserRouter.DELETE("deleteBmtUserByIds", bmtUserApi.DeleteBmtUserByIds) // 批量删除BmtUser
		bmtUserRouter.PUT("updateBmtUser", bmtUserApi.UpdateBmtUser)              // 更新BmtUser
	}
	{
		bmtUserRouterWithoutRecord.GET("findBmtUser", bmtUserApi.FindBmtUser)       // 根据ID获取BmtUser
		bmtUserRouterWithoutRecord.GET("getBmtUserList", bmtUserApi.GetBmtUserList) // 获取BmtUser列表
	}
}
