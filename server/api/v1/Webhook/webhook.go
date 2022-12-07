package Webhook

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/flipped-aurora/gin-vue-admin/server/model/FacebookMessgener"
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
	// var messgenerResponse map[string]interface{}
	// err := c.ShouldBindJSON(&messgenerResponse)
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// }
	// fmt.Println(messgenerResponse)
	var callback FacebookMessgener.Callback
	err := c.ShouldBindJSON(&callback)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	if callback.Object == "page" {
		for _, entry := range callback.Entry {
			for _, event := range entry.Messaging {
				if !reflect.DeepEqual(event.Message, FacebookMessgener.Message{}) && event.Message.Text != "" {
					// ProcessMessage(event)
					fmt.Println("event:")
					fmt.Println(event)
				}
			}
		}
		// w.WriteHeader(200)
		// w.Write([]byte("Got your message"))
		response.OkWithMessage("Got your message", c)
	} else {
		// w.WriteHeader(404)
		// w.Write([]byte("Message not supported"))
		response.FailWithMessage("Message not supported", c)
	}
}

// Debug Campaign
// @Tags Campaign
// @Summary Debug Campaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully debug"}"
// @Router /webhook/facebookMessgenerHook [get]
func (webhookApi *WebhookApi) VerifierFacebookMessgenerWebHook(c *gin.Context) {
	fmt.Print("VerifierFacebookMessgenerWebHook")
	// var query map[string]interface{}
	// err := c.ShouldBindQuery(&query)
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	challenge := c.Query("hub.challenge")
	mode := c.Query("hub.mode")
	token := c.Query("hub.verify_token")
	verifyToken := "gvhe8Yg@C9YG9pqR"
	if mode != "" && token == verifyToken {
		// w.WriteHeader(200)
		// w.Write([]byte(challenge))
		// response.OkWithData(challenge, c)
		// c.Writer()
		c.String(http.StatusOK, challenge)
	} else {
		// w.WriteHeader(404)
		// w.Write([]byte("Error, wrong validation token"))
		response.FailWithMessage("Error, wrong validation token", c)
	}

	// challenge := r.URL.Query().Get("hub.challenge")
	// mode := r.URL.Query().Get("hub.mode")
	// token := r.URL.Query().Get("hub.verify_token")
	// if mode != "" && token == os.Getenv("VERIFY_TOKEN") {
	// 	w.WriteHeader(200)
	// 	w.Write([]byte(challenge))
	// } else {
	// 	w.WriteHeader(404)
	// 	w.Write([]byte("Error, wrong validation token"))
	// }
}
