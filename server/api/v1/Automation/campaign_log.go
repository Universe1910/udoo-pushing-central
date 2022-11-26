package Automation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Automation"
	AutomationReq "github.com/flipped-aurora/gin-vue-admin/server/model/Automation/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CampaignLogApi struct {
}

var campaignLogService = service.ServiceGroupApp.AutomationServiceGroup.CampaignLogService

// CreateCampaignLog Create CampaignLog
// @Tags CampaignLog
// @Summary Create CampaignLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Automation.CampaignLog true "Create CampaignLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /campaignLog/createCampaignLog [post]
func (campaignLogApi *CampaignLogApi) CreateCampaignLog(c *gin.Context) {
	var campaignLog Automation.CampaignLog
	err := c.ShouldBindJSON(&campaignLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := campaignLogService.CreateCampaignLog(campaignLog); err != nil {
		global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create", c)
	} else {
		response.OkWithMessage("Successful creation", c)
	}
}

// DeleteCampaignLog Delete CampaignLog
// @Tags CampaignLog
// @Summary Delete CampaignLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Automation.CampaignLog true "Delete CampaignLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /campaignLog/deleteCampaignLog [delete]
func (campaignLogApi *CampaignLogApi) DeleteCampaignLog(c *gin.Context) {
	var campaignLog Automation.CampaignLog
	err := c.ShouldBindJSON(&campaignLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := campaignLogService.DeleteCampaignLog(campaignLog); err != nil {
		global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// DeleteCampaignLogByIds Delete CampaignLog
// @Tags CampaignLog
// @Summary Delete CampaignLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete CampaignLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /campaignLog/deleteCampaignLogByIds [delete]
func (campaignLogApi *CampaignLogApi) DeleteCampaignLogByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := campaignLogService.DeleteCampaignLogByIds(IDS); err != nil {
		global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// UpdateCampaignLog 更新CampaignLog
// @Tags CampaignLog
// @Summary 更新CampaignLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Automation.CampaignLog true "更新CampaignLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /campaignLog/updateCampaignLog [put]
func (campaignLogApi *CampaignLogApi) UpdateCampaignLog(c *gin.Context) {
	var campaignLog Automation.CampaignLog
	err := c.ShouldBindJSON(&campaignLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := campaignLogService.UpdateCampaignLog(campaignLog); err != nil {
		global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
	} else {
		response.OkWithMessage("update completed", c)
	}
}

// FindCampaignLog 用idSearchCampaignLog
// @Tags CampaignLog
// @Summary 用idSearchCampaignLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Automation.CampaignLog true "用idSearchCampaignLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /campaignLog/findCampaignLog [get]
func (campaignLogApi *CampaignLogApi) FindCampaignLog(c *gin.Context) {
	var campaignLog Automation.CampaignLog
	err := c.ShouldBindQuery(&campaignLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recampaignLog, err := campaignLogService.GetCampaignLog(campaignLog.ID); err != nil {
		global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		response.OkWithData(gin.H{"recampaignLog": recampaignLog}, c)
	}
}

// FindCampaignLogByCampaign
// @Tags CampaignLog
// @Summary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Automation.CampaignLog true "用idSearchCampaignLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /campaignLog/findCampaignLogByCampaign [get]
func (campaignLogApi *CampaignLogApi) FindCampaignLogByCampaign(c *gin.Context) {
	var campaignLog Automation.CampaignLog
	err := c.ShouldBindQuery(&campaignLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recampaignLog, err := campaignLogService.GetCampaignLogByCampaign(campaignLog.CampaignID); err != nil {
		global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		response.OkWithData(gin.H{"recampaignLog": recampaignLog}, c)
	}
}

// GetCampaignLogList 分页获取CampaignLog列表
// @Tags CampaignLog
// @Summary 分页获取CampaignLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query AutomationReq.CampaignLogSearch true "分页获取CampaignLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /campaignLog/getCampaignLogList [get]
func (campaignLogApi *CampaignLogApi) GetCampaignLogList(c *gin.Context) {
	var pageInfo AutomationReq.CampaignLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := campaignLogService.GetCampaignLogInfoList(pageInfo); err != nil {
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
