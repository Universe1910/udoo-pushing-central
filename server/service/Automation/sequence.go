package Automation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Automation"
	AutomationReq "github.com/flipped-aurora/gin-vue-admin/server/model/Automation/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SequenceService struct {
}

// CreateSequence Create Sequence记录
// Author [piexlmax](https://github.com/piexlmax)
func (sequenceService *SequenceService) CreateSequence(sequence *Automation.Sequence) (err error) {
	err = global.GVA_DB.Create(&sequence).Error
	return err
}

// DeleteSequence Delete Sequence记录
// Author [piexlmax](https://github.com/piexlmax)
func (sequenceService *SequenceService) DeleteSequence(sequence Automation.Sequence) (err error) {
	err = global.GVA_DB.Delete(&sequence).Error
	return err
}

// DeleteSequenceByIds 批量Delete Sequence记录
// Author [piexlmax](https://github.com/piexlmax)
func (sequenceService *SequenceService) DeleteSequenceByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Automation.Sequence{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSequence 更新Sequence记录
// Author [piexlmax](https://github.com/piexlmax)
func (sequenceService *SequenceService) UpdateSequence(sequence *Automation.Sequence) (err error) {
	err = global.GVA_DB.Save(&sequence).Error
	return err
}

// GetSequence 根据id获取Sequence记录
// Author [piexlmax](https://github.com/piexlmax)
func (sequenceService *SequenceService) GetSequence(id uint) (sequence Automation.Sequence, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sequence).Error
	return
}

// GetSequenceByCampaign 根据id获取Sequence记录
// Author [piexlmax](https://github.com/piexlmax)
func (sequenceService *SequenceService) GetSequenceByCampaign(id uint) (list []Automation.Sequence, err error) {
	err = global.GVA_DB.Where("campaign_id = ?", id).Find(&list).Error
	return
}

// GetSequenceInfoList 分页获取Sequence记录
// Author [piexlmax](https://github.com/piexlmax)
func (sequenceService *SequenceService) GetSequenceInfoList(info AutomationReq.SequenceSearch) (list []Automation.Sequence, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// Create db
	db := global.GVA_DB.Model(&Automation.Sequence{})
	var sequences []Automation.Sequence
	// 如果有条件搜索 下方会自动Create 搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&sequences).Error
	return sequences, total, err
}
