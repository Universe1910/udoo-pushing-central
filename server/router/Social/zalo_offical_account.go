package Social

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ZaloOfficalAccountRouter struct {
}

// InitZaloOfficalAccountRouter 初始化 ZaloOfficalAccount 路由信息
func (s *ZaloOfficalAccountRouter) InitZaloOfficalAccountRouter(Router *gin.RouterGroup) {
	zaloOfficalAccountRouter := Router.Group("zaloOfficalAccount").Use(middleware.OperationRecord())
	zaloOfficalAccountRouterWithoutRecord := Router.Group("zaloOfficalAccount")
	var zaloOfficalAccountApi = v1.ApiGroupApp.SocialApiGroup.ZaloOfficalAccountApi
	{
		zaloOfficalAccountRouter.POST("createZaloOfficalAccount", zaloOfficalAccountApi.CreateZaloOfficalAccount)   // 新建ZaloOfficalAccount
		zaloOfficalAccountRouter.DELETE("deleteZaloOfficalAccount", zaloOfficalAccountApi.DeleteZaloOfficalAccount) // Delete ZaloOfficalAccount
		zaloOfficalAccountRouter.DELETE("deleteZaloOfficalAccountByIds", zaloOfficalAccountApi.DeleteZaloOfficalAccountByIds) // 批量Delete ZaloOfficalAccount
		zaloOfficalAccountRouter.PUT("updateZaloOfficalAccount", zaloOfficalAccountApi.UpdateZaloOfficalAccount)    // 更新ZaloOfficalAccount
	}
	{
		zaloOfficalAccountRouterWithoutRecord.GET("findZaloOfficalAccount", zaloOfficalAccountApi.FindZaloOfficalAccount)        // 根据ID获取ZaloOfficalAccount
		zaloOfficalAccountRouterWithoutRecord.GET("getZaloOfficalAccountList", zaloOfficalAccountApi.GetZaloOfficalAccountList)  // 获取ZaloOfficalAccount列表
	}
}
