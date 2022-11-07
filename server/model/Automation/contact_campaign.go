// 自动生成模板ContactCampaign
package Automation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// ContactCampaign 结构体
type ContactCampaign struct {
      global.GVA_MODEL
      ContactId  *int `json:"contactId" form:"contactId" gorm:"column:contact_id;comment:;"`
      CampaignId  *int `json:"campaignId" form:"campaignId" gorm:"column:campaign_id;comment:;"`
}


// TableName ContactCampaign 表名
func (ContactCampaign) TableName() string {
  return "contact_campaign"
}

