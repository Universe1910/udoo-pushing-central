package Social

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Social"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    SocialReq "github.com/flipped-aurora/gin-vue-admin/server/model/Social/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type FacebookApi struct {
}

var facebookService = service.ServiceGroupApp.SocialServiceGroup.FacebookService


// CreateFacebook Create Facebook
// @Tags Facebook
// @Summary Create Facebook
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Social.Facebook true "Create Facebook"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /facebook/createFacebook [post]
func (facebookApi *FacebookApi) CreateFacebook(c *gin.Context) {
	var facebook Social.Facebook
	err := c.ShouldBindJSON(&facebook)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := facebookService.CreateFacebook(facebook); err != nil {
        global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create", c)
	} else {
		response.OkWithMessage("Successful creation", c)
	}
}

// DeleteFacebook Delete Facebook
// @Tags Facebook
// @Summary Delete Facebook
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Social.Facebook true "Delete Facebook"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /facebook/deleteFacebook [delete]
func (facebookApi *FacebookApi) DeleteFacebook(c *gin.Context) {
	var facebook Social.Facebook
	err := c.ShouldBindJSON(&facebook)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := facebookService.DeleteFacebook(facebook); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// DeleteFacebookByIds Delete Facebook
// @Tags Facebook
// @Summary Delete Facebook
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete Facebook"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /facebook/deleteFacebookByIds [delete]
func (facebookApi *FacebookApi) DeleteFacebookByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := facebookService.DeleteFacebookByIds(IDS); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// UpdateFacebook 更新Facebook
// @Tags Facebook
// @Summary 更新Facebook
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Social.Facebook true "更新Facebook"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /facebook/updateFacebook [put]
func (facebookApi *FacebookApi) UpdateFacebook(c *gin.Context) {
	var facebook Social.Facebook
	err := c.ShouldBindJSON(&facebook)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := facebookService.UpdateFacebook(facebook); err != nil {
        global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
	} else {
		response.OkWithMessage("update completed", c)
	}
}

// FindFacebook 用idSearchFacebook
// @Tags Facebook
// @Summary 用idSearchFacebook
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Social.Facebook true "用idSearchFacebook"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /facebook/findFacebook [get]
func (facebookApi *FacebookApi) FindFacebook(c *gin.Context) {
	var facebook Social.Facebook
	err := c.ShouldBindQuery(&facebook)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refacebook, err := facebookService.GetFacebook(facebook.ID); err != nil {
        global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		response.OkWithData(gin.H{"refacebook": refacebook}, c)
	}
}

// GetFacebookList 分页获取Facebook列表
// @Tags Facebook
// @Summary 分页获取Facebook列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query SocialReq.FacebookSearch true "分页获取Facebook列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /facebook/getFacebookList [get]
func (facebookApi *FacebookApi) GetFacebookList(c *gin.Context) {
	var pageInfo SocialReq.FacebookSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := facebookService.GetFacebookInfoList(pageInfo); err != nil {
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
