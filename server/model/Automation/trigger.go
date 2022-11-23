// 自动生成模板Trigger
package Automation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Trigger 结构体
type Trigger struct {
	global.GVA_MODEL
	Data       string `json:"data" form:"data" gorm:"column:data;comment:;"`
	Flow       string `json:"flow" form:"flow" gorm:"column:flow; comment:;"`
	CampaignId *int   `json:"campaignId" form:"campaignId" gorm:"column:campaign_id;comment:;"`
	Status     string `json:"status" form:"status" gorm:"column:status;comment:;"`
}

// TableName Trigger 表名
func (Trigger) TableName() string {
	return "trigger"
}
