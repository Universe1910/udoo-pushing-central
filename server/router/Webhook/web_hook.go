package Webhook

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type WebhookRouter struct {
}

// InitCampaignRouter 初始化 Campaign 路由信息
func (s *WebhookRouter) InitWebhookRouter(Router *gin.RouterGroup) {
	// campaignRouter := Router.Group("campaign").Use(middleware.OperationRecord())
	webhookWithoutRecord := Router.Group("webhook")
	// var campaignApi = v1.ApiGroupApp.AutomationApiGroup.CampaignApi
	var webhookApi = v1.ApiGroupApp.WebhookApiGroup.WebhookApi
	// {
	// campaignRouter.POST("createCampaign", campaignApi.CreateCampaign)             // 新建Campaign
	// campaignRouter.DELETE("deleteCampaign", campaignApi.DeleteCampaign)           // Delete Campaign
	// campaignRouter.DELETE("deleteCampaignByIds", campaignApi.DeleteCampaignByIds) // 批量Delete Campaign
	// campaignRouter.PUT("updateCampaign", campaignApi.UpdateCampaign)              // 更新Campaign
	// }
	// {
	webhookWithoutRecord.POST("facebookMessgenerHook", webhookApi.FacebookMessgenerWebHook)
	// 	webhookWithoutRecord.POST("debugSequence", campaignApi.DebugSequence)
	// 	webhookWithoutRecord.POST("startSequence", campaignApi.StartSequence)
	// 	webhookWithoutRecord.GET("findCampaign", campaignApi.FindCampaign)       // 根据ID获取Campaign
	// 	webhookWithoutRecord.GET("getCampaignList", campaignApi.GetCampaignList) // 获取Campaign列表
	// }
}
