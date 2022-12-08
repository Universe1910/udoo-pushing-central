package Scenario

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Scenario"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ScenarioReq "github.com/flipped-aurora/gin-vue-admin/server/model/Scenario/request"
)

type ScenarioService struct {
}

// CreateScenario Create Scenario记录
// Author [piexlmax](https://github.com/piexlmax)
func (scenarioService *ScenarioService) CreateScenario(scenario Scenario.Scenario) (err error) {
	err = global.GVA_DB.Create(&scenario).Error
	return err
}

// DeleteScenario Delete Scenario记录
// Author [piexlmax](https://github.com/piexlmax)
func (scenarioService *ScenarioService)DeleteScenario(scenario Scenario.Scenario) (err error) {
	err = global.GVA_DB.Delete(&scenario).Error
	return err
}

// DeleteScenarioByIds 批量Delete Scenario记录
// Author [piexlmax](https://github.com/piexlmax)
func (scenarioService *ScenarioService)DeleteScenarioByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Scenario.Scenario{},"id in ?",ids.Ids).Error
	return err
}

// UpdateScenario 更新Scenario记录
// Author [piexlmax](https://github.com/piexlmax)
func (scenarioService *ScenarioService)UpdateScenario(scenario Scenario.Scenario) (err error) {
	err = global.GVA_DB.Save(&scenario).Error
	return err
}

// GetScenario 根据id获取Scenario记录
// Author [piexlmax](https://github.com/piexlmax)
func (scenarioService *ScenarioService)GetScenario(id uint) (scenario Scenario.Scenario, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&scenario).Error
	return
}

// GetScenarioInfoList 分页获取Scenario记录
// Author [piexlmax](https://github.com/piexlmax)
func (scenarioService *ScenarioService)GetScenarioInfoList(info ScenarioReq.ScenarioSearch) (list []Scenario.Scenario, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // Create db
	db := global.GVA_DB.Model(&Scenario.Scenario{})
    var scenarios []Scenario.Scenario
    // 如果有条件搜索 下方会自动Create 搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&scenarios).Error
	return  scenarios, total, err
}
