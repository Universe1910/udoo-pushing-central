import service from '@/utils/request'

// @Tags Facebook
// @Summary Create Facebook
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Facebook true "Create Facebook"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /facebook/createFacebook [post]
export const createFacebook = (data) => {
  return service({
    url: '/facebook/createFacebook',
    method: 'post',
    data
  })
}

// @Tags Facebook
// @Summary Delete Facebook
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Facebook true "Delete Facebook"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /facebook/deleteFacebook [delete]
export const deleteFacebook = (data) => {
  return service({
    url: '/facebook/deleteFacebook',
    method: 'delete',
    data
  })
}

// @Tags Facebook
// @Summary Delete Facebook
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量Delete Facebook"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /facebook/deleteFacebook [delete]
export const deleteFacebookByIds = (data) => {
  return service({
    url: '/facebook/deleteFacebookByIds',
    method: 'delete',
    data
  })
}

// @Tags Facebook
// @Summary 更新Facebook
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Facebook true "更新Facebook"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /facebook/updateFacebook [put]
export const updateFacebook = (data) => {
  return service({
    url: '/facebook/updateFacebook',
    method: 'put',
    data
  })
}

// @Tags Facebook
// @Summary 用idSearchFacebook
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Facebook true "用idSearchFacebook"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /facebook/findFacebook [get]
export const findFacebook = (params) => {
  return service({
    url: '/facebook/findFacebook',
    method: 'get',
    params
  })
}

// @Tags Facebook
// @Summary 分页获取Facebook列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Facebook列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /facebook/getFacebookList [get]
export const getFacebookList = (params) => {
  return service({
    url: '/facebook/getFacebookList',
    method: 'get',
    params
  })
}
