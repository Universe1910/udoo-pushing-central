package Social

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Social"
	SocialReq "github.com/flipped-aurora/gin-vue-admin/server/model/Social/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ZaloOfficalAccountApi struct {
}

var zaloOfficalAccountService = service.ServiceGroupApp.SocialServiceGroup.ZaloOfficalAccountService

// CreateZaloOfficalAccount Create ZaloOfficalAccount
// @Tags ZaloOfficalAccount
// @Summary Create ZaloOfficalAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Social.ZaloOfficalAccount true "Create ZaloOfficalAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /zaloOfficalAccount/createZaloOfficalAccount [post]
func (zaloOfficalAccountApi *ZaloOfficalAccountApi) CreateZaloOfficalAccount(c *gin.Context) {
	var zaloOfficalAccount Social.ZaloOfficalAccount
	err := c.ShouldBindJSON(&zaloOfficalAccount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Oaid":            {utils.NotEmpty()},
		"OfficialAccount": {utils.NotEmpty()},
	}
	if err := utils.Verify(zaloOfficalAccount, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zaloOfficalAccountService.CreateZaloOfficalAccount(zaloOfficalAccount); err != nil {
		global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create", c)
	} else {
		response.OkWithMessage("Successful creation", c)
	}
}

// DeleteZaloOfficalAccount Delete ZaloOfficalAccount
// @Tags ZaloOfficalAccount
// @Summary Delete ZaloOfficalAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Social.ZaloOfficalAccount true "Delete ZaloOfficalAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /zaloOfficalAccount/deleteZaloOfficalAccount [delete]
func (zaloOfficalAccountApi *ZaloOfficalAccountApi) DeleteZaloOfficalAccount(c *gin.Context) {
	var zaloOfficalAccount Social.ZaloOfficalAccount
	err := c.ShouldBindJSON(&zaloOfficalAccount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zaloOfficalAccountService.DeleteZaloOfficalAccount(zaloOfficalAccount); err != nil {
		global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// DeleteZaloOfficalAccountByIds Delete ZaloOfficalAccount
// @Tags ZaloOfficalAccount
// @Summary Delete ZaloOfficalAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete ZaloOfficalAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /zaloOfficalAccount/deleteZaloOfficalAccountByIds [delete]
func (zaloOfficalAccountApi *ZaloOfficalAccountApi) DeleteZaloOfficalAccountByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zaloOfficalAccountService.DeleteZaloOfficalAccountByIds(IDS); err != nil {
		global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// UpdateZaloOfficalAccount 更新ZaloOfficalAccount
// @Tags ZaloOfficalAccount
// @Summary 更新ZaloOfficalAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Social.ZaloOfficalAccount true "更新ZaloOfficalAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /zaloOfficalAccount/updateZaloOfficalAccount [put]
func (zaloOfficalAccountApi *ZaloOfficalAccountApi) UpdateZaloOfficalAccount(c *gin.Context) {
	var zaloOfficalAccount Social.ZaloOfficalAccount
	err := c.ShouldBindJSON(&zaloOfficalAccount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Oaid":            {utils.NotEmpty()},
		"OfficialAccount": {utils.NotEmpty()},
	}
	if err := utils.Verify(zaloOfficalAccount, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zaloOfficalAccountService.UpdateZaloOfficalAccount(zaloOfficalAccount); err != nil {
		global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
	} else {
		response.OkWithMessage("update completed", c)
	}
}

// FindZaloOfficalAccount 用idSearchZaloOfficalAccount
// @Tags ZaloOfficalAccount
// @Summary 用idSearchZaloOfficalAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Social.ZaloOfficalAccount true "用idSearchZaloOfficalAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /zaloOfficalAccount/findZaloOfficalAccount [get]
func (zaloOfficalAccountApi *ZaloOfficalAccountApi) FindZaloOfficalAccount(c *gin.Context) {
	var zaloOfficalAccount Social.ZaloOfficalAccount
	err := c.ShouldBindQuery(&zaloOfficalAccount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rezaloOfficalAccount, err := zaloOfficalAccountService.GetZaloOfficalAccount(zaloOfficalAccount.ID); err != nil {
		global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		response.OkWithData(gin.H{"rezaloOfficalAccount": rezaloOfficalAccount}, c)
	}
}

// GetZaloOfficalAccountList 分页获取ZaloOfficalAccount列表
// @Tags ZaloOfficalAccount
// @Summary 分页获取ZaloOfficalAccount列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query SocialReq.ZaloOfficalAccountSearch true "分页获取ZaloOfficalAccount列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /zaloOfficalAccount/getZaloOfficalAccountList [get]
func (zaloOfficalAccountApi *ZaloOfficalAccountApi) GetZaloOfficalAccountList(c *gin.Context) {
	var pageInfo SocialReq.ZaloOfficalAccountSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := zaloOfficalAccountService.GetZaloOfficalAccountInfoList(pageInfo); err != nil {
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
