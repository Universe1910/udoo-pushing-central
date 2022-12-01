import service from '@/utils/request'

// @Tags EmailTemplate
// @Summary 创建EmailTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EmailTemplate true "创建EmailTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /email_template/createEmailTemplate [post]
export const createEmailTemplate = (data) => {
  return service({
    url: '/email_template/createEmailTemplate',
    method: 'post',
    data
  })
}

// @Tags EmailTemplate
// @Summary 删除EmailTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EmailTemplate true "删除EmailTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /email_template/deleteEmailTemplate [delete]
export const deleteEmailTemplate = (data) => {
  return service({
    url: '/email_template/deleteEmailTemplate',
    method: 'delete',
    data
  })
}

// @Tags EmailTemplate
// @Summary 删除EmailTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EmailTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /email_template/deleteEmailTemplate [delete]
export const deleteEmailTemplateByIds = (data) => {
  return service({
    url: '/email_template/deleteEmailTemplateByIds',
    method: 'delete',
    data
  })
}

// @Tags EmailTemplate
// @Summary 更新EmailTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EmailTemplate true "更新EmailTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /email_template/updateEmailTemplate [put]
export const updateEmailTemplate = (data) => {
  return service({
    url: '/email_template/updateEmailTemplate',
    method: 'put',
    data
  })
}

// @Tags EmailTemplate
// @Summary 用idSearchEmailTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EmailTemplate true "用idSearchEmailTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Search成功"}"
// @Router /email_template/findEmailTemplate [get]
export const findEmailTemplate = (params) => {
  return service({
    url: '/email_template/findEmailTemplate',
    method: 'get',
    params
  })
}

// @Tags EmailTemplate
// @Summary 分页获取EmailTemplate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EmailTemplate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /email_template/getEmailTemplateList [get]
export const getEmailTemplateList = (params) => {
  return service({
    url: '/email_template/getEmailTemplateList',
    method: 'get',
    params
  })
}

// @Tags System
// @Summary 发送邮件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body email_response.Email true "发送邮件必须的参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /email_template/sendEmailTemplate [post]
export const sendEmail = (data) => {
  return service({
    url: '/email_template/sendEmailTemplate',
    method: 'post',
    data
  })
}
