import service from '@/utils/request'

// @Tags BmtActivity
// @Summary 创建BmtActivity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BmtActivity true "创建BmtActivity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bmtActivity/createBmtActivity [post]
export const createBmtActivity = (data) => {
  return service({
    url: '/bmtActivity/createBmtActivity',
    method: 'post',
    data
  })
}

// @Tags BmtActivity
// @Summary 删除BmtActivity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BmtActivity true "删除BmtActivity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bmtActivity/deleteBmtActivity [delete]
export const deleteBmtActivity = (data) => {
  return service({
    url: '/bmtActivity/deleteBmtActivity',
    method: 'delete',
    data
  })
}

// @Tags BmtActivity
// @Summary 删除BmtActivity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BmtActivity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bmtActivity/deleteBmtActivity [delete]
export const deleteBmtActivityByIds = (data) => {
  return service({
    url: '/bmtActivity/deleteBmtActivityByIds',
    method: 'delete',
    data
  })
}

// @Tags BmtActivity
// @Summary 更新BmtActivity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BmtActivity true "更新BmtActivity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bmtActivity/updateBmtActivity [put]
export const updateBmtActivity = (data) => {
  return service({
    url: '/bmtActivity/updateBmtActivity',
    method: 'put',
    data
  })
}

// @Tags BmtActivity
// @Summary 用id查询BmtActivity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BmtActivity true "用id查询BmtActivity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bmtActivity/findBmtActivity [get]
export const findBmtActivity = (params) => {
  return service({
    url: '/bmtActivity/findBmtActivity',
    method: 'get',
    params
  })
}

// @Tags BmtActivity
// @Summary 分页获取BmtActivity列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BmtActivity列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bmtActivity/getBmtActivityList [get]
export const getBmtActivityList = (params) => {
  return service({
    url: '/bmtActivity/getBmtActivityList',
    method: 'get',
    params
  })
}
