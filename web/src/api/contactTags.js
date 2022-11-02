import service from '@/utils/request'

// @Tags ContactTags
// @Summary Create ContactTags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContactTags true "Create ContactTags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /contactTags/createContactTags [post]
export const createContactTags = (data) => {
  return service({
    url: '/contactTags/createContactTags',
    method: 'post',
    data
  })
}

// @Tags ContactTags
// @Summary Delete ContactTags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContactTags true "Delete ContactTags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /contactTags/deleteContactTags [delete]
export const deleteContactTags = (data) => {
  return service({
    url: '/contactTags/deleteContactTags',
    method: 'delete',
    data
  })
}

// @Tags ContactTags
// @Summary Delete ContactTags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量Delete ContactTags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /contactTags/deleteContactTags [delete]
export const deleteContactTagsByIds = (data) => {
  return service({
    url: '/contactTags/deleteContactTagsByIds',
    method: 'delete',
    data
  })
}

// @Tags ContactTags
// @Summary 更新ContactTags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContactTags true "更新ContactTags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /contactTags/updateContactTags [put]
export const updateContactTags = (data) => {
  return service({
    url: '/contactTags/updateContactTags',
    method: 'put',
    data
  })
}

// @Tags ContactTags
// @Summary 用id查询ContactTags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ContactTags true "用id查询ContactTags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /contactTags/findContactTags [get]
export const findContactTags = (params) => {
  return service({
    url: '/contactTags/findContactTags',
    method: 'get',
    params
  })
}

// @Tags ContactTags
// @Summary 分页获取ContactTags列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ContactTags列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /contactTags/getContactTagsList [get]
export const getContactTagsList = (params) => {
  return service({
    url: '/contactTags/getContactTagsList',
    method: 'get',
    params
  })
}
