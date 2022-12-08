package Scenario

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Scenario"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ScenarioReq "github.com/flipped-aurora/gin-vue-admin/server/model/Scenario/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type ScenarioApi struct {
}

var scenarioService = service.ServiceGroupApp.ScenarioServiceGroup.ScenarioService


// CreateScenario Create Scenario
// @Tags Scenario
// @Summary Create Scenario
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Scenario.Scenario true "Create Scenario"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /scenario/createScenario [post]
func (scenarioApi *ScenarioApi) CreateScenario(c *gin.Context) {
	var scenario Scenario.Scenario
	err := c.ShouldBindJSON(&scenario)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "CreatedBy":{utils.NotEmpty()},
    }
	if err := utils.Verify(scenario, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := scenarioService.CreateScenario(scenario); err != nil {
        global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create", c)
	} else {
		response.OkWithMessage("Successful creation", c)
	}
}

// DeleteScenario Delete Scenario
// @Tags Scenario
// @Summary Delete Scenario
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Scenario.Scenario true "Delete Scenario"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /scenario/deleteScenario [delete]
func (scenarioApi *ScenarioApi) DeleteScenario(c *gin.Context) {
	var scenario Scenario.Scenario
	err := c.ShouldBindJSON(&scenario)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := scenarioService.DeleteScenario(scenario); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// DeleteScenarioByIds Delete Scenario
// @Tags Scenario
// @Summary Delete Scenario
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete Scenario"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /scenario/deleteScenarioByIds [delete]
func (scenarioApi *ScenarioApi) DeleteScenarioByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := scenarioService.DeleteScenarioByIds(IDS); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// UpdateScenario 更新Scenario
// @Tags Scenario
// @Summary 更新Scenario
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Scenario.Scenario true "更新Scenario"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /scenario/updateScenario [put]
func (scenarioApi *ScenarioApi) UpdateScenario(c *gin.Context) {
	var scenario Scenario.Scenario
	err := c.ShouldBindJSON(&scenario)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "CreatedBy":{utils.NotEmpty()},
      }
    if err := utils.Verify(scenario, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := scenarioService.UpdateScenario(scenario); err != nil {
        global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
	} else {
		response.OkWithMessage("update completed", c)
	}
}

// FindScenario 用idSearchScenario
// @Tags Scenario
// @Summary 用idSearchScenario
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Scenario.Scenario true "用idSearchScenario"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /scenario/findScenario [get]
func (scenarioApi *ScenarioApi) FindScenario(c *gin.Context) {
	var scenario Scenario.Scenario
	err := c.ShouldBindQuery(&scenario)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rescenario, err := scenarioService.GetScenario(scenario.ID); err != nil {
        global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		response.OkWithData(gin.H{"rescenario": rescenario}, c)
	}
}

// GetScenarioList 分页获取Scenario列表
// @Tags Scenario
// @Summary 分页获取Scenario列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ScenarioReq.ScenarioSearch true "分页获取Scenario列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /scenario/getScenarioList [get]
func (scenarioApi *ScenarioApi) GetScenarioList(c *gin.Context) {
	var pageInfo ScenarioReq.ScenarioSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := scenarioService.GetScenarioInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("Fail!", zap.Error(err))
        response.FailWithMessage("Fail", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "Successful", c)
    }
}
