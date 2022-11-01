package Contacts

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Contacts"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ContactsReq "github.com/flipped-aurora/gin-vue-admin/server/model/Contacts/request"
)

type ContactTagsService struct {
}

// CreateContactTags Create ContactTags记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactTagsService *ContactTagsService) CreateContactTags(contactTags Contacts.ContactTags) (err error) {
	err = global.GVA_DB.Create(&contactTags).Error
	return err
}

// DeleteContactTags Delete ContactTags记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactTagsService *ContactTagsService)DeleteContactTags(contactTags Contacts.ContactTags) (err error) {
	err = global.GVA_DB.Delete(&contactTags).Error
	return err
}

// DeleteContactTagsByIds 批量Delete ContactTags记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactTagsService *ContactTagsService)DeleteContactTagsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Contacts.ContactTags{},"id in ?",ids.Ids).Error
	return err
}

// UpdateContactTags 更新ContactTags记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactTagsService *ContactTagsService)UpdateContactTags(contactTags Contacts.ContactTags) (err error) {
	err = global.GVA_DB.Save(&contactTags).Error
	return err
}

// GetContactTags 根据id获取ContactTags记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactTagsService *ContactTagsService)GetContactTags(id uint) (contactTags Contacts.ContactTags, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&contactTags).Error
	return
}

// GetContactTagsInfoList 分页获取ContactTags记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactTagsService *ContactTagsService)GetContactTagsInfoList(info ContactsReq.ContactTagsSearch) (list []Contacts.ContactTags, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // Create db
	db := global.GVA_DB.Model(&Contacts.ContactTags{})
    var contactTagss []Contacts.ContactTags
    // 如果有条件搜索 下方会自动Create 搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&contactTagss).Error
	return  contactTagss, total, err
}
