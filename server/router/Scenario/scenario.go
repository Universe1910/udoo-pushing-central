package Scenario

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ScenarioRouter struct {
}

// InitScenarioRouter 初始化 Scenario 路由信息
func (s *ScenarioRouter) InitScenarioRouter(Router *gin.RouterGroup) {
	scenarioRouter := Router.Group("scenario").Use(middleware.OperationRecord())
	scenarioRouterWithoutRecord := Router.Group("scenario")
	var scenarioApi = v1.ApiGroupApp.ScenarioApiGroup.ScenarioApi
	{
		scenarioRouter.POST("createScenario", scenarioApi.CreateScenario)   // 新建Scenario
		scenarioRouter.DELETE("deleteScenario", scenarioApi.DeleteScenario) // Delete Scenario
		scenarioRouter.DELETE("deleteScenarioByIds", scenarioApi.DeleteScenarioByIds) // 批量Delete Scenario
		scenarioRouter.PUT("updateScenario", scenarioApi.UpdateScenario)    // 更新Scenario
	}
	{
		scenarioRouterWithoutRecord.GET("findScenario", scenarioApi.FindScenario)        // 根据ID获取Scenario
		scenarioRouterWithoutRecord.GET("getScenarioList", scenarioApi.GetScenarioList)  // 获取Scenario列表
	}
}
