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

type BmtUlApi struct {
}

var bmtUlService = service.ServiceGroupApp.BmtServiceGroup.BmtUlService

// CreateBmtUl 创建BmtUl
// @Tags BmtUl
// @Summary 创建BmtUl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bmt.BmtUl true "创建BmtUl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bmtUl/createBmtUl [post]
func (bmtUlApi *BmtUlApi) CreateBmtUl(c *gin.Context) {
	var bmtUl bmt.BmtUl
	_ = c.ShouldBindJSON(&bmtUl)
	if err := bmtUlService.CreateBmtUl(bmtUl); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBmtUl 删除BmtUl
// @Tags BmtUl
// @Summary 删除BmtUl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bmt.BmtUl true "删除BmtUl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bmtUl/deleteBmtUl [delete]
func (bmtUlApi *BmtUlApi) DeleteBmtUl(c *gin.Context) {
	var bmtUl bmt.BmtUl
	_ = c.ShouldBindJSON(&bmtUl)
	if err := bmtUlService.DeleteBmtUl(bmtUl); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBmtUlByIds 批量删除BmtUl
// @Tags BmtUl
// @Summary 批量删除BmtUl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BmtUl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bmtUl/deleteBmtUlByIds [delete]
func (bmtUlApi *BmtUlApi) DeleteBmtUlByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bmtUlService.DeleteBmtUlByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBmtUl 更新BmtUl
// @Tags BmtUl
// @Summary 更新BmtUl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bmt.BmtUl true "更新BmtUl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bmtUl/updateBmtUl [put]
func (bmtUlApi *BmtUlApi) UpdateBmtUl(c *gin.Context) {
	var bmtUl bmt.BmtUl
	_ = c.ShouldBindJSON(&bmtUl)
	if err := bmtUlService.UpdateBmtUl(bmtUl); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBmtUl 用id查询BmtUl
// @Tags BmtUl
// @Summary 用id查询BmtUl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bmt.BmtUl true "用id查询BmtUl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bmtUl/findBmtUl [get]
func (bmtUlApi *BmtUlApi) FindBmtUl(c *gin.Context) {
	var bmtUl bmt.BmtUl
	_ = c.ShouldBindQuery(&bmtUl)
	if err, rebmtUl := bmtUlService.GetBmtUl(bmtUl.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebmtUl": rebmtUl}, c)
	}
}

// GetBmtUlList 分页获取BmtUl列表
// @Tags BmtUl
// @Summary 分页获取BmtUl列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bmtReq.BmtUlSearch true "分页获取BmtUl列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bmtUl/getBmtUlList [get]
func (bmtUlApi *BmtUlApi) GetBmtUlList(c *gin.Context) {
	var pageInfo bmtReq.BmtUlSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := bmtUlService.GetBmtUlInfoList(pageInfo); err != nil {
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
