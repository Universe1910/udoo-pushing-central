import service from '@/utils/request'

// @Tags Sequence
// @Summary Create Sequence
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sequence true "Create Sequence"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /sequence/createSequence [post]
export const createSequence = (data) => {
  return service({
    url: '/sequence/createSequence',
    method: 'post',
    data
  })
}

// @Tags Sequence
// @Summary Delete Sequence
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sequence true "Delete Sequence"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /sequence/deleteSequence [delete]
export const deleteSequence = (data) => {
  return service({
    url: '/sequence/deleteSequence',
    method: 'delete',
    data
  })
}

// @Tags Sequence
// @Summary Delete Sequence
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量Delete Sequence"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /sequence/deleteSequence [delete]
export const deleteSequenceByIds = (data) => {
  return service({
    url: '/sequence/deleteSequenceByIds',
    method: 'delete',
    data
  })
}

// @Tags Sequence
// @Summary 更新Sequence
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sequence true "更新Sequence"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /sequence/updateSequence [put]
export const updateSequence = (data) => {
  return service({
    url: '/sequence/updateSequence',
    method: 'put',
    data
  })
}

// @Tags Sequence
// @Summary 用idSearchSequence
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Sequence true "用idSearchSequence"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /sequence/findSequence [get]
export const findSequence = (params) => {
  return service({
    url: '/sequence/findSequence',
    method: 'get',
    params
  })
}

// @Tags Sequence
// @Summary 分页获取Sequence列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Sequence列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /sequence/getSequenceList [get]
export const getSequenceList = (params) => {
  return service({
    url: '/sequence/getSequenceList',
    method: 'get',
    params
  })
}

// @Tags Sequence
// @Summary 分页获取Sequence列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Sequence列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /sequence/getSequenceByCampaign [get]
export const getSequenceByCampaign = (params) => {
  return service({
    url: '/sequence/getSequenceByCampaign',
    method: 'get',
    params
  })
}
