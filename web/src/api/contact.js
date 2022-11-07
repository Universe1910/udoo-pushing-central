import service from '@/utils/request'

// @Tags Contact
// @Summary Create Contact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Contact true "Create Contact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /contact/createContact [post]
export const createContact = (data) => {
  return service({
    url: '/contact/createContact',
    method: 'post',
    data
  })
}

// @Tags Contact
// @Summary Delete Contact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Contact true "Delete Contact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /contact/deleteContact [delete]
export const deleteContact = (data) => {
  return service({
    url: '/contact/deleteContact',
    method: 'delete',
    data
  })
}

// @Tags Contact
// @Summary Delete Contact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量Delete Contact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /contact/deleteContact [delete]
export const deleteContactByIds = (data) => {
  return service({
    url: '/contact/deleteContactByIds',
    method: 'delete',
    data
  })
}

// @Tags Contact
// @Summary 更新Contact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Contact true "更新Contact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /contact/updateContact [put]
export const updateContact = (data) => {
  return service({
    url: '/contact/updateContact',
    method: 'put',
    data
  })
}

// @Tags Contact
// @Summary 用idSearchContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Contact true "用idSearchContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /contact/findContact [get]
export const findContact = (params) => {
  return service({
    url: '/contact/findContact',
    method: 'get',
    params
  })
}

// @Tags Contact
// @Summary 分页获取Contact列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Contact列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /contact/getContactList [get]
export const getContactList = (params) => {
  return service({
    url: '/contact/getContactList',
    method: 'get',
    params
  })
}
