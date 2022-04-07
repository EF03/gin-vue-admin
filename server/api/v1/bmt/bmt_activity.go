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

type BmtActivityApi struct {
}

var bmtActivityService = service.ServiceGroupApp.BmtServiceGroup.BmtActivityService

// CreateBmtActivity 创建BmtActivity
// @Tags BmtActivity
// @Summary 创建BmtActivity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bmt.BmtActivity true "创建BmtActivity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bmtActivity/createBmtActivity [post]
func (bmtActivityApi *BmtActivityApi) CreateBmtActivity(c *gin.Context) {
	var bmtActivity bmt.BmtActivity
	_ = c.ShouldBindJSON(&bmtActivity)
	if err := bmtActivityService.CreateBmtActivity(bmtActivity); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBmtActivity 删除BmtActivity
// @Tags BmtActivity
// @Summary 删除BmtActivity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bmt.BmtActivity true "删除BmtActivity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bmtActivity/deleteBmtActivity [delete]
func (bmtActivityApi *BmtActivityApi) DeleteBmtActivity(c *gin.Context) {
	var bmtActivity bmt.BmtActivity
	_ = c.ShouldBindJSON(&bmtActivity)
	if err := bmtActivityService.DeleteBmtActivity(bmtActivity); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBmtActivityByIds 批量删除BmtActivity
// @Tags BmtActivity
// @Summary 批量删除BmtActivity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BmtActivity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bmtActivity/deleteBmtActivityByIds [delete]
func (bmtActivityApi *BmtActivityApi) DeleteBmtActivityByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bmtActivityService.DeleteBmtActivityByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBmtActivity 更新BmtActivity
// @Tags BmtActivity
// @Summary 更新BmtActivity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bmt.BmtActivity true "更新BmtActivity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bmtActivity/updateBmtActivity [put]
func (bmtActivityApi *BmtActivityApi) UpdateBmtActivity(c *gin.Context) {
	var bmtActivity bmt.BmtActivity
	_ = c.ShouldBindJSON(&bmtActivity)
	if err := bmtActivityService.UpdateBmtActivity(bmtActivity); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBmtActivity 用id查询BmtActivity
// @Tags BmtActivity
// @Summary 用id查询BmtActivity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bmt.BmtActivity true "用id查询BmtActivity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bmtActivity/findBmtActivity [get]
func (bmtActivityApi *BmtActivityApi) FindBmtActivity(c *gin.Context) {
	var bmtActivity bmt.BmtActivity
	_ = c.ShouldBindQuery(&bmtActivity)
	if err, rebmtActivity := bmtActivityService.GetBmtActivity(bmtActivity.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebmtActivity": rebmtActivity}, c)
	}
}

// GetBmtActivityList 分页获取BmtActivity列表
// @Tags BmtActivity
// @Summary 分页获取BmtActivity列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bmtReq.BmtActivitySearch true "分页获取BmtActivity列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bmtActivity/getBmtActivityList [get]
func (bmtActivityApi *BmtActivityApi) GetBmtActivityList(c *gin.Context) {
	var pageInfo bmtReq.BmtActivitySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := bmtActivityService.GetBmtActivityInfoList(pageInfo); err != nil {
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
