package Automation

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SequenceRouter struct {
}

// InitSequenceRouter 初始化 Sequence 路由信息
func (s *SequenceRouter) InitSequenceRouter(Router *gin.RouterGroup) {
	sequenceRouter := Router.Group("sequence").Use(middleware.OperationRecord())
	sequenceRouterWithoutRecord := Router.Group("sequence")
	var sequenceApi = v1.ApiGroupApp.AutomationApiGroup.SequenceApi
	{
		sequenceRouter.POST("createSequence", sequenceApi.CreateSequence)             // 新建Sequence
		sequenceRouter.DELETE("deleteSequence", sequenceApi.DeleteSequence)           // Delete Sequence
		sequenceRouter.DELETE("deleteSequenceByIds", sequenceApi.DeleteSequenceByIds) // 批量Delete Sequence
		sequenceRouter.PUT("updateSequence", sequenceApi.UpdateSequence)              // 更新Sequence
	}
	{
		sequenceRouterWithoutRecord.GET("findSequence", sequenceApi.FindSequence)                   // 根据ID获取Sequence
		sequenceRouterWithoutRecord.GET("getSequenceList", sequenceApi.GetSequenceList)             // 获取Sequence列表
		sequenceRouterWithoutRecord.GET("getSequenceByCampaign", sequenceApi.GetSequenceByCampaign) // 获取Sequence列表
	}
}
