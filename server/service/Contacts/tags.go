package Contacts

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Contacts"
	ContactsReq "github.com/flipped-aurora/gin-vue-admin/server/model/Contacts/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TagsService struct {
}

// CreateTags Create Tags记录
// Author [piexlmax](https://github.com/piexlmax)
func (tagsService *TagsService) CreateTags(tags Contacts.Tags) (err error) {
	err = global.GVA_DB.Create(&tags).Error
	return err
}

// DeleteTags Delete Tags记录
// Author [piexlmax](https://github.com/piexlmax)
func (tagsService *TagsService) DeleteTags(tags Contacts.Tags) (err error) {
	err = global.GVA_DB.Delete(&tags).Error
	return err
}

// DeleteTagsByIds 批量Delete Tags记录
// Author [piexlmax](https://github.com/piexlmax)
func (tagsService *TagsService) DeleteTagsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Contacts.Tags{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTags 更新Tags记录
// Author [piexlmax](https://github.com/piexlmax)
func (tagsService *TagsService) UpdateTags(tags Contacts.Tags) (err error) {
	err = global.GVA_DB.Save(&tags).Error
	return err
}

// GetTags 根据id获取Tags记录
// Author [piexlmax](https://github.com/piexlmax)
func (tagsService *TagsService) GetTags(id uint) (tags Contacts.Tags, err error) {
	err = global.GVA_DB.Where("id = ?", id).Preload("CreatedObject").First(&tags).Error
	return
}

// GetTagsInfoList 分页获取Tags记录
// Author [piexlmax](https://github.com/piexlmax)
func (tagsService *TagsService) GetTagsInfoList(info ContactsReq.TagsSearch) (list []Contacts.Tags, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// Create db
	db := global.GVA_DB.Model(&Contacts.Tags{})
	var tagss []Contacts.Tags
	// 如果有条件搜索 下方会自动Create 搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("CreatedObject").Find(&tagss).Error
	return tagss, total, err
}
