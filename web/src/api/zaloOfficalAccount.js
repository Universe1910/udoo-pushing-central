import service from '@/utils/request'

// @Tags ZaloOfficalAccount
// @Summary Create ZaloOfficalAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZaloOfficalAccount true "Create ZaloOfficalAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /zaloOfficalAccount/createZaloOfficalAccount [post]
export const createZaloOfficalAccount = (data) => {
  return service({
    url: '/zaloOfficalAccount/createZaloOfficalAccount',
    method: 'post',
    data
  })
}

// @Tags ZaloOfficalAccount
// @Summary Delete ZaloOfficalAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZaloOfficalAccount true "Delete ZaloOfficalAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /zaloOfficalAccount/deleteZaloOfficalAccount [delete]
export const deleteZaloOfficalAccount = (data) => {
  return service({
    url: '/zaloOfficalAccount/deleteZaloOfficalAccount',
    method: 'delete',
    data
  })
}

// @Tags ZaloOfficalAccount
// @Summary Delete ZaloOfficalAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量Delete ZaloOfficalAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /zaloOfficalAccount/deleteZaloOfficalAccount [delete]
export const deleteZaloOfficalAccountByIds = (data) => {
  return service({
    url: '/zaloOfficalAccount/deleteZaloOfficalAccountByIds',
    method: 'delete',
    data
  })
}

// @Tags ZaloOfficalAccount
// @Summary 更新ZaloOfficalAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZaloOfficalAccount true "更新ZaloOfficalAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /zaloOfficalAccount/updateZaloOfficalAccount [put]
export const updateZaloOfficalAccount = (data) => {
  return service({
    url: '/zaloOfficalAccount/updateZaloOfficalAccount',
    method: 'put',
    data
  })
}

// @Tags ZaloOfficalAccount
// @Summary 用idSearchZaloOfficalAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ZaloOfficalAccount true "用idSearchZaloOfficalAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /zaloOfficalAccount/findZaloOfficalAccount [get]
export const findZaloOfficalAccount = (params) => {
  return service({
    url: '/zaloOfficalAccount/findZaloOfficalAccount',
    method: 'get',
    params
  })
}

// @Tags ZaloOfficalAccount
// @Summary 分页获取ZaloOfficalAccount列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ZaloOfficalAccount列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /zaloOfficalAccount/getZaloOfficalAccountList [get]
export const getZaloOfficalAccountList = (params) => {
  return service({
    url: '/zaloOfficalAccount/getZaloOfficalAccountList',
    method: 'get',
    params
  })
}
