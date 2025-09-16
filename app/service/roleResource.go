package service

import (
	"NewProject/app/biz"
	"NewProject/app/scheme"
	"NewProject/pkg/wapper"
	"github.com/gin-gonic/gin"
)

type RoleResourceService struct {
	RoleResourceBiz *biz.RoleResourceBiz
}

func NewRoleResourceService(roleResourceService *biz.RoleResourceBiz) *RoleResourceService {
	return &RoleResourceService{
		RoleResourceBiz: roleResourceService,
	}
}

// 角色资源绑定
func (s *RoleResourceService) RoleResourceBind(c *gin.Context) {
	var bindInfo scheme.AddRoleResourceReq
	err := c.ShouldBind(&bindInfo)
	if err != nil {
		wapper.ResError(c, wapper.AddResourceFailed)
		return
	}
	usernameInterface, exists := c.Get("username")
	if !exists {
		wapper.ResError(c, wapper.GetUserNameFailed)
		return
	}
	username, ok := usernameInterface.(string)
	if !ok {
		wapper.ResError(c, wapper.TypeAssertionFailed)
		return
	}
	addRoleResourceInfo := scheme.AddRoleOwnedResourceInfo{
		RoleId:     bindInfo.RoleId,
		ResourceId: bindInfo.ResourceId,
		CreatedBy:  username,
	}

	bindData, errCode := s.RoleResourceBiz.RoleResourceBind(addRoleResourceInfo)
	if errCode != wapper.Success {
		wapper.ResError(c, wapper.AddResourceFailed)
		return
	}
	wapper.ResSuccess(c, bindData)
}

// 查询角色拥有的资源列表
func (s *RoleResourceService) GetRoleOwnedResourceList(c *gin.Context) {
	var (
		roleId scheme.RoleOwnedResourceListReq
		err    error
	)
	err = c.ShouldBindJSON(&roleId)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}
	roleOwnedResourceData, errCode := s.RoleResourceBiz.GetRoleOwnedResourceList(roleId)
	if errCode != wapper.Success {
		wapper.ResError(c, wapper.GetRoleResourceFailed)
		return
	}
	wapper.ResSuccess(c, roleOwnedResourceData)
}

// 批量删除角色拥有的资源
func (s *RoleResourceService) DelRoleOwnedResource(c *gin.Context) {
	var (
		delRoleResourceInfo scheme.DelRoleOwnedResourceReq
		err                 error
	)
	err = c.ShouldBindJSON(&delRoleResourceInfo)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}
	errCode := s.RoleResourceBiz.DelRoleOwnedResource(delRoleResourceInfo)
	if errCode != wapper.Success {
		wapper.ResError(c, wapper.DeleteRoleOwnedResourceFailed)
		return
	}
	wapper.ResSuccess(c, "删除成功")
}
