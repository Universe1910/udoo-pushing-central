package Automation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ContactCampaignRouter struct {
}

// InitContactCampaignRouter 初始化 ContactCampaign 路由信息
func (s *ContactCampaignRouter) InitContactCampaignRouter(Router *gin.RouterGroup) {
	contactCampaignRouter := Router.Group("contactCampaign").Use(middleware.OperationRecord())
	contactCampaignRouterWithoutRecord := Router.Group("contactCampaign")
	var contactCampaignApi = v1.ApiGroupApp.AutomationApiGroup.ContactCampaignApi
	{
		contactCampaignRouter.POST("createContactCampaign", contactCampaignApi.CreateContactCampaign)   // 新建ContactCampaign
		contactCampaignRouter.DELETE("deleteContactCampaign", contactCampaignApi.DeleteContactCampaign) // Delete ContactCampaign
		contactCampaignRouter.DELETE("deleteContactCampaignByIds", contactCampaignApi.DeleteContactCampaignByIds) // 批量Delete ContactCampaign
		contactCampaignRouter.PUT("updateContactCampaign", contactCampaignApi.UpdateContactCampaign)    // 更新ContactCampaign
	}
	{
		contactCampaignRouterWithoutRecord.GET("findContactCampaign", contactCampaignApi.FindContactCampaign)        // 根据ID获取ContactCampaign
		contactCampaignRouterWithoutRecord.GET("getContactCampaignList", contactCampaignApi.GetContactCampaignList)  // 获取ContactCampaign列表
	}
}
