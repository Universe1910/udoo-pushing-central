package zalo

import (
	"fmt"
)

type ZaloAPI struct {
	apiList      map[string]interface{}
	oauthUrl     string
	openApiUrl   string
	openVersion  string
	oauthVersion string
}

func (api *ZaloAPI) initializeAPILists() {
	api.openApiUrl = ""
	api.oauthVersion = "v4"
	api.openVersion = "v2.0"
	api.oauthUrl = "https://oauth.zaloapp.com "
	api.apiList = make(map[string]interface{})
	api.apiList["builder"] = api.getAccessTokenAPI()
	api.apiList["refreshAccessToken"] = api.refrestTokenAPI()
}

func (api *ZaloAPI) getAccessTokenAPI() map[string]interface{} {
	builder := make(map[string]interface{})
	// builder["url"] = "https://oauth.zaloapp.com/v4/oa/access_token"
	builder["url"] = fmt.Sprintf("%s/%s/oa/access_token", api.oauthUrl, api.oauthVersion)
	builder["method"] = "POST"
	builder["header"] = map[string]interface{}{"secret_key": ""}
	builder["body"] = map[string]interface{}{
		"code":          "",
		"app_id":        "",
		"grant_type":    "authorization_code",
		"code_verifier": "",
	}
	builder["dataType"] = "urlencode"
	builder["response"] = []string{"access_token", "refresh_token", "expires_in"}
	return builder
}

func (api *ZaloAPI) refrestTokenAPI() map[string]interface{} {
	builder := make(map[string]interface{})
	builder["url"] = fmt.Sprintf("%s/%s/oa/access_token", api.oauthUrl, api.oauthVersion)
	builder["method"] = "POST"
	builder["header"] = map[string]interface{}{"secret_key": ""}
	builder["body"] = map[string]interface{}{
		"refresh_token": "",
		"app_id":        "",
		"grant_type":    "refresh_token",
		"code_verifier": "",
	}
	builder["dataType"] = "urlencode"
	builder["response"] = []string{"access_token", "refresh_token", "expires_in"}
	return builder
}

func (api *ZaloAPI) getFollowers() map[string]interface{} {
	builder := make(map[string]interface{})
	// builder["url"] = "https://openapi.zalo.me/v2.0/oa/getfollowers"
	builder["url"] = fmt.Sprintf("%s/%s/oa/getfollowers", api.openApiUrl, api.openVersion)
	builder["method"] = "GET"
	builder["header"] = map[string]interface{}{
		"access_token": "",
	}
	builder["data"] = map[string]interface{}{
		"offset":   0,
		"count":    5,
		"tag_name": "",
	}
	return builder
}
