import service from '@/utils/request'

// @Tags Trigger
// @Summary Create Trigger
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Trigger true "Create Trigger"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /trigger/createTrigger [post]
export const createTrigger = (data) => {
  return service({
    url: '/trigger/createTrigger',
    method: 'post',
    data
  })
}

// @Tags Trigger
// @Summary Delete Trigger
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Trigger true "Delete Trigger"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /trigger/deleteTrigger [delete]
export const deleteTrigger = (data) => {
  return service({
    url: '/trigger/deleteTrigger',
    method: 'delete',
    data
  })
}



// @Tags Trigger
// @Summary Delete Trigger
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量Delete Trigger"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /trigger/deleteTrigger [delete]
export const deleteTriggerByIds = (data) => {
  return service({
    url: '/trigger/deleteTriggerByIds',
    method: 'delete',
    data
  })
}

// @Tags Trigger
// @Summary 更新Trigger
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Trigger true "更新Trigger"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /trigger/updateTrigger [put]
export const updateTrigger = (data) => {
  return service({
    url: '/trigger/updateTrigger',
    method: 'put',
    data
  })
}

// @Tags Trigger
// @Summary 用idSearchTrigger
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Trigger true "用idSearchTrigger"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /trigger/findTrigger [get]
export const findTrigger = (params) => {
  return service({
    url: '/trigger/findTrigger',
    method: 'get',
    params
  })
}

// @Tags Trigger
// @Summary 分页获取Trigger列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Trigger列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /trigger/getTriggerList [get]
export const getTriggerList = (params) => {
  return service({
    url: '/trigger/getTriggerList',
    method: 'get',
    params
  })
}
