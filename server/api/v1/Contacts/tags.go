package Contacts

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Contacts"
	ContactsReq "github.com/flipped-aurora/gin-vue-admin/server/model/Contacts/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TagsApi struct {
}

var tagsService = service.ServiceGroupApp.ContactsServiceGroup.TagsService

// CreateTags Create Tags
// @Tags Tags
// @Summary Create Tags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Contacts.Tags true "Create Tags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /tags/createTags [post]
func (tagsApi *TagsApi) CreateTags(c *gin.Context) {
	var tags Contacts.Tags
	err := c.ShouldBindJSON(&tags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(tags, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tagsService.CreateTags(tags); err != nil {
		global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create", c)
	} else {
		response.OkWithMessage("Successful creation", c)
	}
}

// DeleteTags Delete Tags
// @Tags Tags
// @Summary Delete Tags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Contacts.Tags true "Delete Tags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /tags/deleteTags [delete]
func (tagsApi *TagsApi) DeleteTags(c *gin.Context) {
	var tags Contacts.Tags
	err := c.ShouldBindJSON(&tags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tagsService.DeleteTags(tags); err != nil {
		global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// DeleteTagsByIds Delete Tags
// @Tags Tags
// @Summary Delete Tags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete Tags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /tags/deleteTagsByIds [delete]
func (tagsApi *TagsApi) DeleteTagsByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tagsService.DeleteTagsByIds(IDS); err != nil {
		global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// UpdateTags 更新Tags
// @Tags Tags
// @Summary 更新Tags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Contacts.Tags true "更新Tags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /tags/updateTags [put]
func (tagsApi *TagsApi) UpdateTags(c *gin.Context) {
	var tags Contacts.Tags
	err := c.ShouldBindJSON(&tags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(tags, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tagsService.UpdateTags(tags); err != nil {
		global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
	} else {
		response.OkWithMessage("update completed", c)
	}
}

// FindTags 用idSearchTags
// @Tags Tags
// @Summary 用idSearchTags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Contacts.Tags true "用idSearchTags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /tags/findTags [get]
func (tagsApi *TagsApi) FindTags(c *gin.Context) {
	var tags Contacts.Tags
	err := c.ShouldBindQuery(&tags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retags, err := tagsService.GetTags(tags.ID); err != nil {
		global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		response.OkWithData(gin.H{"retags": retags}, c)
	}
}

// GetTagsList 分页获取Tags列表
// @Tags Tags
// @Summary 分页获取Tags列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ContactsReq.TagsSearch true "分页获取Tags列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /tags/getTagsList [get]
func (tagsApi *TagsApi) GetTagsList(c *gin.Context) {
	var pageInfo ContactsReq.TagsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := tagsService.GetTagsInfoList(pageInfo); err != nil {
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
