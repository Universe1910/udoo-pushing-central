import service from '@/utils/request'

// @Tags ContactCampaign
// @Summary Create ContactCampaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContactCampaign true "Create ContactCampaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /contactCampaign/createContactCampaign [post]
export const createContactCampaign = (data) => {
  return service({
    url: '/contactCampaign/createContactCampaign',
    method: 'post',
    data
  })
}

// @Tags ContactCampaign
// @Summary Delete ContactCampaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContactCampaign true "Delete ContactCampaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /contactCampaign/deleteContactCampaign [delete]
export const deleteContactCampaign = (data) => {
  return service({
    url: '/contactCampaign/deleteContactCampaign',
    method: 'delete',
    data
  })
}

// @Tags ContactCampaign
// @Summary Delete ContactCampaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量Delete ContactCampaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /contactCampaign/deleteContactCampaign [delete]
export const deleteContactCampaignByIds = (data) => {
  return service({
    url: '/contactCampaign/deleteContactCampaignByIds',
    method: 'delete',
    data
  })
}

// @Tags ContactCampaign
// @Summary 更新ContactCampaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContactCampaign true "更新ContactCampaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /contactCampaign/updateContactCampaign [put]
export const updateContactCampaign = (data) => {
  return service({
    url: '/contactCampaign/updateContactCampaign',
    method: 'put',
    data
  })
}

// @Tags ContactCampaign
// @Summary 用idSearchContactCampaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ContactCampaign true "用idSearchContactCampaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /contactCampaign/findContactCampaign [get]
export const findContactCampaign = (params) => {
  return service({
    url: '/contactCampaign/findContactCampaign',
    method: 'get',
    params
  })
}

// @Tags ContactCampaign
// @Summary 分页获取ContactCampaign列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ContactCampaign列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /contactCampaign/getContactCampaignList [get]
export const getContactCampaignList = (params) => {
  return service({
    url: '/contactCampaign/getContactCampaignList',
    method: 'get',
    params
  })
}
