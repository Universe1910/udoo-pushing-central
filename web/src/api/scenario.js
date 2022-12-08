import service from '@/utils/request'

// @Tags Scenario
// @Summary Create Scenario
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Scenario true "Create Scenario"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /scenario/createScenario [post]
export const createScenario = (data) => {
  return service({
    url: '/scenario/createScenario',
    method: 'post',
    data
  })
}

// @Tags Scenario
// @Summary Delete Scenario
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Scenario true "Delete Scenario"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /scenario/deleteScenario [delete]
export const deleteScenario = (data) => {
  return service({
    url: '/scenario/deleteScenario',
    method: 'delete',
    data
  })
}

// @Tags Scenario
// @Summary Delete Scenario
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量Delete Scenario"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /scenario/deleteScenario [delete]
export const deleteScenarioByIds = (data) => {
  return service({
    url: '/scenario/deleteScenarioByIds',
    method: 'delete',
    data
  })
}

// @Tags Scenario
// @Summary 更新Scenario
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Scenario true "更新Scenario"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /scenario/updateScenario [put]
export const updateScenario = (data) => {
  return service({
    url: '/scenario/updateScenario',
    method: 'put',
    data
  })
}

// @Tags Scenario
// @Summary 用idSearchScenario
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Scenario true "用idSearchScenario"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /scenario/findScenario [get]
export const findScenario = (params) => {
  return service({
    url: '/scenario/findScenario',
    method: 'get',
    params
  })
}

// @Tags Scenario
// @Summary 分页获取Scenario列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Scenario列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /scenario/getScenarioList [get]
export const getScenarioList = (params) => {
  return service({
    url: '/scenario/getScenarioList',
    method: 'get',
    params
  })
}
