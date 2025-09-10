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
	bindData, errCode := s.RoleResourceBiz.RoleResourceBind(bindInfo)
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
