import service from '@/utils/request'

// @Tags Tags
// @Summary Create Tags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tags true "Create Tags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /tags/createTags [post]
export const createTags = (data) => {
  return service({
    url: '/tags/createTags',
    method: 'post',
    data
  })
}

// @Tags Tags
// @Summary Delete Tags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tags true "Delete Tags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /tags/deleteTags [delete]
export const deleteTags = (data) => {
  return service({
    url: '/tags/deleteTags',
    method: 'delete',
    data
  })
}

// @Tags Tags
// @Summary Delete Tags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量Delete Tags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /tags/deleteTags [delete]
export const deleteTagsByIds = (data) => {
  return service({
    url: '/tags/deleteTagsByIds',
    method: 'delete',
    data
  })
}

// @Tags Tags
// @Summary 更新Tags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tags true "更新Tags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /tags/updateTags [put]
export const updateTags = (data) => {
  return service({
    url: '/tags/updateTags',
    method: 'put',
    data
  })
}

// @Tags Tags
// @Summary 用idSearchTags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Tags true "用idSearchTags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /tags/findTags [get]
export const findTags = (params) => {
  return service({
    url: '/tags/findTags',
    method: 'get',
    params
  })
}

// @Tags Tags
// @Summary 分页获取Tags列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Tags列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /tags/getTagsList [get]
export const getTagsList = (params) => {
  return service({
    url: '/tags/getTagsList',
    method: 'get',
    params
  })
}
