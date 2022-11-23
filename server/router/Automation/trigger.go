package Automation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TriggerRouter struct {
}

// InitTriggerRouter 初始化 Trigger 路由信息
func (s *TriggerRouter) InitTriggerRouter(Router *gin.RouterGroup) {
	triggerRouter := Router.Group("trigger").Use(middleware.OperationRecord())
	triggerRouterWithoutRecord := Router.Group("trigger")
	var triggerApi = v1.ApiGroupApp.AutomationApiGroup.TriggerApi
	{
		triggerRouter.POST("createTrigger", triggerApi.CreateTrigger)   // 新建Trigger
		triggerRouter.DELETE("deleteTrigger", triggerApi.DeleteTrigger) // Delete Trigger
		triggerRouter.DELETE("deleteTriggerByIds", triggerApi.DeleteTriggerByIds) // 批量Delete Trigger
		triggerRouter.PUT("updateTrigger", triggerApi.UpdateTrigger)    // 更新Trigger
	}
	{
		triggerRouterWithoutRecord.GET("findTrigger", triggerApi.FindTrigger)        // 根据ID获取Trigger
		triggerRouterWithoutRecord.GET("getTriggerList", triggerApi.GetTriggerList)  // 获取Trigger列表
	}
}
