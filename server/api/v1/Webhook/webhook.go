package Webhook

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Automation"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type WebhookApi struct {
}

// Debug Campaign
// @Tags Campaign
// @Summary Debug Campaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Automation.Campaign true "用idSearchCampaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully debug"}"
// @Router /webhook/facebookMessgenerHook [post]
func (webhookApi *WebhookApi) FacebookMessgenerWebHook(c *gin.Context) {
	var campaign Automation.Campaign
	err := c.ShouldBindJSON(&campaign)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	// if res, err := campaignService.DebugCampaign(campaign.ID); err != nil {
	// 	global.GVA_LOG.Error("failed to debug!", zap.Error(err))
	// 	response.FailWithMessage(err.Error(), c)
	// } else {
	// 	if res["error"] == 0 {
	// 		global.GVA_LOG.Error(res["message"].(string), zap.Error(err))
	// 		response.FailWithMessage(res["message"].(string), c)
	// 		return
	// 	}
	// 	response.OkWithMessage("successfully debuged", c)
	// }
}
