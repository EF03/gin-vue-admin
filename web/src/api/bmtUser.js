import service from '@/utils/request'

// @Tags BmtUser
// @Summary 创建BmtUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BmtUser true "创建BmtUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bmtUser/createBmtUser [post]
export const createBmtUser = (data) => {
  return service({
    url: '/bmtUser/createBmtUser',
    method: 'post',
    data
  })
}

// @Tags BmtUser
// @Summary 删除BmtUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BmtUser true "删除BmtUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bmtUser/deleteBmtUser [delete]
export const deleteBmtUser = (data) => {
  return service({
    url: '/bmtUser/deleteBmtUser',
    method: 'delete',
    data
  })
}

// @Tags BmtUser
// @Summary 删除BmtUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BmtUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bmtUser/deleteBmtUser [delete]
export const deleteBmtUserByIds = (data) => {
  return service({
    url: '/bmtUser/deleteBmtUserByIds',
    method: 'delete',
    data
  })
}

// @Tags BmtUser
// @Summary 更新BmtUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BmtUser true "更新BmtUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bmtUser/updateBmtUser [put]
export const updateBmtUser = (data) => {
  return service({
    url: '/bmtUser/updateBmtUser',
    method: 'put',
    data
  })
}

// @Tags BmtUser
// @Summary 用id查询BmtUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BmtUser true "用id查询BmtUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bmtUser/findBmtUser [get]
export const findBmtUser = (params) => {
  return service({
    url: '/bmtUser/findBmtUser',
    method: 'get',
    params
  })
}

// @Tags BmtUser
// @Summary 分页获取BmtUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BmtUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bmtUser/getBmtUserList [get]
export const getBmtUserList = (params) => {
  return service({
    url: '/bmtUser/getBmtUserList',
    method: 'get',
    params
  })
}
