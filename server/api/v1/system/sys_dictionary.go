package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DictionaryApi struct{}

// CreateSysDictionary
// @Tags      SysDictionary
// @Summary   Create SysDictionary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDictionary           true  "SysDictionary模型"
// @Success   200   {object}  response.Response{msg=string}  "Create SysDictionary"
// @Router    /sysDictionary/createSysDictionary [post]
func (s *DictionaryApi) CreateSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryService.CreateSysDictionary(dictionary)
	if err != nil {
		global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create", c)
		return
	}
	response.OkWithMessage("Successful creation", c)
}

// DeleteSysDictionary
// @Tags      SysDictionary
// @Summary   Delete SysDictionary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDictionary           true  "SysDictionary模型"
// @Success   200   {object}  response.Response{msg=string}  "Delete SysDictionary"
// @Router    /sysDictionary/deleteSysDictionary [delete]
func (s *DictionaryApi) DeleteSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryService.DeleteSysDictionary(dictionary)
	if err != nil {
		global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
		return
	}
	response.OkWithMessage("successfully deleted", c)
}

// UpdateSysDictionary
// @Tags      SysDictionary
// @Summary   更新SysDictionary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDictionary           true  "SysDictionary模型"
// @Success   200   {object}  response.Response{msg=string}  "更新SysDictionary"
// @Router    /sysDictionary/updateSysDictionary [put]
func (s *DictionaryApi) UpdateSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryService.UpdateSysDictionary(&dictionary)
	if err != nil {
		global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
		return
	}
	response.OkWithMessage("update completed", c)
}

// FindSysDictionary
// @Tags      SysDictionary
// @Summary   用id查询SysDictionary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     system.SysDictionary                                       true  "ID或字典英名"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "用id查询SysDictionary"
// @Router    /sysDictionary/findSysDictionary [get]
func (s *DictionaryApi) FindSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	err := c.ShouldBindQuery(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sysDictionary, err := dictionaryService.GetSysDictionary(dictionary.Type, dictionary.ID, dictionary.Status)
	if err != nil {
		global.GVA_LOG.Error("字典未Create 或未开启!", zap.Error(err))
		response.FailWithMessage("字典未Create 或未开启", c)
		return
	}
	response.OkWithDetailed(gin.H{"resysDictionary": sysDictionary}, "Find Successfully", c)
}

// GetSysDictionaryList
// @Tags      SysDictionary
// @Summary   分页获取SysDictionary列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.SysDictionarySearch                             true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取SysDictionary列表,返回包括列表,总数,页码,每页数量"
// @Router    /sysDictionary/getSysDictionaryList [get]
func (s *DictionaryApi) GetSysDictionaryList(c *gin.Context) {
	var pageInfo request.SysDictionarySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := dictionaryService.GetSysDictionaryInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Fail!", zap.Error(err))
		response.FailWithMessage("Fail", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Successful", c)
}
