package {{.Package}}

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/{{.Package}}"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    {{.Package}}Req "github.com/flipped-aurora/gin-vue-admin/server/model/{{.Package}}/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    {{- if .NeedValid }}
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
    {{- else if .AutoCreateResource}}
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
    {{- end }}
)

type {{.StructName}}Api struct {
}

var {{.Abbreviation}}Service = service.ServiceGroupApp.{{.PackageT}}ServiceGroup.{{.StructName}}Service


// Create{{.StructName}} Create {{.StructName}}
// @Tags {{.StructName}}
// @Summary Create {{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {{.Package}}.{{.StructName}} true "Create {{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /{{.Abbreviation}}/create{{.StructName}} [post]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Create{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} {{.Package}}.{{.StructName}}
	err := c.ShouldBindJSON(&{{.Abbreviation}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	{{- if .AutoCreateResource }}
    {{.Abbreviation}}.CreatedBy = utils.GetUserID(c)
	{{- end }}
    {{- if .NeedValid }}
    verify := utils.Rules{
    {{- range $index,$element := .Fields }}
       {{- if $element.Require }}
        "{{$element.FieldName}}":{utils.NotEmpty()},
        {{- end }}
    {{- end }}
    }
	if err := utils.Verify({{.Abbreviation}}, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
    {{- end }}
	if err := {{.Abbreviation}}Service.Create{{.StructName}}({{.Abbreviation}}); err != nil {
        global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create", c)
	} else {
		response.OkWithMessage("Successful creation", c)
	}
}

// Delete{{.StructName}} Delete {{.StructName}}
// @Tags {{.StructName}}
// @Summary Delete {{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {{.Package}}.{{.StructName}} true "Delete {{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /{{.Abbreviation}}/delete{{.StructName}} [delete]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Delete{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} {{.Package}}.{{.StructName}}
	err := c.ShouldBindJSON(&{{.Abbreviation}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
		{{- if .AutoCreateResource }}
    {{.Abbreviation}}.DeletedBy = utils.GetUserID(c)
        {{- end }}
	if err := {{.Abbreviation}}Service.Delete{{.StructName}}({{.Abbreviation}}); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// Delete{{.StructName}}ByIds Delete {{.StructName}}
// @Tags {{.StructName}}
// @Summary Delete {{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete {{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /{{.Abbreviation}}/delete{{.StructName}}ByIds [delete]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Delete{{.StructName}}ByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    	{{- if .AutoCreateResource }}
    deletedBy := utils.GetUserID(c)
        {{- end }}
	if err := {{.Abbreviation}}Service.Delete{{.StructName}}ByIds(IDS{{- if .AutoCreateResource }},deletedBy{{- end }}); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// Update{{.StructName}} 更新{{.StructName}}
// @Tags {{.StructName}}
// @Summary 更新{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {{.Package}}.{{.StructName}} true "更新{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /{{.Abbreviation}}/update{{.StructName}} [put]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Update{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} {{.Package}}.{{.StructName}}
	err := c.ShouldBindJSON(&{{.Abbreviation}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	    {{- if .AutoCreateResource }}
    {{.Abbreviation}}.UpdatedBy = utils.GetUserID(c)
        {{- end }}
	{{- if .NeedValid }}
      verify := utils.Rules{
      {{- range $index,$element := .Fields }}
         {{- if $element.Require }}
          "{{$element.FieldName}}":{utils.NotEmpty()},
          {{- end }}
      {{- end }}
      }
    if err := utils.Verify({{.Abbreviation}}, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
    {{- end }}
	if err := {{.Abbreviation}}Service.Update{{.StructName}}({{.Abbreviation}}); err != nil {
        global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
	} else {
		response.OkWithMessage("update completed", c)
	}
}

// Find{{.StructName}} 用id查询{{.StructName}}
// @Tags {{.StructName}}
// @Summary 用id查询{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query {{.Package}}.{{.StructName}} true "用id查询{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /{{.Abbreviation}}/find{{.StructName}} [get]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Find{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} {{.Package}}.{{.StructName}}
	err := c.ShouldBindQuery(&{{.Abbreviation}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if re{{.Abbreviation}}, err := {{.Abbreviation}}Service.Get{{.StructName}}({{.Abbreviation}}.ID); err != nil {
        global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		response.OkWithData(gin.H{"re{{.Abbreviation}}": re{{.Abbreviation}}}, c)
	}
}

// Get{{.StructName}}List 分页获取{{.StructName}}列表
// @Tags {{.StructName}}
// @Summary 分页获取{{.StructName}}列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query {{.Package}}Req.{{.StructName}}Search true "分页获取{{.StructName}}列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /{{.Abbreviation}}/get{{.StructName}}List [get]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Get{{.StructName}}List(c *gin.Context) {
	var pageInfo {{.Package}}Req.{{.StructName}}Search
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := {{.Abbreviation}}Service.Get{{.StructName}}InfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("Fail!", zap.Error(err))
        response.FailWithMessage("Fail", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "Successful", c)
    }
}
