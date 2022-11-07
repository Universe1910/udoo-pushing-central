package Automation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Automation"
	AutomationReq "github.com/flipped-aurora/gin-vue-admin/server/model/Automation/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CampaignService struct {
}

// CreateCampaign Create Campaign记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignService *CampaignService) CreateCampaign(campaign Automation.Campaign) (err error) {
	err = global.GVA_DB.Create(&campaign).Error
	return err
}

// DeleteCampaign Delete Campaign记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignService *CampaignService) DeleteCampaign(campaign Automation.Campaign) (err error) {
	err = global.GVA_DB.Delete(&campaign).Error
	return err
}

// DeleteCampaignByIds 批量Delete Campaign记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignService *CampaignService) DeleteCampaignByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Automation.Campaign{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCampaign 更新Campaign记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignService *CampaignService) UpdateCampaign(campaign Automation.Campaign) (err error) {
	err = global.GVA_DB.Save(&campaign).Error
	return err
}

// GetCampaign 根据id获取Campaign记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignService *CampaignService) GetCampaign(id uint) (campaign Automation.Campaign, err error) {
	err = global.GVA_DB.Where("id = ?", id).Preload("Contacts").First(&campaign).Error
	return
}

// GetCampaignInfoList 分页获取Campaign记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignService *CampaignService) GetCampaignInfoList(info AutomationReq.CampaignSearch) (list []Automation.Campaign, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// Create db
	db := global.GVA_DB.Model(&Automation.Campaign{})
	var campaigns []Automation.Campaign
	// 如果有条件搜索 下方会自动Create 搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Contacts").Find(&campaigns).Error
	return campaigns, total, err
}
