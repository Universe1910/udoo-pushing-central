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

type CampaignApi struct {
}

var campaignService = service.ServiceGroupApp.AutomationServiceGroup.CampaignService


// CreateCampaign Create Campaign
// @Tags Campaign
// @Summary Create Campaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Automation.Campaign true "Create Campaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /campaign/createCampaign [post]
func (campaignApi *CampaignApi) CreateCampaign(c *gin.Context) {
	var campaign Automation.Campaign
	err := c.ShouldBindJSON(&campaign)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "CreatedBy":{utils.NotEmpty()},
    }
	if err := utils.Verify(campaign, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := campaignService.CreateCampaign(campaign); err != nil {
        global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create", c)
	} else {
		response.OkWithMessage("Successful creation", c)
	}
}

// DeleteCampaign Delete Campaign
// @Tags Campaign
// @Summary Delete Campaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Automation.Campaign true "Delete Campaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /campaign/deleteCampaign [delete]
func (campaignApi *CampaignApi) DeleteCampaign(c *gin.Context) {
	var campaign Automation.Campaign
	err := c.ShouldBindJSON(&campaign)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := campaignService.DeleteCampaign(campaign); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// DeleteCampaignByIds Delete Campaign
// @Tags Campaign
// @Summary Delete Campaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete Campaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /campaign/deleteCampaignByIds [delete]
func (campaignApi *CampaignApi) DeleteCampaignByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := campaignService.DeleteCampaignByIds(IDS); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// UpdateCampaign 更新Campaign
// @Tags Campaign
// @Summary 更新Campaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Automation.Campaign true "更新Campaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /campaign/updateCampaign [put]
func (campaignApi *CampaignApi) UpdateCampaign(c *gin.Context) {
	var campaign Automation.Campaign
	err := c.ShouldBindJSON(&campaign)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "CreatedBy":{utils.NotEmpty()},
      }
    if err := utils.Verify(campaign, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := campaignService.UpdateCampaign(campaign); err != nil {
        global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
	} else {
		response.OkWithMessage("update completed", c)
	}
}

// FindCampaign 用idSearchCampaign
// @Tags Campaign
// @Summary 用idSearchCampaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Automation.Campaign true "用idSearchCampaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /campaign/findCampaign [get]
func (campaignApi *CampaignApi) FindCampaign(c *gin.Context) {
	var campaign Automation.Campaign
	err := c.ShouldBindQuery(&campaign)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recampaign, err := campaignService.GetCampaign(campaign.ID); err != nil {
        global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		response.OkWithData(gin.H{"recampaign": recampaign}, c)
	}
}

// GetCampaignList 分页获取Campaign列表
// @Tags Campaign
// @Summary 分页获取Campaign列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query AutomationReq.CampaignSearch true "分页获取Campaign列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /campaign/getCampaignList [get]
func (campaignApi *CampaignApi) GetCampaignList(c *gin.Context) {
	var pageInfo AutomationReq.CampaignSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := campaignService.GetCampaignInfoList(pageInfo); err != nil {
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
