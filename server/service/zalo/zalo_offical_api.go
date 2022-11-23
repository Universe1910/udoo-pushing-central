package zalo

import (
	"fmt"
	"net/http"
)

type ZaloOfficalAPI struct {
	apiList      map[string]interface{}
	oauthUrl     string
	openApiUrl   string
	openVersion  string
	oauthVersion string
}

func (api *ZaloOfficalAPI) InitializeAPI() {
	api.openApiUrl = ""
	api.oauthVersion = "v4"
	api.openVersion = "v2.0"
	api.oauthUrl = "https://oauth.zaloapp.com"
	// api.apiList = make(map[string]interface{})
	// api.apiList["getAccessTokenAPI"] = api.getAccessTokenAPI()
	// api.apiList["refreshAccessToken"] = api.refrestTokenAPI()
}

func (api *ZaloOfficalAPI) GetHeader(secretKey string) http.Header {
	// var header http.Header
	// header.Add("secret_key", secretKey)
	// header.Add("Content-Type", "application/json")
	// header.Add("Accept", "*/*")
	// return header
	header := http.Header{
		"secret_key":    {secretKey},
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer Token"},
	}
	return header
}

func (api *ZaloOfficalAPI) GetAccessTokenAPI(secretKey, code, appId, codeVerifier string) (res map[string]interface{}, err error) {
	builder := make(map[string]interface{})
	// builder["endpoint"] = "https://oauth.zaloapp.com/v4/oa/access_token"
	// https://oauth.zaloapp.com/v4/oa/access_token&app_id&grant_type=authorization_code&code_verifier=MEYuxmQt5Xw9adAXWTyyCceene52XywZ5tHPW8TNDMX
	builder["endpoint"] = fmt.Sprintf("%s/%s/oa/access_token?code=%s&app_id=%s&grant_type=authorization_code&code_verifier=%s", api.oauthUrl, api.oauthVersion, code, appId, codeVerifier)
	builder["method"] = "POST"
	builder["header"] = api.GetHeader(secretKey)
	res, err = ExecuteRequest(builder)
	return
}

func (api *ZaloOfficalAPI) RefrestTokenAPI(secretKey, refreshToken, appId string) (res map[string]interface{}, err error) {
	builder := make(map[string]interface{})
	// builder["endpoint"] = fmt.Sprintf("%s/%s/oa/access_token", api.oauthUrl, api.oauthVersion)
	// builder["method"] = "POST"
	builder["endpoint"] = fmt.Sprintf("%s/%s/oa/access_token?refresh_token=%s&app_id=%s&grant_type=refresh_token", api.oauthUrl, api.oauthVersion, refreshToken, appId)
	builder["method"] = "POST"
	builder["header"] = api.GetHeader(secretKey)
	res, err = ExecuteRequest(builder)
	return
}

func (api *ZaloOfficalAPI) getFollowers() map[string]interface{} {
	builder := make(map[string]interface{})
	// builder["endpoint"] = "https://openapi.zalo.me/v2.0/oa/getfollowers"
	builder["endpoint"] = fmt.Sprintf("%s/%s/oa/getfollowers", api.openApiUrl, api.openVersion)
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
