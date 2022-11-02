// 自动生成模板ContactTags
package Contacts

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// ContactTags 结构体
type ContactTags struct {
      global.GVA_MODEL
      ContactID  *int `json:"contactID" form:"contactID" gorm:"column:contact_id;comment:;"`
      TagID  *int `json:"tagID" form:"tagID" gorm:"column:tag_id;comment:;"`
}


// TableName ContactTags 表名
func (ContactTags) TableName() string {
  return "contact_tags"
}

