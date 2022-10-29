import service from '@/utils/request'

// @Tags UdooEmail
// @Summary 创建UdooEmail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UdooEmail true "创建UdooEmail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /udooEmail/createUdooEmail [post]
export const createUdooEmail = (data) => {
  return service({
    url: '/udooEmail/createUdooEmail',
    method: 'post',
    data
  })
}

// @Tags UdooEmail
// @Summary 删除UdooEmail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UdooEmail true "删除UdooEmail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /udooEmail/deleteUdooEmail [delete]
export const deleteUdooEmail = (data) => {
  return service({
    url: '/udooEmail/deleteUdooEmail',
    method: 'delete',
    data
  })
}

// @Tags UdooEmail
// @Summary 删除UdooEmail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UdooEmail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /udooEmail/deleteUdooEmail [delete]
export const deleteUdooEmailByIds = (data) => {
  return service({
    url: '/udooEmail/deleteUdooEmailByIds',
    method: 'delete',
    data
  })
}

// @Tags UdooEmail
// @Summary 更新UdooEmail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UdooEmail true "更新UdooEmail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /udooEmail/updateUdooEmail [put]
export const updateUdooEmail = (data) => {
  return service({
    url: '/udooEmail/updateUdooEmail',
    method: 'put',
    data
  })
}

// @Tags UdooEmail
// @Summary 用id查询UdooEmail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UdooEmail true "用id查询UdooEmail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /udooEmail/findUdooEmail [get]
export const findUdooEmail = (params) => {
  return service({
    url: '/udooEmail/findUdooEmail',
    method: 'get',
    params
  })
}

// @Tags UdooEmail
// @Summary 分页获取UdooEmail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取UdooEmail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /udooEmail/getUdooEmailList [get]
export const getUdooEmailList = (params) => {
  return service({
    url: '/udooEmail/getUdooEmailList',
    method: 'get',
    params
  })
}
