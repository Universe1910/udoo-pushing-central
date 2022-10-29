package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityApi struct{}

// CreateAuthority
// @Tags      Authority
// @Summary   Create 角色
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "Create 角色,返回包括系统角色详情"
// @Router    /authority/createAuthority [post]
func (a *AuthorityApi) CreateAuthority(c *gin.Context) {
	var authority system.SysAuthority
	err := c.ShouldBindJSON(&authority)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(authority, utils.AuthorityVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if authBack, err := authorityService.CreateAuthority(authority); err != nil {
		global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create"+err.Error(), c)
	} else {
		_ = menuService.AddMenuAuthority(systemReq.DefaultMenu(), authority.AuthorityId)
		_ = casbinService.UpdateCasbin(authority.AuthorityId, systemReq.DefaultCasbin())
		response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "Successful creation", c)
	}
}

// CopyAuthority
// @Tags      Authority
// @Summary   拷贝角色
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      response.SysAuthorityCopyResponse                                  true  "旧角色id, 新权限id, 新权限名, 新父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "拷贝角色,返回包括系统角色详情"
// @Router    /authority/copyAuthority [post]
func (a *AuthorityApi) CopyAuthority(c *gin.Context) {
	var copyInfo systemRes.SysAuthorityCopyResponse
	err := c.ShouldBindJSON(&copyInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(copyInfo, utils.OldAuthorityVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(copyInfo.Authority, utils.AuthorityVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	authBack, err := authorityService.CopyAuthority(copyInfo)
	if err != nil {
		global.GVA_LOG.Error("拷贝失败!", zap.Error(err))
		response.FailWithMessage("拷贝失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "拷贝成功", c)
}

// DeleteAuthority
// @Tags      Authority
// @Summary   Delete 角色
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority            true  "Delete 角色"
// @Success   200   {object}  response.Response{msg=string}  "Delete 角色"
// @Router    /authority/deleteAuthority [post]
func (a *AuthorityApi) DeleteAuthority(c *gin.Context) {
	var authority system.SysAuthority
	err := c.ShouldBindJSON(&authority)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(authority, utils.AuthorityIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = authorityService.DeleteAuthority(&authority)
	if err != nil { // Delete 角色之前需要判断是否有用户正在使用此角色
		global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete"+err.Error(), c)
		return
	}
	response.OkWithMessage("successfully deleted", c)
}

// UpdateAuthority
// @Tags      Authority
// @Summary   更新角色信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "更新角色信息,返回包括系统角色详情"
// @Router    /authority/updateAuthority [post]
func (a *AuthorityApi) UpdateAuthority(c *gin.Context) {
	var auth system.SysAuthority
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(auth, utils.AuthorityVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	authority, err := authorityService.UpdateAuthority(auth)
	if err != nil {
		global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure"+err.Error(), c)
		return
	}
	response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authority}, "update completed", c)
}

// GetAuthorityList
// @Tags      Authority
// @Summary   分页获取角色列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取角色列表,返回包括列表,总数,页码,每页数量"
// @Router    /authority/getAuthorityList [post]
func (a *AuthorityApi) GetAuthorityList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := authorityService.GetAuthorityInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Fail!", zap.Error(err))
		response.FailWithMessage("Fail"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Successful", c)
}

// SetDataAuthority
// @Tags      Authority
// @Summary   设置角色资源权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority            true  "设置角色资源权限"
// @Success   200   {object}  response.Response{msg=string}  "设置角色资源权限"
// @Router    /authority/setDataAuthority [post]
func (a *AuthorityApi) SetDataAuthority(c *gin.Context) {
	var auth system.SysAuthority
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(auth, utils.AuthorityIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = authorityService.SetDataAuthority(auth)
	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("设置成功", c)
}
