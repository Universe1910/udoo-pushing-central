package Social

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Social"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    SocialReq "github.com/flipped-aurora/gin-vue-admin/server/model/Social/request"
)

type ZaloOfficalAccountService struct {
}

// CreateZaloOfficalAccount Create ZaloOfficalAccount记录
// Author [piexlmax](https://github.com/piexlmax)
func (zaloOfficalAccountService *ZaloOfficalAccountService) CreateZaloOfficalAccount(zaloOfficalAccount Social.ZaloOfficalAccount) (err error) {
	err = global.GVA_DB.Create(&zaloOfficalAccount).Error
	return err
}

// DeleteZaloOfficalAccount Delete ZaloOfficalAccount记录
// Author [piexlmax](https://github.com/piexlmax)
func (zaloOfficalAccountService *ZaloOfficalAccountService)DeleteZaloOfficalAccount(zaloOfficalAccount Social.ZaloOfficalAccount) (err error) {
	err = global.GVA_DB.Delete(&zaloOfficalAccount).Error
	return err
}

// DeleteZaloOfficalAccountByIds 批量Delete ZaloOfficalAccount记录
// Author [piexlmax](https://github.com/piexlmax)
func (zaloOfficalAccountService *ZaloOfficalAccountService)DeleteZaloOfficalAccountByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Social.ZaloOfficalAccount{},"id in ?",ids.Ids).Error
	return err
}

// UpdateZaloOfficalAccount 更新ZaloOfficalAccount记录
// Author [piexlmax](https://github.com/piexlmax)
func (zaloOfficalAccountService *ZaloOfficalAccountService)UpdateZaloOfficalAccount(zaloOfficalAccount Social.ZaloOfficalAccount) (err error) {
	err = global.GVA_DB.Save(&zaloOfficalAccount).Error
	return err
}

// GetZaloOfficalAccount 根据id获取ZaloOfficalAccount记录
// Author [piexlmax](https://github.com/piexlmax)
func (zaloOfficalAccountService *ZaloOfficalAccountService)GetZaloOfficalAccount(id uint) (zaloOfficalAccount Social.ZaloOfficalAccount, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&zaloOfficalAccount).Error
	return
}

// GetZaloOfficalAccountInfoList 分页获取ZaloOfficalAccount记录
// Author [piexlmax](https://github.com/piexlmax)
func (zaloOfficalAccountService *ZaloOfficalAccountService)GetZaloOfficalAccountInfoList(info SocialReq.ZaloOfficalAccountSearch) (list []Social.ZaloOfficalAccount, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // Create db
	db := global.GVA_DB.Model(&Social.ZaloOfficalAccount{})
    var zaloOfficalAccounts []Social.ZaloOfficalAccount
    // 如果有条件搜索 下方会自动Create 搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&zaloOfficalAccounts).Error
	return  zaloOfficalAccounts, total, err
}
