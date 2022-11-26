package Automation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Automation"
	AutomationReq "github.com/flipped-aurora/gin-vue-admin/server/model/Automation/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CampaignLogService struct {
}

// CreateCampaignLog Create CampaignLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignLogService *CampaignLogService) CreateCampaignLog(campaignLog Automation.CampaignLog) (err error) {
	err = global.GVA_DB.Create(&campaignLog).Error
	return err
}

// DeleteCampaignLog Delete CampaignLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignLogService *CampaignLogService) DeleteCampaignLog(campaignLog Automation.CampaignLog) (err error) {
	err = global.GVA_DB.Delete(&campaignLog).Error
	return err
}

// DeleteCampaignLogByIds 批量Delete CampaignLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignLogService *CampaignLogService) DeleteCampaignLogByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Automation.CampaignLog{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCampaignLog 更新CampaignLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignLogService *CampaignLogService) UpdateCampaignLog(campaignLog Automation.CampaignLog) (err error) {
	err = global.GVA_DB.Save(&campaignLog).Error
	return err
}

// GetCampaignLog 根据id获取CampaignLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignLogService *CampaignLogService) GetCampaignLog(id uint) (campaignLog Automation.CampaignLog, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&campaignLog).Error
	return
}

// GetCampaignLogByCampaign 根据id获取CampaignLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignLogService *CampaignLogService) GetCampaignLogByCampaign(id uint) ([]Automation.CampaignLog, error) {
	var campaignLogs []Automation.CampaignLog
	db := global.GVA_DB.Model(&Automation.CampaignLog{})
	db = db.Where("campaign_id = ?", id)
	err := db.Find(&campaignLogs).Error
	return campaignLogs, err
}

// GetCampaignLogInfoList 分页获取CampaignLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignLogService *CampaignLogService) GetCampaignLogInfoList(info AutomationReq.CampaignLogSearch) (list []Automation.CampaignLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// Create db
	db := global.GVA_DB.Model(&Automation.CampaignLog{})
	var campaignLogs []Automation.CampaignLog
	// 如果有条件搜索 下方会自动Create 搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&campaignLogs).Error
	return campaignLogs, total, err
}
