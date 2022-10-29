package EmailMarketing

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EmailTemplateRouter struct {
}

// InitEmailTemplateRouter 初始化 EmailTemplate 路由信息
func (s *EmailTemplateRouter) InitEmailTemplateRouter(Router *gin.RouterGroup) {
	email_templateRouter := Router.Group("email_template").Use(middleware.OperationRecord())
	email_templateRouterWithoutRecord := Router.Group("email_template")
	var email_templateApi = v1.ApiGroupApp.EmailmarketingApiGroup.EmailTemplateApi
	{
		email_templateRouter.POST("createEmailTemplate", email_templateApi.CreateEmailTemplate)             // 新建EmailTemplate
		email_templateRouter.DELETE("deleteEmailTemplate", email_templateApi.DeleteEmailTemplate)           // Delete EmailTemplate
		email_templateRouter.DELETE("deleteEmailTemplateByIds", email_templateApi.DeleteEmailTemplateByIds) // 批量Delete EmailTemplate
		email_templateRouter.PUT("updateEmailTemplate", email_templateApi.UpdateEmailTemplate)              // 更新EmailTemplate
	}
	{
		email_templateRouterWithoutRecord.GET("findEmailTemplate", email_templateApi.FindEmailTemplate)       // 根据ID获取EmailTemplate
		email_templateRouterWithoutRecord.GET("getEmailTemplateList", email_templateApi.GetEmailTemplateList) // 获取EmailTemplate列表
	}
}
