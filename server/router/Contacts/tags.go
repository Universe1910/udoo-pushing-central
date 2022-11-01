package Contacts

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TagsRouter struct {
}

// InitTagsRouter 初始化 Tags 路由信息
func (s *TagsRouter) InitTagsRouter(Router *gin.RouterGroup) {
	tagsRouter := Router.Group("tags").Use(middleware.OperationRecord())
	tagsRouterWithoutRecord := Router.Group("tags")
	var tagsApi = v1.ApiGroupApp.ContactsApiGroup.TagsApi
	{
		tagsRouter.POST("createTags", tagsApi.CreateTags)   // 新建Tags
		tagsRouter.DELETE("deleteTags", tagsApi.DeleteTags) // Delete Tags
		tagsRouter.DELETE("deleteTagsByIds", tagsApi.DeleteTagsByIds) // 批量Delete Tags
		tagsRouter.PUT("updateTags", tagsApi.UpdateTags)    // 更新Tags
	}
	{
		tagsRouterWithoutRecord.GET("findTags", tagsApi.FindTags)        // 根据ID获取Tags
		tagsRouterWithoutRecord.GET("getTagsList", tagsApi.GetTagsList)  // 获取Tags列表
	}
}
