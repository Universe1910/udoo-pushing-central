package Automation

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CampaignRouter struct {
}

// InitCampaignRouter 初始化 Campaign 路由信息
func (s *CampaignRouter) InitCampaignRouter(Router *gin.RouterGroup) {
	campaignRouter := Router.Group("campaign").Use(middleware.OperationRecord())
	campaignRouterWithoutRecord := Router.Group("campaign")
	var campaignApi = v1.ApiGroupApp.AutomationApiGroup.CampaignApi
	{
		campaignRouter.POST("createCampaign", campaignApi.CreateCampaign)             // 新建Campaign
		campaignRouter.DELETE("deleteCampaign", campaignApi.DeleteCampaign)           // Delete Campaign
		campaignRouter.DELETE("deleteCampaignByIds", campaignApi.DeleteCampaignByIds) // 批量Delete Campaign
		campaignRouter.PUT("updateCampaign", campaignApi.UpdateCampaign)              // 更新Campaign
	}
	{
		campaignRouterWithoutRecord.POST("debugCampaign", campaignApi.DebugCampaign)
		campaignRouterWithoutRecord.GET("findCampaign", campaignApi.FindCampaign)       // 根据ID获取Campaign
		campaignRouterWithoutRecord.GET("getCampaignList", campaignApi.GetCampaignList) // 获取Campaign列表
	}
}
