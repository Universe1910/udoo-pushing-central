package EmailMarketing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EmailMarketing"
	EmailMarketingReq "github.com/flipped-aurora/gin-vue-admin/server/model/EmailMarketing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type EmailTemplateService struct {
}

// CreateEmailTemplate Create EmailTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (email_templateService *EmailTemplateService) CreateEmailTemplate(email_template EmailMarketing.EmailTemplate) (err error) {
	err = global.GVA_DB.Create(&email_template).Error
	return err
}

// DeleteEmailTemplate Delete EmailTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (email_templateService *EmailTemplateService) DeleteEmailTemplate(email_template EmailMarketing.EmailTemplate) (err error) {
	err = global.GVA_DB.Delete(&email_template).Error
	return err
}

// DeleteEmailTemplateByIds 批量Delete EmailTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (email_templateService *EmailTemplateService) DeleteEmailTemplateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]EmailMarketing.EmailTemplate{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateEmailTemplate 更新EmailTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (email_templateService *EmailTemplateService) UpdateEmailTemplate(email_template EmailMarketing.EmailTemplate) (err error) {
	err = global.GVA_DB.Save(&email_template).Error
	return err
}

// GetEmailTemplate 根据id获取EmailTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (email_templateService *EmailTemplateService) GetEmailTemplate(id uint) (email_template EmailMarketing.EmailTemplate, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&email_template).Error
	return
}

// GetEmailTemplateInfoList 分页获取EmailTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (email_templateService *EmailTemplateService) GetEmailTemplateInfoList(info EmailMarketingReq.EmailTemplateSearch) (list []EmailMarketing.EmailTemplate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// Create db
	db := global.GVA_DB.Model(&EmailMarketing.EmailTemplate{})
	var email_templates []EmailMarketing.EmailTemplate
	// 如果有条件搜索 下方会自动Create 搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&email_templates).Error
	return email_templates, total, err
}
