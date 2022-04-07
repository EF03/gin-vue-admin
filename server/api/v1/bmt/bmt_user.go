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

type BmtUserApi struct {
}

var bmtUserService = service.ServiceGroupApp.BmtServiceGroup.BmtUserService

// CreateBmtUser 创建BmtUser
// @Tags BmtUser
// @Summary 创建BmtUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bmt.BmtUser true "创建BmtUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bmtUser/createBmtUser [post]
func (bmtUserApi *BmtUserApi) CreateBmtUser(c *gin.Context) {
	var bmtUser bmt.BmtUser
	_ = c.ShouldBindJSON(&bmtUser)
	if err := bmtUserService.CreateBmtUser(bmtUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBmtUser 删除BmtUser
// @Tags BmtUser
// @Summary 删除BmtUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bmt.BmtUser true "删除BmtUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bmtUser/deleteBmtUser [delete]
func (bmtUserApi *BmtUserApi) DeleteBmtUser(c *gin.Context) {
	var bmtUser bmt.BmtUser
	_ = c.ShouldBindJSON(&bmtUser)
	if err := bmtUserService.DeleteBmtUser(bmtUser); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBmtUserByIds 批量删除BmtUser
// @Tags BmtUser
// @Summary 批量删除BmtUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BmtUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bmtUser/deleteBmtUserByIds [delete]
func (bmtUserApi *BmtUserApi) DeleteBmtUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bmtUserService.DeleteBmtUserByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBmtUser 更新BmtUser
// @Tags BmtUser
// @Summary 更新BmtUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bmt.BmtUser true "更新BmtUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bmtUser/updateBmtUser [put]
func (bmtUserApi *BmtUserApi) UpdateBmtUser(c *gin.Context) {
	var bmtUser bmt.BmtUser
	_ = c.ShouldBindJSON(&bmtUser)
	if err := bmtUserService.UpdateBmtUser(bmtUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBmtUser 用id查询BmtUser
// @Tags BmtUser
// @Summary 用id查询BmtUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bmt.BmtUser true "用id查询BmtUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bmtUser/findBmtUser [get]
func (bmtUserApi *BmtUserApi) FindBmtUser(c *gin.Context) {
	var bmtUser bmt.BmtUser
	_ = c.ShouldBindQuery(&bmtUser)
	if err, rebmtUser := bmtUserService.GetBmtUser(bmtUser.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebmtUser": rebmtUser}, c)
	}
}

// GetBmtUserList 分页获取BmtUser列表
// @Tags BmtUser
// @Summary 分页获取BmtUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bmtReq.BmtUserSearch true "分页获取BmtUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bmtUser/getBmtUserList [get]
func (bmtUserApi *BmtUserApi) GetBmtUserList(c *gin.Context) {
	var pageInfo bmtReq.BmtUserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := bmtUserService.GetBmtUserInfoList(pageInfo); err != nil {
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
