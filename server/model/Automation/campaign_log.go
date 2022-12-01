// 自动生成模板CampaignLog
package Automation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CampaignLog 结构体
type CampaignLog struct {
	global.GVA_MODEL
	CampaignID uint   `json:"campaignID" form:"campaignID" gorm:"column:campaign_id;comment:;"`
	Action     string `json:"action" form:"action" gorm:"column:action;comment:;"`
	ContactID  uint   `json:"contactID" form:"contactID" gorm:"column:contact_id;comment:;"`
	Message    string `json:"message" form:"message" gorm:"column:message;comment:;"`
	Type       string `json:"type" form:"type" gorm:"column:type;comment:;"`
	Source       string `json:"source" form:"source" gorm:"column:source;comment:;"`
}

// TableName CampaignLog 表名
func (CampaignLog) TableName() string {
	return "campaign_log"
}
