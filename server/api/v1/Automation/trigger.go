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

type TriggerApi struct {
}

var triggerService = service.ServiceGroupApp.AutomationServiceGroup.TriggerService

// CreateTrigger Create Trigger
// @Tags Trigger
// @Summary Create Trigger
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Automation.Trigger true "Create Trigger"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /trigger/createTrigger [post]
func (triggerApi *TriggerApi) CreateTrigger(c *gin.Context) {
	var trigger Automation.Trigger
	err := c.ShouldBindJSON(&trigger)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := triggerService.CreateTrigger(&trigger); err != nil {
		global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create", c)
	} else {

		// response.OkWithMessage("Successful creation", c)
		response.OkWithMessageAndData("Successful creation", trigger, c)
	}
}

// DeleteTrigger Delete Trigger
// @Tags Trigger
// @Summary Delete Trigger
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Automation.Trigger true "Delete Trigger"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /trigger/deleteTrigger [delete]
func (triggerApi *TriggerApi) DeleteTrigger(c *gin.Context) {
	var trigger Automation.Trigger
	err := c.ShouldBindJSON(&trigger)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := triggerService.DeleteTrigger(trigger); err != nil {
		global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// DeleteTriggerByIds Delete Trigger
// @Tags Trigger
// @Summary Delete Trigger
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete Trigger"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /trigger/deleteTriggerByIds [delete]
func (triggerApi *TriggerApi) DeleteTriggerByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := triggerService.DeleteTriggerByIds(IDS); err != nil {
		global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// UpdateTrigger 更新Trigger
// @Tags Trigger
// @Summary 更新Trigger
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Automation.Trigger true "更新Trigger"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /trigger/updateTrigger [put]
func (triggerApi *TriggerApi) UpdateTrigger(c *gin.Context) {
	var trigger Automation.Trigger
	err := c.ShouldBindJSON(&trigger)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := triggerService.UpdateTrigger(&trigger); err != nil {
		global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
	} else {
		response.OkWithMessage("update completed", c)
	}
}

// FindTrigger 用idSearchTrigger
// @Tags Trigger
// @Summary 用idSearchTrigger
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Automation.Trigger true "用idSearchTrigger"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /trigger/findTrigger [get]
func (triggerApi *TriggerApi) FindTrigger(c *gin.Context) {
	var trigger Automation.Trigger
	err := c.ShouldBindQuery(&trigger)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retrigger, err := triggerService.GetTrigger(trigger.ID); err != nil {
		global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		response.OkWithData(gin.H{"retrigger": retrigger}, c)
	}
}

// GetTriggerList 分页获取Trigger列表
// @Tags Trigger
// @Summary 分页获取Trigger列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query AutomationReq.TriggerSearch true "分页获取Trigger列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /trigger/getTriggerList [get]
func (triggerApi *TriggerApi) GetTriggerList(c *gin.Context) {
	var pageInfo AutomationReq.TriggerSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := triggerService.GetTriggerInfoList(pageInfo); err != nil {
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
