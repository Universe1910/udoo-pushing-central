package Automation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Automation"
	AutomationReq "github.com/flipped-aurora/gin-vue-admin/server/model/Automation/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SequenceApi struct {
}

var sequenceService = service.ServiceGroupApp.AutomationServiceGroup.SequenceService

// CreateSequence Create Sequence
// @Tags Sequence
// @Summary Create Sequence
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Automation.Sequence true "Create Sequence"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /sequence/createSequence [post]
func (sequenceApi *SequenceApi) CreateSequence(c *gin.Context) {
	var sequence Automation.Sequence
	err := c.ShouldBindJSON(&sequence)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"CampaignId": {utils.NotEmpty()},
	}
	if err := utils.Verify(sequence, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sequenceService.CreateSequence(&sequence); err != nil {
		global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create", c)
	} else {
		response.OkWithMessageAndData("Successful creation", sequence, c)
	}
}

// DeleteSequence Delete Sequence
// @Tags Sequence
// @Summary Delete Sequence
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Automation.Sequence true "Delete Sequence"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /sequence/deleteSequence [delete]
func (sequenceApi *SequenceApi) DeleteSequence(c *gin.Context) {
	var sequence Automation.Sequence
	err := c.ShouldBindJSON(&sequence)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sequenceService.DeleteSequence(sequence); err != nil {
		global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// DeleteSequenceByIds Delete Sequence
// @Tags Sequence
// @Summary Delete Sequence
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete Sequence"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /sequence/deleteSequenceByIds [delete]
func (sequenceApi *SequenceApi) DeleteSequenceByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sequenceService.DeleteSequenceByIds(IDS); err != nil {
		global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// UpdateSequence 更新Sequence
// @Tags Sequence
// @Summary 更新Sequence
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Automation.Sequence true "更新Sequence"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /sequence/updateSequence [put]
func (sequenceApi *SequenceApi) UpdateSequence(c *gin.Context) {
	var sequence Automation.Sequence
	err := c.ShouldBindJSON(&sequence)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"CampaignId": {utils.NotEmpty()},
		"Date":       {utils.NotEmpty()},
		"Time":       {utils.NotEmpty()},
		"ActionName": {utils.NotEmpty()},
	}
	if err := utils.Verify(sequence, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sequenceService.UpdateSequence(&sequence); err != nil {
		global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
	} else {
		response.OkWithMessageAndData("update completed", sequence, c)
	}
}

// FindSequence 用idSearchSequence
// @Tags Sequence
// @Summary 用idSearchSequence
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Automation.Sequence true "用idSearchSequence"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /sequence/findSequence [get]
func (sequenceApi *SequenceApi) FindSequence(c *gin.Context) {
	var sequence Automation.Sequence
	err := c.ShouldBindQuery(&sequence)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resequence, err := sequenceService.GetSequence(sequence.ID); err != nil {
		global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		response.OkWithData(gin.H{"resequence": resequence}, c)
	}
}

// GetSequenceList 分页获取Sequence列表
// @Tags Sequence
// @Summary 分页获取Sequence列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query AutomationReq.SequenceSearch true "分页获取Sequence列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /sequence/getSequenceList [get]
func (sequenceApi *SequenceApi) GetSequenceList(c *gin.Context) {
	var pageInfo AutomationReq.SequenceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sequenceService.GetSequenceInfoList(pageInfo); err != nil {
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

// GetSequenceList 分页获取Sequence列表
// @Tags Sequence
// @Summary 分页获取Sequence列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query AutomationReq.SequenceSearch true "分页获取Sequence列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /sequence/getSequenceByCampaign [get]
func (sequenceApi *SequenceApi) GetSequenceByCampaign(c *gin.Context) {
	var sequence Automation.Sequence
	err := c.ShouldBindQuery(&sequence)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resequence, err := sequenceService.GetSequenceByCampaign(sequence.CampaignId); err != nil {
		global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		response.OkWithData(gin.H{"resequence": resequence}, c)
	}
}
