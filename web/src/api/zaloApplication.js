import service from '@/utils/request'

// @Tags ZaloApplication
// @Summary Create ZaloApplication
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZaloApplication true "Create ZaloApplication"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /zaloApplication/createZaloApplication [post]
export const createZaloApplication = (data) => {
  return service({
    url: '/zaloApplication/createZaloApplication',
    method: 'post',
    data
  })
}

// @Tags ZaloApplication
// @Summary Delete ZaloApplication
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZaloApplication true "Delete ZaloApplication"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /zaloApplication/deleteZaloApplication [delete]
export const deleteZaloApplication = (data) => {
  return service({
    url: '/zaloApplication/deleteZaloApplication',
    method: 'delete',
    data
  })
}

// @Tags ZaloApplication
// @Summary Delete ZaloApplication
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量Delete ZaloApplication"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /zaloApplication/deleteZaloApplication [delete]
export const deleteZaloApplicationByIds = (data) => {
  return service({
    url: '/zaloApplication/deleteZaloApplicationByIds',
    method: 'delete',
    data
  })
}

// @Tags ZaloApplication
// @Summary 更新ZaloApplication
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZaloApplication true "更新ZaloApplication"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /zaloApplication/updateZaloApplication [put]
export const updateZaloApplication = (data) => {
  return service({
    url: '/zaloApplication/updateZaloApplication',
    method: 'put',
    data
  })
}

// @Tags ZaloApplication
// @Summary 用idSearchZaloApplication
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ZaloApplication true "用idSearchZaloApplication"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /zaloApplication/findZaloApplication [get]
export const findZaloApplication = (params) => {
  return service({
    url: '/zaloApplication/findZaloApplication',
    method: 'get',
    params
  })
}

// @Tags ZaloApplication
// @Summary 分页获取ZaloApplication列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ZaloApplication列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /zaloApplication/getZaloApplicationList [get]
export const getZaloApplicationList = (params) => {
  return service({
    url: '/zaloApplication/getZaloApplicationList',
    method: 'get',
    params
  })
}
