package Social

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ZaloApplicationRouter struct {
}

// InitZaloApplicationRouter 初始化 ZaloApplication 路由信息
func (s *ZaloApplicationRouter) InitZaloApplicationRouter(Router *gin.RouterGroup) {
	zaloApplicationRouter := Router.Group("zaloApplication").Use(middleware.OperationRecord())
	zaloApplicationRouterWithoutRecord := Router.Group("zaloApplication")
	var zaloApplicationApi = v1.ApiGroupApp.SocialApiGroup.ZaloApplicationApi
	{
		zaloApplicationRouter.POST("createZaloApplication", zaloApplicationApi.CreateZaloApplication)             // 新建ZaloApplication
		zaloApplicationRouter.DELETE("deleteZaloApplication", zaloApplicationApi.DeleteZaloApplication)           // Delete ZaloApplication
		zaloApplicationRouter.DELETE("deleteZaloApplicationByIds", zaloApplicationApi.DeleteZaloApplicationByIds) // 批量Delete ZaloApplication
		zaloApplicationRouter.PUT("updateZaloApplication", zaloApplicationApi.UpdateZaloApplication)              // 更新ZaloApplication
	}
	{
		zaloApplicationRouterWithoutRecord.GET("findZaloApplication", zaloApplicationApi.FindZaloApplication)                   // 根据ID获取ZaloApplication
		zaloApplicationRouterWithoutRecord.GET("getZaloApplicationList", zaloApplicationApi.GetZaloApplicationList)             // 获取ZaloApplication列表
		zaloApplicationRouterWithoutRecord.GET("getZaloApplicationActiveList", zaloApplicationApi.GetZaloApplicationActiveList) // ZaloApplication Active
	}
}
