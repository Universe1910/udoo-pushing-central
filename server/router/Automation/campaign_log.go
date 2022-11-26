package Automation

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CampaignLogRouter struct {
}

// InitCampaignLogRouter 初始化 CampaignLog 路由信息
func (s *CampaignLogRouter) InitCampaignLogRouter(Router *gin.RouterGroup) {
	campaignLogRouter := Router.Group("campaignLog").Use(middleware.OperationRecord())
	campaignLogRouterWithoutRecord := Router.Group("campaignLog")
	var campaignLogApi = v1.ApiGroupApp.AutomationApiGroup.CampaignLogApi
	{
		campaignLogRouter.POST("createCampaignLog", campaignLogApi.CreateCampaignLog)             // 新建CampaignLog
		campaignLogRouter.DELETE("deleteCampaignLog", campaignLogApi.DeleteCampaignLog)           // Delete CampaignLog
		campaignLogRouter.DELETE("deleteCampaignLogByIds", campaignLogApi.DeleteCampaignLogByIds) // 批量Delete CampaignLog
		campaignLogRouter.PUT("updateCampaignLog", campaignLogApi.UpdateCampaignLog)              // 更新CampaignLog
	}
	{
		campaignLogRouterWithoutRecord.GET("findCampaignLog", campaignLogApi.FindCampaignLog)
		campaignLogRouterWithoutRecord.GET("findCampaignLogByCampaign", campaignLogApi.FindCampaignLogByCampaign) // 根据ID获取CampaignLog
		campaignLogRouterWithoutRecord.GET("getCampaignLogList", campaignLogApi.GetCampaignLogList)               // 获取CampaignLog列表
	}
}
