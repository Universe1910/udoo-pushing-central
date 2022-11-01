package Contacts

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ContactRouter struct {
}

// InitContactRouter 初始化 Contact 路由信息
func (s *ContactRouter) InitContactRouter(Router *gin.RouterGroup) {
	contactRouter := Router.Group("contact").Use(middleware.OperationRecord())
	contactRouterWithoutRecord := Router.Group("contact")
	var contactApi = v1.ApiGroupApp.ContactsApiGroup.ContactApi
	{
		contactRouter.POST("createContact", contactApi.CreateContact)             // 新建Contact
		contactRouter.DELETE("deleteContact", contactApi.DeleteContact)           // Delete Contact
		contactRouter.DELETE("deleteContactByIds", contactApi.DeleteContactByIds) // 批量Delete Contact
		contactRouter.PUT("updateContact", contactApi.UpdateContact)              // 更新Contact
	}
	{
		contactRouterWithoutRecord.GET("findContact", contactApi.FindContact)       // 根据ID获取Contact
		contactRouterWithoutRecord.GET("getContactList", contactApi.GetContactList) // 获取Contact列表
	}
}

func (s *ContactRouter) InitPublicContactRouter(Router *gin.RouterGroup) {
	//contactPublicRouter := Router.Group("publicContact")
	contactPublicRouterWithoutRecord := Router.Group("publicContact")
	var contactApi = v1.ApiGroupApp.ContactsApiGroup.ContactApi
	{
		contactPublicRouterWithoutRecord.POST("create", contactApi.CreateContact)  // 新建Contact
		contactPublicRouterWithoutRecord.GET("getList", contactApi.GetContactList) // 获取Contact列表
	}

}
