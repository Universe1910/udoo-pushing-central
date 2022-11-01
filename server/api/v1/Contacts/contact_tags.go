package Contacts

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Contacts"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ContactsReq "github.com/flipped-aurora/gin-vue-admin/server/model/Contacts/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type ContactTagsApi struct {
}

var contactTagsService = service.ServiceGroupApp.ContactsServiceGroup.ContactTagsService


// CreateContactTags Create ContactTags
// @Tags ContactTags
// @Summary Create ContactTags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Contacts.ContactTags true "Create ContactTags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /contactTags/createContactTags [post]
func (contactTagsApi *ContactTagsApi) CreateContactTags(c *gin.Context) {
	var contactTags Contacts.ContactTags
	err := c.ShouldBindJSON(&contactTags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "ContactID":{utils.NotEmpty()},
        "TagID":{utils.NotEmpty()},
    }
	if err := utils.Verify(contactTags, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := contactTagsService.CreateContactTags(contactTags); err != nil {
        global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create", c)
	} else {
		response.OkWithMessage("Successful creation", c)
	}
}

// DeleteContactTags Delete ContactTags
// @Tags ContactTags
// @Summary Delete ContactTags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Contacts.ContactTags true "Delete ContactTags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /contactTags/deleteContactTags [delete]
func (contactTagsApi *ContactTagsApi) DeleteContactTags(c *gin.Context) {
	var contactTags Contacts.ContactTags
	err := c.ShouldBindJSON(&contactTags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := contactTagsService.DeleteContactTags(contactTags); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// DeleteContactTagsByIds Delete ContactTags
// @Tags ContactTags
// @Summary Delete ContactTags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete ContactTags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /contactTags/deleteContactTagsByIds [delete]
func (contactTagsApi *ContactTagsApi) DeleteContactTagsByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := contactTagsService.DeleteContactTagsByIds(IDS); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// UpdateContactTags 更新ContactTags
// @Tags ContactTags
// @Summary 更新ContactTags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Contacts.ContactTags true "更新ContactTags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /contactTags/updateContactTags [put]
func (contactTagsApi *ContactTagsApi) UpdateContactTags(c *gin.Context) {
	var contactTags Contacts.ContactTags
	err := c.ShouldBindJSON(&contactTags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "ContactID":{utils.NotEmpty()},
          "TagID":{utils.NotEmpty()},
      }
    if err := utils.Verify(contactTags, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := contactTagsService.UpdateContactTags(contactTags); err != nil {
        global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
	} else {
		response.OkWithMessage("update completed", c)
	}
}

// FindContactTags 用id查询ContactTags
// @Tags ContactTags
// @Summary 用id查询ContactTags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Contacts.ContactTags true "用id查询ContactTags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /contactTags/findContactTags [get]
func (contactTagsApi *ContactTagsApi) FindContactTags(c *gin.Context) {
	var contactTags Contacts.ContactTags
	err := c.ShouldBindQuery(&contactTags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recontactTags, err := contactTagsService.GetContactTags(contactTags.ID); err != nil {
        global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		response.OkWithData(gin.H{"recontactTags": recontactTags}, c)
	}
}

// GetContactTagsList 分页获取ContactTags列表
// @Tags ContactTags
// @Summary 分页获取ContactTags列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ContactsReq.ContactTagsSearch true "分页获取ContactTags列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /contactTags/getContactTagsList [get]
func (contactTagsApi *ContactTagsApi) GetContactTagsList(c *gin.Context) {
	var pageInfo ContactsReq.ContactTagsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := contactTagsService.GetContactTagsInfoList(pageInfo); err != nil {
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
