import service from '@/utils/request'

// @Tags {{.StructName}}
// @Summary Create {{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "Create {{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /{{.Abbreviation}}/create{{.StructName}} [post]
export const create{{.StructName}} = (data) => {
  return service({
    url: '/{{.Abbreviation}}/create{{.StructName}}',
    method: 'post',
    data
  })
}

// @Tags {{.StructName}}
// @Summary Delete {{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "Delete {{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /{{.Abbreviation}}/delete{{.StructName}} [delete]
export const delete{{.StructName}} = (data) => {
  return service({
    url: '/{{.Abbreviation}}/delete{{.StructName}}',
    method: 'delete',
    data
  })
}

// @Tags {{.StructName}}
// @Summary Delete {{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Batch delete {{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /{{.Abbreviation}}/delete{{.StructName}} [delete]
export const delete{{.StructName}}ByIds = (data) => {
  return service({
    url: '/{{.Abbreviation}}/delete{{.StructName}}ByIds',
    method: 'delete',
    data
  })
}

// @Tags {{.StructName}}
// @Summary Update{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "Update{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /{{.Abbreviation}}/update{{.StructName}} [put]
export const update{{.StructName}} = (data) => {
  return service({
    url: '/{{.Abbreviation}}/update{{.StructName}}',
    method: 'put',
    data
  })
}

// @Tags {{.StructName}}
// @Summary Get by ID{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.{{.StructName}} true "Get by ID{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /{{.Abbreviation}}/find{{.StructName}} [get]
export const find{{.StructName}} = (params) => {
  return service({
    url: '/{{.Abbreviation}}/find{{.StructName}}',
    method: 'get',
    params
  })
}

// @Tags {{.StructName}}
// @Summary get{{.StructName}}List
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "get{{.StructName}}List"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /{{.Abbreviation}}/get{{.StructName}}List [get]
export const get{{.StructName}}List = (params) => {
  return service({
    url: '/{{.Abbreviation}}/get{{.StructName}}List',
    method: 'get',
    params
  })
}
