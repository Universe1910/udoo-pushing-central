package Social

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Social"
	SocialReq "github.com/flipped-aurora/gin-vue-admin/server/model/Social/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/service/zalo"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ZaloApplicationApi struct {
}

var zaloApplicationService = service.ServiceGroupApp.SocialServiceGroup.ZaloApplicationService
var zaloNotificationServiceAPI zalo.ZaloNotificationServiceAPI

// CreateZaloApplication Create ZaloApplication
// @Tags ZaloApplication
// @Summary Create ZaloApplication
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Social.ZaloApplication true "Create ZaloApplication"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /zaloApplication/createZaloApplication [post]
func (zaloApplicationApi *ZaloApplicationApi) CreateZaloApplication(c *gin.Context) {
	var zaloApplication Social.ZaloApplication
	err := c.ShouldBindJSON(&zaloApplication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"ApplicationName": {utils.NotEmpty()},
		"ApplicationID":   {utils.NotEmpty()},
		"SecretKey":       {utils.NotEmpty()},
	}
	if err := utils.Verify(zaloApplication, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zaloApplicationService.CreateZaloApplication(zaloApplication); err != nil {
		global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create", c)
	} else {
		response.OkWithMessage("Successful creation", c)
	}
}

// DeleteZaloApplication Delete ZaloApplication
// @Tags ZaloApplication
// @Summary Delete ZaloApplication
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Social.ZaloApplication true "Delete ZaloApplication"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /zaloApplication/deleteZaloApplication [delete]
func (zaloApplicationApi *ZaloApplicationApi) DeleteZaloApplication(c *gin.Context) {
	var zaloApplication Social.ZaloApplication
	err := c.ShouldBindJSON(&zaloApplication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zaloApplicationService.DeleteZaloApplication(zaloApplication); err != nil {
		global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// DeleteZaloApplicationByIds Delete ZaloApplication
// @Tags ZaloApplication
// @Summary Delete ZaloApplication
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete ZaloApplication"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /zaloApplication/deleteZaloApplicationByIds [delete]
func (zaloApplicationApi *ZaloApplicationApi) DeleteZaloApplicationByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zaloApplicationService.DeleteZaloApplicationByIds(IDS); err != nil {
		global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// UpdateZaloApplication 更新ZaloApplication
// @Tags ZaloApplication
// @Summary 更新ZaloApplication
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Social.ZaloApplication true "更新ZaloApplication"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /zaloApplication/updateZaloApplication [put]
func (zaloApplicationApi *ZaloApplicationApi) UpdateZaloApplication(c *gin.Context) {
	var zaloApplication Social.ZaloApplication
	err := c.ShouldBindJSON(&zaloApplication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"ApplicationName": {utils.NotEmpty()},
		"ApplicationID":   {utils.NotEmpty()},
		"SecretKey":       {utils.NotEmpty()},
	}
	if err := utils.Verify(zaloApplication, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zaloApplicationService.UpdateZaloApplication(&zaloApplication); err != nil {
		global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
	} else {
		response.OkWithMessage("update completed", c)
	}
}

// FindZaloApplication 用idSearchZaloApplication
// @Tags ZaloApplication
// @Summary 用idSearchZaloApplication
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Social.ZaloApplication true "用idSearchZaloApplication"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /zaloApplication/findZaloApplication [get]
func (zaloApplicationApi *ZaloApplicationApi) FindZaloApplication(c *gin.Context) {
	var zaloApplication Social.ZaloApplication
	err := c.ShouldBindQuery(&zaloApplication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rezaloApplication, err := zaloApplicationService.GetZaloApplication(zaloApplication.ID); err != nil {
		global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		authorizationData := zalo.GetVerifierAndChallengeCode(rezaloApplication.CodeVerifier)
		codeChallenge := authorizationData["codeChallenge"]
		authorizationURL := fmt.Sprintf("https://oauth.zaloapp.com/v4/oa/permission?app_id=%s&redirect_uri=%s&code_challenge=%s", rezaloApplication.ApplicationID, rezaloApplication.CallbackURL, authorizationData["codeChallenge"])
		rezaloApplication.AuthorizationURL = authorizationURL
		rezaloApplication.CodeChallenge = fmt.Sprintf("%s", codeChallenge)
		response.OkWithData(gin.H{"rezaloApplication": rezaloApplication}, c)
	}
}

// GetZaloApplicationList 分页获取ZaloApplication列表
// @Tags ZaloApplication
// @Summary 分页获取ZaloApplication列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query SocialReq.ZaloApplicationSearch true "分页获取ZaloApplication列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /zaloApplication/getZaloApplicationList [get]
func (zaloApplicationApi *ZaloApplicationApi) GetZaloApplicationList(c *gin.Context) {
	var pageInfo SocialReq.ZaloApplicationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := zaloApplicationService.GetZaloApplicationInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("Fail!", zap.Error(err))
		response.FailWithMessage("Fail", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "Successful", c)
	}
}

// GetZaloApplicationActiveList
// @Tags ZaloApplicationActive
// @Summary GetZaloApplicationActiveList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query SocialReq.ZaloApplicationSearch true "分页获取ZaloApplication列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /zaloApplication/getZaloApplicationActiveList [get]
func (zaloApplicationApi *ZaloApplicationApi) GetZaloApplicationActiveList(c *gin.Context) {
	var pageInfo SocialReq.ZaloApplicationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := zaloApplicationService.GetZaloApplicationActiveInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("Fail!", zap.Error(err))
		response.FailWithMessage("Fail", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "Successful", c)
	}
}

// GetZaloNotificationTemplate
// @Tags ZaloApplication
// @Summary GetZaloNotificationTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Social.ZaloApplication true "用idSearchZaloApplication"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /zaloApplication/getZaloNotificationTemplate [get]
func (zaloApplicationApi *ZaloApplicationApi) GetZaloNotificationTemplate(c *gin.Context) {
	var zaloApplication Social.ZaloApplication
	err := c.ShouldBindQuery(&zaloApplication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var res map[string]interface{}
	_, err = zaloApplicationService.FetchAccessToken(&zaloApplication)
	if err != nil {
		global.GVA_LOG.Error("Fail!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	zaloNotificationServiceAPI.InitializeAPI()
	res, err = zaloNotificationServiceAPI.GetTemplateList(zaloApplication.AccessToken)
	if err != nil {
		global.GVA_LOG.Error("Fail!", zap.Error(err))
		fmt.Println(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	// fmt.Println(res)
	if res["message"] != "Success" {
		global.GVA_LOG.Error("Fail!", zap.Error(err))
		response.FailWithDetailed(res, "Zalo: Error", c)
		return
	}
	templates := res["data"].([]interface{})
	var results []map[string]interface{}
	for _, t := range templates {
		mia := t.(map[string]interface{})
		id := mia["templateId"].(float64)
		res, err = zaloNotificationServiceAPI.GetTemplateDetail(zaloApplication.AccessToken, id)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if res["message"] != "Success" {
			fmt.Println(err.Error())
			continue
		}

		data := res["data"].(map[string]interface{})
		var templateData []map[string]string
		listParams := data["listParams"].([]interface{})
		for _, p := range listParams {
			params := p.(map[string]interface{})
			templateData = append(templateData, map[string]string{
				"replaceKey": params["name"].(string),
				"replaceVal": "",
			})
		}

		results = append(results, map[string]interface{}{
			"name":       data["templateName"],
			"id":         id,
			"data":       templateData,
			"previewUrl": data["previewUrl"],
			"price":      data["price"],
			"timeout":    data["timeout"],
		})

	}

	response.OkWithData(results, c)

}
