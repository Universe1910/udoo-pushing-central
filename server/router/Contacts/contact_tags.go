package Contacts

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ContactTagsRouter struct {
}

// InitContactTagsRouter 初始化 ContactTags 路由信息
func (s *ContactTagsRouter) InitContactTagsRouter(Router *gin.RouterGroup) {
	contactTagsRouter := Router.Group("contactTags").Use(middleware.OperationRecord())
	contactTagsRouterWithoutRecord := Router.Group("contactTags")
	var contactTagsApi = v1.ApiGroupApp.ContactsApiGroup.ContactTagsApi
	{
		contactTagsRouter.POST("createContactTags", contactTagsApi.CreateContactTags)   // 新建ContactTags
		contactTagsRouter.DELETE("deleteContactTags", contactTagsApi.DeleteContactTags) // Delete ContactTags
		contactTagsRouter.DELETE("deleteContactTagsByIds", contactTagsApi.DeleteContactTagsByIds) // 批量Delete ContactTags
		contactTagsRouter.PUT("updateContactTags", contactTagsApi.UpdateContactTags)    // 更新ContactTags
	}
	{
		contactTagsRouterWithoutRecord.GET("findContactTags", contactTagsApi.FindContactTags)        // 根据ID获取ContactTags
		contactTagsRouterWithoutRecord.GET("getContactTagsList", contactTagsApi.GetContactTagsList)  // 获取ContactTags列表
	}
}
