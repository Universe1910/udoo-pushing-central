// 自动生成模板Scenario
package Scenario

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Scenario 结构体
type Scenario struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
      CreatedBy  *int `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:;"`
      Data  string `json:"data" form:"data" gorm:"column:data;comment:;"`
}


// TableName Scenario 表名
func (Scenario) TableName() string {
  return "scenario"
}

