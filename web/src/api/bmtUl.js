import service from '@/utils/request'

// @Tags BmtUl
// @Summary 创建BmtUl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BmtUl true "创建BmtUl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bmtUl/createBmtUl [post]
export const createBmtUl = (data) => {
  return service({
    url: '/bmtUl/createBmtUl',
    method: 'post',
    data
  })
}

// @Tags BmtUl
// @Summary 删除BmtUl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BmtUl true "删除BmtUl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bmtUl/deleteBmtUl [delete]
export const deleteBmtUl = (data) => {
  return service({
    url: '/bmtUl/deleteBmtUl',
    method: 'delete',
    data
  })
}

// @Tags BmtUl
// @Summary 删除BmtUl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BmtUl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bmtUl/deleteBmtUl [delete]
export const deleteBmtUlByIds = (data) => {
  return service({
    url: '/bmtUl/deleteBmtUlByIds',
    method: 'delete',
    data
  })
}

// @Tags BmtUl
// @Summary 更新BmtUl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BmtUl true "更新BmtUl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bmtUl/updateBmtUl [put]
export const updateBmtUl = (data) => {
  return service({
    url: '/bmtUl/updateBmtUl',
    method: 'put',
    data
  })
}

// @Tags BmtUl
// @Summary 用id查询BmtUl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BmtUl true "用id查询BmtUl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bmtUl/findBmtUl [get]
export const findBmtUl = (params) => {
  return service({
    url: '/bmtUl/findBmtUl',
    method: 'get',
    params
  })
}

// @Tags BmtUl
// @Summary 分页获取BmtUl列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BmtUl列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bmtUl/getBmtUlList [get]
export const getBmtUlList = (params) => {
  return service({
    url: '/bmtUl/getBmtUlList',
    method: 'get',
    params
  })
}
