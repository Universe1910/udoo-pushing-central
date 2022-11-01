// 自动生成模板EmailTemplate
package EmailMarketing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// EmailTemplate 结构体
type EmailTemplate struct {
	global.GVA_MODEL
	Name          string           `json:"name" form:"name" gorm:"column:name;comment:;"`
	Content       string           `json:"content" form:"content" gorm:"column:content;comment:Content HTML;"`
	Subject       string           `json:"subject" form:"subject" gorm:"column:subject;comment:Subject;"`
	CreatedBy     *int             `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:User create email;"`
	CreatedObject sysModel.SysUser `json:"createdObject" gorm:"foreignKey:CreatedBy"`
}

// TableName EmailTemplate 表名
func (EmailTemplate) TableName() string {
	return "email_template"
}
