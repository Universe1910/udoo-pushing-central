import service from '@/utils/request'

// @Tags CampaignLog
// @Summary Create CampaignLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CampaignLog true "Create CampaignLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /campaignLog/createCampaignLog [post]
export const createCampaignLog = (data) => {
  return service({
    url: '/campaignLog/createCampaignLog',
    method: 'post',
    data
  })
}

// @Tags CampaignLog
// @Summary Delete CampaignLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CampaignLog true "Delete CampaignLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /campaignLog/deleteCampaignLog [delete]
export const deleteCampaignLog = (data) => {
  return service({
    url: '/campaignLog/deleteCampaignLog',
    method: 'delete',
    data
  })
}

// @Tags CampaignLog
// @Summary Delete CampaignLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量Delete CampaignLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /campaignLog/deleteCampaignLog [delete]
export const deleteCampaignLogByIds = (data) => {
  return service({
    url: '/campaignLog/deleteCampaignLogByIds',
    method: 'delete',
    data
  })
}

// @Tags CampaignLog
// @Summary 更新CampaignLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CampaignLog true "更新CampaignLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /campaignLog/updateCampaignLog [put]
export const updateCampaignLog = (data) => {
  return service({
    url: '/campaignLog/updateCampaignLog',
    method: 'put',
    data
  })
}

// @Tags CampaignLog
// @Summary 用idSearchCampaignLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CampaignLog true "用idSearchCampaignLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /campaignLog/findCampaignLog [get]
export const findCampaignLog = (params) => {
  return service({
    url: '/campaignLog/findCampaignLog',
    method: 'get',
    params
  })
}

// @Tags CampaignLog
// @Summary 用idSearchCampaignLogByCampaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CampaignLog true "用idSearchCampaignLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /campaignLog/findCampaignLogByCampaign [get]
export const findCampaignLogByCampaign = (params) => {
  return service({
    url: '/campaignLog/findCampaignLogByCampaign',
    method: 'get',
    params
  })
}

// @Tags CampaignLog
// @Summary 分页获取CampaignLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CampaignLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /campaignLog/getCampaignLogList [get]
export const getCampaignLogList = (params) => {
  return service({
    url: '/campaignLog/getCampaignLogList',
    method: 'get',
    params
  })
}
