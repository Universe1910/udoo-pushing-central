// 自动生成模板Tags
package Contacts

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// Tags 结构体
type Tags struct {
	global.GVA_MODEL
	Name          string           `json:"name" form:"name" gorm:"column:name;comment:;"`
	CreatedBy     *int             `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:;"`
	CreatedObject sysModel.SysUser `json:"createdObject" form:"createdObject" gorm:"foreignKey:CreatedBy"`
}

// TableName Tags 表名
func (Tags) TableName() string {
	return "tags"
}
