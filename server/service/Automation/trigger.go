package Automation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Automation"
	AutomationReq "github.com/flipped-aurora/gin-vue-admin/server/model/Automation/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TriggerService struct {
}

// CreateTrigger Create Trigger记录
// Author [piexlmax](https://github.com/piexlmax)
func (triggerService *TriggerService) CreateTrigger(trigger *Automation.Trigger) (err error) {
	err = global.GVA_DB.Create(&trigger).Error
	return err
}

// DeleteTrigger Delete Trigger记录
// Author [piexlmax](https://github.com/piexlmax)
func (triggerService *TriggerService) DeleteTrigger(trigger Automation.Trigger) (err error) {
	err = global.GVA_DB.Delete(&trigger).Error
	return err
}

// DeleteTriggerByIds 批量Delete Trigger记录
// Author [piexlmax](https://github.com/piexlmax)
func (triggerService *TriggerService) DeleteTriggerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Automation.Trigger{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTrigger 更新Trigger记录
// Author [piexlmax](https://github.com/piexlmax)
func (triggerService *TriggerService) UpdateTrigger(trigger *Automation.Trigger) (err error) {
	err = global.GVA_DB.Save(&trigger).Error
	return err
}

// GetTrigger 根据id获取Trigger记录
// Author [piexlmax](https://github.com/piexlmax)
func (triggerService *TriggerService) GetTrigger(id uint) (trigger Automation.Trigger, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&trigger).Error
	return
}

// GetTriggerInfoList 分页获取Trigger记录
// Author [piexlmax](https://github.com/piexlmax)
func (triggerService *TriggerService) GetTriggerInfoList(info AutomationReq.TriggerSearch) (list []Automation.Trigger, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// Create db
	db := global.GVA_DB.Model(&Automation.Trigger{})
	var triggers []Automation.Trigger
	// 如果有条件搜索 下方会自动Create 搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("TriggerObject").Find(&triggers).Error
	return triggers, total, err
}
