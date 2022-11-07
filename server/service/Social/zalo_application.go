package Social

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Social"
	SocialReq "github.com/flipped-aurora/gin-vue-admin/server/model/Social/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ZaloApplicationService struct {
}

// CreateZaloApplication Create ZaloApplication记录
// Author [piexlmax](https://github.com/piexlmax)
func (zaloApplicationService *ZaloApplicationService) CreateZaloApplication(zaloApplication Social.ZaloApplication) (err error) {
	err = global.GVA_DB.Create(&zaloApplication).Error
	return err
}

// DeleteZaloApplication Delete ZaloApplication记录
// Author [piexlmax](https://github.com/piexlmax)
func (zaloApplicationService *ZaloApplicationService) DeleteZaloApplication(zaloApplication Social.ZaloApplication) (err error) {
	err = global.GVA_DB.Delete(&zaloApplication).Error
	return err
}

// DeleteZaloApplicationByIds 批量Delete ZaloApplication记录
// Author [piexlmax](https://github.com/piexlmax)
func (zaloApplicationService *ZaloApplicationService) DeleteZaloApplicationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Social.ZaloApplication{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateZaloApplication 更新ZaloApplication记录
// Author [piexlmax](https://github.com/piexlmax)
func (zaloApplicationService *ZaloApplicationService) UpdateZaloApplication(zaloApplication Social.ZaloApplication) (err error) {
	err = global.GVA_DB.Save(&zaloApplication).Error
	return err
}

// GetZaloApplication 根据id获取ZaloApplication记录
// Author [piexlmax](https://github.com/piexlmax)
func (zaloApplicationService *ZaloApplicationService) GetZaloApplication(id uint) (zaloApplication Social.ZaloApplication, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&zaloApplication).Error
	return
}

// GetZaloApplicationInfoList 分页获取ZaloApplication记录
// Author [piexlmax](https://github.com/piexlmax)
func (zaloApplicationService *ZaloApplicationService) GetZaloApplicationInfoList(info SocialReq.ZaloApplicationSearch) (list []Social.ZaloApplication, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// Create db
	db := global.GVA_DB.Model(&Social.ZaloApplication{})
	var zaloApplications []Social.ZaloApplication
	// 如果有条件搜索 下方会自动Create 搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&zaloApplications).Error
	return zaloApplications, total, err
}

// GetZaloApplicationInfoList 分页获取ZaloApplication记录
// Author [piexlmax](https://github.com/piexlmax)
func (zaloApplicationService *ZaloApplicationService) GetZaloApplicationActiveInfoList(info SocialReq.ZaloApplicationSearch) (list []Social.ZaloApplication, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// Create db
	db := global.GVA_DB.Model(&Social.ZaloApplication{})
	var zaloApplications []Social.ZaloApplication
	// 如果有条件搜索 下方会自动Create 搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ? AND status =?", info.StartCreatedAt, info.EndCreatedAt, 1)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&zaloApplications).Error
	return zaloApplications, total, err
}
