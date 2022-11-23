package Social

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Social"
	SocialReq "github.com/flipped-aurora/gin-vue-admin/server/model/Social/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service/zalo"
	"go.uber.org/zap"
	// "github.com/flipped-aurora/gin-vue-admin/server/service/zalo"
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
func (zaloApplicationService *ZaloApplicationService) UpdateZaloApplication(zaloApplication *Social.ZaloApplication) (err error) {
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

// GetZaloActiveApplication 根据id获取ZaloApplication记录
// Author [piexlmax](https://github.com/piexlmax)
func (zaloApplicationService *ZaloApplicationService) GetZaloActiveApplication() (zaloApplication Social.ZaloApplication, err error) {
	err = global.GVA_DB.Where("default_app = ? AND status = ?", "TRUE", "TRUE").First(&zaloApplication).Error
	return
}

func (zaloApplicationService *ZaloApplicationService) FetchAccessToken(zaloApplication *Social.ZaloApplication) (res map[string]interface{}, err error) {
	var zoApi zalo.ZaloOfficalAPI
	zoApi.InitializeAPI()
	res, err = zoApi.GetAccessTokenAPI(zaloApplication.SecretKey, zaloApplication.CodeAuthorization, zaloApplication.ApplicationID, zaloApplication.CodeVerifier)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if res["error_name"] == "Authorized code expired" {
		res, err = zoApi.RefrestTokenAPI(zaloApplication.SecretKey, zaloApplication.RefreshToken, zaloApplication.ApplicationID)
		fmt.Println(res)
	}
	if res["error_name"] != nil {
		return nil, errors.New(fmt.Sprintf("Zalo: %s", res["error_name"].(string)))
	}
	zaloApplication.AccessToken = res["access_token"].(string)
	zaloApplication.RefreshToken = res["refresh_token"].(string)
	zaloApplication.ExpiresIn = res["expires_in"].(string)
	if err := zaloApplicationService.UpdateZaloApplication(zaloApplication); err != nil {
		global.GVA_LOG.Error("Query Failed", zap.Error(err))
		return res, err
	}
	return
}
