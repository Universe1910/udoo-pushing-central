package Contacts

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Automation"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Contacts"
	ContactsReq "github.com/flipped-aurora/gin-vue-admin/server/model/Contacts/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type ContactApi struct {
}

var contactService = service.ServiceGroupApp.ContactsServiceGroup.ContactService
var contactCampaignService = service.ServiceGroupApp.AutomationServiceGroup.ContactCampaignService
var campaignService = service.ServiceGroupApp.AutomationServiceGroup.CampaignService

// CreateContact Create Contact
// @Tags Contact
// @Summary Create Contact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Contacts.ContactRegister true "Create Contact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /contact/createContact [post]
func (contactApi *ContactApi) CreateContact(c *gin.Context) {
	var reg Contacts.ContactRegister
	err := c.ShouldBindJSON(&reg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		//"Uid":{utils.NotEmpty()},
		"Email": {utils.NotEmpty()},
		"Phone": {utils.NotEmpty()},
	}
	if err := utils.Verify(reg, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	contact := Contacts.Contact{
		Firstname: reg.Firstname,
		Lastname:  reg.Lastname,
		Email:     reg.Email,
		Phone:     reg.Phone,
		Fbid:      reg.Fbid,
		Zid:       reg.Zid,
		Address:   reg.Address,
		City:      reg.City,
		State:     reg.State,
		Zipcode:   reg.Zipcode,
		Country:   reg.Country,
	}

	contact.UUID = uuid.NewV4()
	if err := contactService.CreateContact(&contact); err != nil {
		global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create", c)
	} else {
		if reg.CampaignID != nil {
			contactCampaign := Automation.ContactCampaign{
				ContactId:  contact.ID,
				CampaignId: reg.CampaignID,
			}
			if err := contactCampaignService.CreateContactCampaign(contactCampaign); err != nil {
				global.GVA_LOG.Error("Canot add contact to campaign!", zap.Error(err))
				response.FailWithMessage("Canot add contact to campaign", c)
			} else {
				err := campaignService.RunTrigger(*reg.CampaignID, contact)
				if err != nil {
					global.GVA_LOG.Error("Add success but trigger run fail!", zap.Error(err))
					response.FailWithMessage("Add success but trigger run fail", c)
				}
				response.OkWithMessage("Successful creation", c)
			}
		}

	}
}

// DeleteContact Delete Contact
// @Tags Contact
// @Summary Delete Contact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Contacts.Contact true "Delete Contact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /contact/deleteContact [delete]
func (contactApi *ContactApi) DeleteContact(c *gin.Context) {
	var contact Contacts.Contact
	err := c.ShouldBindJSON(&contact)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := contactService.DeleteContact(contact); err != nil {
		global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// DeleteContactByIds Delete Contact
// @Tags Contact
// @Summary Delete Contact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete Contact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /contact/deleteContactByIds [delete]
func (contactApi *ContactApi) DeleteContactByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := contactService.DeleteContactByIds(IDS); err != nil {
		global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// UpdateContact 更新Contact
// @Tags Contact
// @Summary 更新Contact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Contacts.Contact true "更新Contact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /contact/updateContact [put]
func (contactApi *ContactApi) UpdateContact(c *gin.Context) {
	var contact Contacts.Contact
	err := c.ShouldBindJSON(&contact)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Uid":   {utils.NotEmpty()},
		"Email": {utils.NotEmpty()},
		"Phone": {utils.NotEmpty()},
	}
	if err := utils.Verify(contact, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := contactService.UpdateContact(contact); err != nil {
		global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
	} else {
		response.OkWithMessage("update completed", c)
	}
}

// FindContact 用idSearchContact
// @Tags Contact
// @Summary 用idSearchContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Contacts.Contact true "用idSearchContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /contact/findContact [get]
func (contactApi *ContactApi) FindContact(c *gin.Context) {
	var contact Contacts.Contact
	err := c.ShouldBindQuery(&contact)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recontact, err := contactService.GetContact(contact.ID); err != nil {
		global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		response.OkWithData(gin.H{"recontact": recontact}, c)
	}
}

// GetContactList 分页获取Contact列表
// @Tags Contact
// @Summary 分页获取Contact列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ContactsReq.ContactSearch true "分页获取Contact列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /contact/getContactList [get]
func (contactApi *ContactApi) GetContactList(c *gin.Context) {
	var pageInfo ContactsReq.ContactSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := contactService.GetContactInfoList(pageInfo); err != nil {
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
