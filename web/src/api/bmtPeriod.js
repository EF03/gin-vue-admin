import service from '@/utils/request'

// @Tags BmtPeriod
// @Summary 创建BmtPeriod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BmtPeriod true "创建BmtPeriod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bmtPeriod/createBmtPeriod [post]
export const createBmtPeriod = (data) => {
  return service({
    url: '/bmtPeriod/createBmtPeriod',
    method: 'post',
    data
  })
}

export const drawBmtPeriod = (data) => {
  return service({
    url: '/bmtPeriod/drawBmtPeriod',
    method: 'post',
    data
  })
}

// @Tags BmtPeriod
// @Summary 删除BmtPeriod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BmtPeriod true "删除BmtPeriod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bmtPeriod/deleteBmtPeriod [delete]
export const deleteBmtPeriod = (data) => {
  return service({
    url: '/bmtPeriod/deleteBmtPeriod',
    method: 'delete',
    data
  })
}

// @Tags BmtPeriod
// @Summary 删除BmtPeriod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BmtPeriod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bmtPeriod/deleteBmtPeriod [delete]
export const deleteBmtPeriodByIds = (data) => {
  return service({
    url: '/bmtPeriod/deleteBmtPeriodByIds',
    method: 'delete',
    data
  })
}

// @Tags BmtPeriod
// @Summary 更新BmtPeriod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BmtPeriod true "更新BmtPeriod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bmtPeriod/updateBmtPeriod [put]
export const updateBmtPeriod = (data) => {
  return service({
    url: '/bmtPeriod/updateBmtPeriod',
    method: 'put',
    data
  })
}

// @Tags BmtPeriod
// @Summary 用id查询BmtPeriod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BmtPeriod true "用id查询BmtPeriod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bmtPeriod/findBmtPeriod [get]
export const findBmtPeriod = (params) => {
  return service({
    url: '/bmtPeriod/findBmtPeriod',
    method: 'get',
    params
  })
}

// @Tags BmtPeriod
// @Summary 分页获取BmtPeriod列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BmtPeriod列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bmtPeriod/getBmtPeriodList [get]
export const getBmtPeriodList = (params) => {
  return service({
    url: '/bmtPeriod/getBmtPeriodList',
    method: 'get',
    params
  })
}
