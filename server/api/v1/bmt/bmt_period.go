package bmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bmt"
	bmtReq "github.com/flipped-aurora/gin-vue-admin/server/model/bmt/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BmtPeriodApi struct {
}

var bmtPeriodService = service.ServiceGroupApp.BmtServiceGroup.BmtPeriodService

// CreateBmtPeriod 创建BmtPeriod
// @Tags BmtPeriod
// @Summary 创建BmtPeriod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bmt.BmtPeriod true "创建BmtPeriod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bmtPeriod/createBmtPeriod [post]
func (bmtPeriodApi *BmtPeriodApi) CreateBmtPeriod(c *gin.Context) {
	var bmtPeriod bmt.BmtPeriod
	_ = c.ShouldBindJSON(&bmtPeriod)
	if err := bmtPeriodService.CreateBmtPeriod(bmtPeriod); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBmtPeriod 删除BmtPeriod
// @Tags BmtPeriod
// @Summary 删除BmtPeriod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bmt.BmtPeriod true "删除BmtPeriod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bmtPeriod/deleteBmtPeriod [delete]
func (bmtPeriodApi *BmtPeriodApi) DeleteBmtPeriod(c *gin.Context) {
	var bmtPeriod bmt.BmtPeriod
	_ = c.ShouldBindJSON(&bmtPeriod)
	if err := bmtPeriodService.DeleteBmtPeriod(bmtPeriod); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBmtPeriodByIds 批量删除BmtPeriod
// @Tags BmtPeriod
// @Summary 批量删除BmtPeriod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BmtPeriod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bmtPeriod/deleteBmtPeriodByIds [delete]
func (bmtPeriodApi *BmtPeriodApi) DeleteBmtPeriodByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bmtPeriodService.DeleteBmtPeriodByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBmtPeriod 更新BmtPeriod
// @Tags BmtPeriod
// @Summary 更新BmtPeriod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bmt.BmtPeriod true "更新BmtPeriod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bmtPeriod/updateBmtPeriod [put]
func (bmtPeriodApi *BmtPeriodApi) UpdateBmtPeriod(c *gin.Context) {
	var bmtPeriod bmt.BmtPeriod
	_ = c.ShouldBindJSON(&bmtPeriod)
	if err := bmtPeriodService.UpdateBmtPeriod(bmtPeriod); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBmtPeriod 用id查询BmtPeriod
// @Tags BmtPeriod
// @Summary 用id查询BmtPeriod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bmt.BmtPeriod true "用id查询BmtPeriod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bmtPeriod/findBmtPeriod [get]
func (bmtPeriodApi *BmtPeriodApi) FindBmtPeriod(c *gin.Context) {
	var bmtPeriod bmt.BmtPeriod
	_ = c.ShouldBindQuery(&bmtPeriod)
	if err, rebmtPeriod := bmtPeriodService.GetBmtPeriod(bmtPeriod.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebmtPeriod": rebmtPeriod}, c)
	}
}

// GetBmtPeriodList 分页获取BmtPeriod列表
// @Tags BmtPeriod
// @Summary 分页获取BmtPeriod列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bmtReq.BmtPeriodSearch true "分页获取BmtPeriod列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bmtPeriod/getBmtPeriodList [get]
func (bmtPeriodApi *BmtPeriodApi) GetBmtPeriodList(c *gin.Context) {
	var pageInfo bmtReq.BmtPeriodSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := bmtPeriodService.GetBmtPeriodInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
