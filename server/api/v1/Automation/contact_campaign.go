package Automation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Automation"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    AutomationReq "github.com/flipped-aurora/gin-vue-admin/server/model/Automation/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type ContactCampaignApi struct {
}

var contactCampaignService = service.ServiceGroupApp.AutomationServiceGroup.ContactCampaignService


// CreateContactCampaign Create ContactCampaign
// @Tags ContactCampaign
// @Summary Create ContactCampaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Automation.ContactCampaign true "Create ContactCampaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /contactCampaign/createContactCampaign [post]
func (contactCampaignApi *ContactCampaignApi) CreateContactCampaign(c *gin.Context) {
	var contactCampaign Automation.ContactCampaign
	err := c.ShouldBindJSON(&contactCampaign)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "ContactId":{utils.NotEmpty()},
    }
	if err := utils.Verify(contactCampaign, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := contactCampaignService.CreateContactCampaign(contactCampaign); err != nil {
        global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create", c)
	} else {
		response.OkWithMessage("Successful creation", c)
	}
}

// DeleteContactCampaign Delete ContactCampaign
// @Tags ContactCampaign
// @Summary Delete ContactCampaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Automation.ContactCampaign true "Delete ContactCampaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /contactCampaign/deleteContactCampaign [delete]
func (contactCampaignApi *ContactCampaignApi) DeleteContactCampaign(c *gin.Context) {
	var contactCampaign Automation.ContactCampaign
	err := c.ShouldBindJSON(&contactCampaign)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := contactCampaignService.DeleteContactCampaign(contactCampaign); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// DeleteContactCampaignByIds Delete ContactCampaign
// @Tags ContactCampaign
// @Summary Delete ContactCampaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete ContactCampaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /contactCampaign/deleteContactCampaignByIds [delete]
func (contactCampaignApi *ContactCampaignApi) DeleteContactCampaignByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := contactCampaignService.DeleteContactCampaignByIds(IDS); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// UpdateContactCampaign 更新ContactCampaign
// @Tags ContactCampaign
// @Summary 更新ContactCampaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Automation.ContactCampaign true "更新ContactCampaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /contactCampaign/updateContactCampaign [put]
func (contactCampaignApi *ContactCampaignApi) UpdateContactCampaign(c *gin.Context) {
	var contactCampaign Automation.ContactCampaign
	err := c.ShouldBindJSON(&contactCampaign)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "ContactId":{utils.NotEmpty()},
      }
    if err := utils.Verify(contactCampaign, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := contactCampaignService.UpdateContactCampaign(contactCampaign); err != nil {
        global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
	} else {
		response.OkWithMessage("update completed", c)
	}
}

// FindContactCampaign 用idSearchContactCampaign
// @Tags ContactCampaign
// @Summary 用idSearchContactCampaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Automation.ContactCampaign true "用idSearchContactCampaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /contactCampaign/findContactCampaign [get]
func (contactCampaignApi *ContactCampaignApi) FindContactCampaign(c *gin.Context) {
	var contactCampaign Automation.ContactCampaign
	err := c.ShouldBindQuery(&contactCampaign)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recontactCampaign, err := contactCampaignService.GetContactCampaign(contactCampaign.ID); err != nil {
        global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		response.OkWithData(gin.H{"recontactCampaign": recontactCampaign}, c)
	}
}

// GetContactCampaignList 分页获取ContactCampaign列表
// @Tags ContactCampaign
// @Summary 分页获取ContactCampaign列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query AutomationReq.ContactCampaignSearch true "分页获取ContactCampaign列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /contactCampaign/getContactCampaignList [get]
func (contactCampaignApi *ContactCampaignApi) GetContactCampaignList(c *gin.Context) {
	var pageInfo AutomationReq.ContactCampaignSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := contactCampaignService.GetContactCampaignInfoList(pageInfo); err != nil {
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
