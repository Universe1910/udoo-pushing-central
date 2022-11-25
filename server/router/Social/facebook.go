package Social

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FacebookRouter struct {
}

// InitFacebookRouter 初始化 Facebook 路由信息
func (s *FacebookRouter) InitFacebookRouter(Router *gin.RouterGroup) {
	facebookRouter := Router.Group("facebook").Use(middleware.OperationRecord())
	facebookRouterWithoutRecord := Router.Group("facebook")
	var facebookApi = v1.ApiGroupApp.SocialApiGroup.FacebookApi
	{
		facebookRouter.POST("createFacebook", facebookApi.CreateFacebook)   // 新建Facebook
		facebookRouter.DELETE("deleteFacebook", facebookApi.DeleteFacebook) // Delete Facebook
		facebookRouter.DELETE("deleteFacebookByIds", facebookApi.DeleteFacebookByIds) // 批量Delete Facebook
		facebookRouter.PUT("updateFacebook", facebookApi.UpdateFacebook)    // 更新Facebook
	}
	{
		facebookRouterWithoutRecord.GET("findFacebook", facebookApi.FindFacebook)        // 根据ID获取Facebook
		facebookRouterWithoutRecord.GET("getFacebookList", facebookApi.GetFacebookList)  // 获取Facebook列表
	}
}
