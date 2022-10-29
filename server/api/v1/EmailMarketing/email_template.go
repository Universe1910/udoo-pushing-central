package EmailMarketing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EmailMarketing"
	EmailMarketingReq "github.com/flipped-aurora/gin-vue-admin/server/model/EmailMarketing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EmailTemplateApi struct {
}

var email_templateService = service.ServiceGroupApp.EmailmarketingServiceGroup.EmailTemplateService

// CreateEmailTemplate Create EmailTemplate
// @Tags EmailTemplate
// @Summary Create EmailTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EmailMarketing.EmailTemplate true "Create EmailTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /email_template/createEmailTemplate [post]
func (email_templateApi *EmailTemplateApi) CreateEmailTemplate(c *gin.Context) {
	var email_template EmailMarketing.EmailTemplate
	err := c.ShouldBindJSON(&email_template)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := email_templateService.CreateEmailTemplate(email_template); err != nil {
		global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create", c)
	} else {
		response.OkWithMessage("Successful creation", c)
	}
}

// DeleteEmailTemplate Delete EmailTemplate
// @Tags EmailTemplate
// @Summary Delete EmailTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EmailMarketing.EmailTemplate true "Delete EmailTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /email_template/deleteEmailTemplate [delete]
func (email_templateApi *EmailTemplateApi) DeleteEmailTemplate(c *gin.Context) {
	var email_template EmailMarketing.EmailTemplate
	err := c.ShouldBindJSON(&email_template)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := email_templateService.DeleteEmailTemplate(email_template); err != nil {
		global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// DeleteEmailTemplateByIds 批量Delete EmailTemplate
// @Tags EmailTemplate
// @Summary 批量Delete EmailTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量Delete EmailTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量successfully deleted"}"
// @Router /email_template/deleteEmailTemplateByIds [delete]
func (email_templateApi *EmailTemplateApi) DeleteEmailTemplateByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := email_templateService.DeleteEmailTemplateByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量failed to delete!", zap.Error(err))
		response.FailWithMessage("批量failed to delete", c)
	} else {
		response.OkWithMessage("批量successfully deleted", c)
	}
}

// UpdateEmailTemplate 更新EmailTemplate
// @Tags EmailTemplate
// @Summary 更新EmailTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EmailMarketing.EmailTemplate true "更新EmailTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /email_template/updateEmailTemplate [put]
func (email_templateApi *EmailTemplateApi) UpdateEmailTemplate(c *gin.Context) {
	var email_template EmailMarketing.EmailTemplate
	err := c.ShouldBindJSON(&email_template)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := email_templateService.UpdateEmailTemplate(email_template); err != nil {
		global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
	} else {
		response.OkWithMessage("update completed", c)
	}
}

// FindEmailTemplate 用id查询EmailTemplate
// @Tags EmailTemplate
// @Summary 用id查询EmailTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EmailMarketing.EmailTemplate true "用id查询EmailTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /email_template/findEmailTemplate [get]
func (email_templateApi *EmailTemplateApi) FindEmailTemplate(c *gin.Context) {
	var email_template EmailMarketing.EmailTemplate
	err := c.ShouldBindQuery(&email_template)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reemail_template, err := email_templateService.GetEmailTemplate(email_template.ID); err != nil {
		global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		response.OkWithData(gin.H{"reemail_template": reemail_template}, c)
	}
}

// GetEmailTemplateList 分页获取EmailTemplate列表
// @Tags EmailTemplate
// @Summary 分页获取EmailTemplate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EmailMarketingReq.EmailTemplateSearch true "分页获取EmailTemplate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /email_template/getEmailTemplateList [get]
func (email_templateApi *EmailTemplateApi) GetEmailTemplateList(c *gin.Context) {
	var pageInfo EmailMarketingReq.EmailTemplateSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := email_templateService.GetEmailTemplateInfoList(pageInfo); err != nil {
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
