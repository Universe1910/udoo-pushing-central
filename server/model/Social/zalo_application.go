// 自动生成模板ZaloApplication
package Social

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ZaloApplication
type ZaloApplication struct {
	global.GVA_MODEL
	ApplicationName   string `json:"applicationName" form:"applicationName" gorm:"column:application_name;comment:;"`
	ApplicationID     string `json:"applicationID" form:"applicationID" gorm:"column:application_id;comment:;"`
	SecretKey         string `json:"secretKey" form:"secretKey" gorm:"column:secret_key;comment:;"`
	CodeVerifier      string `json:"codeVerifier" form:"codeVerifier" gorm:code_verifier" gorm:"column:code_verifier;comment:;"`
	Status            *bool  `json:"status" form:"status" gorm:"column:status;comment:;"`
	DefaultApp        *bool  `json:"defaultApp" form:"defaultApp" gorm:"column:default_app;comment:;"`
	CallbackURL       string `json:"callbackURL" form:"callbackURL" gorm:"column:call_back_url;comment:;"`
	AuthorizationURL  string `json:"authorizationURL"`
	CodeAuthorization string `json:"codeAuthorization" form:"codeAuthorization" gorm:"column:code_authorization;comment:;"`
	CodeChallenge     string `json:"codeChallenge" form:"codeChallenge" gorm:"column:code_challenge;comment:;"`
	OAID              string `json:"oaId" form:"oaId" gorm:"column:oa_id;comment:;"`
	AccessToken       string `json:"accessToken" form:"accessToken" gorm:"column:access_token;comment:;"`
	RefreshToken      string `json:"refreshToken" form:"refreshToken" gorm:"column:refresh_token;comment:;"`
	ExpiresIn         string `json:"expiresIn" form:"expiresIn" gorm:"column:expires_in;comment:;"`
}

// TableName ZaloApplication
func (ZaloApplication) TableName() string {
	return "zalo_appplication"
}
