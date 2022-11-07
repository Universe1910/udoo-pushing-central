// 自动生成模板ZaloOfficalAccount
package Social

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// ZaloOfficalAccount 结构体
type ZaloOfficalAccount struct {
      global.GVA_MODEL
      Oaid  string `json:"oaid" form:"oaid" gorm:"column:oaid;comment:;"`
      OfficialAccount  string `json:"officialAccount" form:"officialAccount" gorm:"column:official_account;comment:;"`
}


// TableName ZaloOfficalAccount 表名
func (ZaloOfficalAccount) TableName() string {
  return "zalo_offical_account"
}

