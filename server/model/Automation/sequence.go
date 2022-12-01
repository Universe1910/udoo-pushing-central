// 自动生成模板Sequence
package Automation

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Sequence 结构体
type Sequence struct {
	global.GVA_MODEL
	CampaignId uint       `json:"campaignId" form:"campaignId" gorm:"column:campaign_id;comment:;"`
	Order      int        `json:"order" form:"order"  gorm:"column:order;comment:;"`
	Date       *time.Time `json:"date" form:"date" gorm:"column:date;comment:;"`
	ActionName string     `json:"actionName" form:"actionName" gorm:"column:action_name;comment:;"`
	Data       string     `json:"data" form:"data" gorm:"column:data;comment:;"`
}

// TableName Sequence 表名
func (Sequence) TableName() string {
	return "sequence"
}
