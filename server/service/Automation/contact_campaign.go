package Automation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Automation"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    AutomationReq "github.com/flipped-aurora/gin-vue-admin/server/model/Automation/request"
)

type ContactCampaignService struct {
}

// CreateContactCampaign Create ContactCampaign记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactCampaignService *ContactCampaignService) CreateContactCampaign(contactCampaign Automation.ContactCampaign) (err error) {
	err = global.GVA_DB.Create(&contactCampaign).Error
	return err
}

// DeleteContactCampaign Delete ContactCampaign记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactCampaignService *ContactCampaignService)DeleteContactCampaign(contactCampaign Automation.ContactCampaign) (err error) {
	err = global.GVA_DB.Delete(&contactCampaign).Error
	return err
}

// DeleteContactCampaignByIds 批量Delete ContactCampaign记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactCampaignService *ContactCampaignService)DeleteContactCampaignByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Automation.ContactCampaign{},"id in ?",ids.Ids).Error
	return err
}

// UpdateContactCampaign 更新ContactCampaign记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactCampaignService *ContactCampaignService)UpdateContactCampaign(contactCampaign Automation.ContactCampaign) (err error) {
	err = global.GVA_DB.Save(&contactCampaign).Error
	return err
}

// GetContactCampaign 根据id获取ContactCampaign记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactCampaignService *ContactCampaignService)GetContactCampaign(id uint) (contactCampaign Automation.ContactCampaign, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&contactCampaign).Error
	return
}

// GetContactCampaignInfoList 分页获取ContactCampaign记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactCampaignService *ContactCampaignService)GetContactCampaignInfoList(info AutomationReq.ContactCampaignSearch) (list []Automation.ContactCampaign, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // Create db
	db := global.GVA_DB.Model(&Automation.ContactCampaign{})
    var contactCampaigns []Automation.ContactCampaign
    // 如果有条件搜索 下方会自动Create 搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&contactCampaigns).Error
	return  contactCampaigns, total, err
}
