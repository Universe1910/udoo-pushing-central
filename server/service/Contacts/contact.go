package Contacts

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Contacts"
	ContactsReq "github.com/flipped-aurora/gin-vue-admin/server/model/Contacts/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	// uuid "github.com/satori/go.uuid"
)

type ContactService struct {
}

// CreateContact Create Contact记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactService *ContactService) CreateContact(contact *Contacts.Contact) (err error) {
	// contact.UUID = uuid.NewV4()
	err = global.GVA_DB.Create(&contact).Error
	return err
}

// DeleteContact Delete Contact记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactService *ContactService) DeleteContact(contact Contacts.Contact) (err error) {
	err = global.GVA_DB.Delete(&contact).Error
	return err
}

// DeleteContactByIds 批量Delete Contact记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactService *ContactService) DeleteContactByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Contacts.Contact{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateContact 更新Contact记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactService *ContactService) UpdateContact(contact Contacts.Contact) (err error) {
	err = global.GVA_DB.Save(&contact).Error
	return err
}

// GetContact 根据id获取Contact记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactService *ContactService) GetContact(id uint) (contact Contacts.Contact, err error) {
	err = global.GVA_DB.Where("id = ?", id).Preload("Tags").First(&contact).Error
	return
}

// GetContactInfoList 分页获取Contact记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactService *ContactService) GetContactInfoList(info ContactsReq.ContactSearch) (list []Contacts.Contact, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// Create db
	db := global.GVA_DB.Model(&Contacts.Contact{})
	var contacts []Contacts.Contact
	// 如果有条件搜索 下方会自动Create 搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Email != "" {
		db = db.Where("email LIKE ?", "%"+info.Email+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Tags").Find(&contacts).Error
	return contacts, total, err
}
