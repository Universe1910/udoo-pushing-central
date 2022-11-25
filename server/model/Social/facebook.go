// 自动生成模板Facebook
package Social

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Facebook 结构体
type Facebook struct {
	global.GVA_MODEL
	AccessToken           string `json:"accessToken" form:"accessToken" gorm:"column:access_token;comment:;"`
	ExpiresIn             *int   `json:"expiresIn" form:"expiresIn" gorm:"column:expires_in;comment:;"`
	ReauthorizeRequiredIn string `json:"reauthorizeRequiredIn" form:"reauthorizeRequiredIn" gorm:"column:reauthorize_required_in;comment:;"`
	UserID                string `json:"userID" form:"userID" gorm:"column:user_id;comment:;"`
	AppName               string `json:"appName" form:"appName" gorm:"column:app_name;comment:;"`
	AppID                 string `json:"appId" form:"appId" gorm:"column:app_id;comment:;"`
	SignedRequest         string `json:"signedRequest" form:"signedRequest" gorm:"column:signed_request;comment:;"`
	Status                string `json:"status" form:"status" gorm:"column:status;comment:;"`
	SecretKey string `json:"secretKey" form:"secretKey" gorm:"column:secret_key;comment:;"`
}

// TableName Facebook 表名
func (Facebook) TableName() string {
	return "facebook"
}
