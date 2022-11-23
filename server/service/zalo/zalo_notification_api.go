package zalo

import (
	"fmt"
	"net/http"
	"strconv"
)

type ZaloNotificationServiceAPI struct {
	openApiUrl string
}

// const (
// 	ENABLE         int = 1
// 	PENDING_REVIEW     = 2
// 	REJECT             = 3
// 	DISABLE            = 4
// )

func (api *ZaloNotificationServiceAPI) InitializeAPI() {
	api.openApiUrl = "https://business.openapi.zalo.me"
}

func (api *ZaloNotificationServiceAPI) GetHeader(accessToken string) http.Header {
	// var header http.Header
	// header.Set("access_token", accessToken)
	header := http.Header{
		"access_token":  {accessToken},
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer Token"},
	}
	return header
}

func (api *ZaloNotificationServiceAPI) GetTemplateList(accessToken string) (res map[string]interface{}, err error) {
	builder := make(map[string]interface{})
	builder["endpoint"] = fmt.Sprintf("%s/template/all?offset=0&limit=100&status=%d", api.openApiUrl, 1)
	builder["method"] = "GET"
	builder["header"] = api.GetHeader(accessToken)
	res, err = ExecuteRequest(builder)
	return
}

func (api *ZaloNotificationServiceAPI) GetTemplateDetail(accessToken string, templateId float64) (res map[string]interface{}, err error) {
	builder := make(map[string]interface{})
	builder["endpoint"] = fmt.Sprintf("%s/template/info?template_id=%s", api.openApiUrl, strconv.FormatFloat(templateId, 'f', -1, 64))
	builder["method"] = "GET"
	builder["header"] = api.GetHeader(accessToken)
	res, err = ExecuteRequest(builder)
	return
}
