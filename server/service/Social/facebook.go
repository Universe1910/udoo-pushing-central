package Social

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Social"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    SocialReq "github.com/flipped-aurora/gin-vue-admin/server/model/Social/request"
)

type FacebookService struct {
}

// CreateFacebook Create Facebook记录
// Author [piexlmax](https://github.com/piexlmax)
func (facebookService *FacebookService) CreateFacebook(facebook Social.Facebook) (err error) {
	err = global.GVA_DB.Create(&facebook).Error
	return err
}

// DeleteFacebook Delete Facebook记录
// Author [piexlmax](https://github.com/piexlmax)
func (facebookService *FacebookService)DeleteFacebook(facebook Social.Facebook) (err error) {
	err = global.GVA_DB.Delete(&facebook).Error
	return err
}

// DeleteFacebookByIds 批量Delete Facebook记录
// Author [piexlmax](https://github.com/piexlmax)
func (facebookService *FacebookService)DeleteFacebookByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Social.Facebook{},"id in ?",ids.Ids).Error
	return err
}

// UpdateFacebook 更新Facebook记录
// Author [piexlmax](https://github.com/piexlmax)
func (facebookService *FacebookService)UpdateFacebook(facebook Social.Facebook) (err error) {
	err = global.GVA_DB.Save(&facebook).Error
	return err
}

// GetFacebook 根据id获取Facebook记录
// Author [piexlmax](https://github.com/piexlmax)
func (facebookService *FacebookService)GetFacebook(id uint) (facebook Social.Facebook, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&facebook).Error
	return
}

// GetFacebookInfoList 分页获取Facebook记录
// Author [piexlmax](https://github.com/piexlmax)
func (facebookService *FacebookService)GetFacebookInfoList(info SocialReq.FacebookSearch) (list []Social.Facebook, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // Create db
	db := global.GVA_DB.Model(&Social.Facebook{})
    var facebooks []Social.Facebook
    // 如果有条件搜索 下方会自动Create 搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&facebooks).Error
	return  facebooks, total, err
}
