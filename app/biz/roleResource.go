package biz

import (
	"NewProject/app/data"
	"NewProject/app/scheme"
	"NewProject/models"
	"NewProject/pkg/wapper"
)

type RoleResourceBiz struct {
	RoleResourceData *data.RoleResourceData
}

func NewRoleResourceBiz(roleResourceData *data.RoleResourceData) *RoleResourceBiz {
	return &RoleResourceBiz{
		RoleResourceData: roleResourceData,
	}
}

// 角色资源绑定
func (b *RoleResourceBiz) RoleResourceBind(AddRoleResourceBind scheme.AddRoleOwnedResourceInfo) ([]models.RoleResource, wapper.ErrorCode) {
	bindData, errCode := b.RoleResourceData.RoleResourceBind(AddRoleResourceBind)
	if errCode != nil {
		return nil, wapper.RoleResourceBindFailed
	}
	return bindData, wapper.Success
}

// 查询角色拥有的资源列表
func (b *RoleResourceBiz) GetRoleOwnedResourceList(roleId scheme.RoleOwnedResourceListReq) ([]models.RoleResource, wapper.ErrorCode) {
	roleOwnedResourceData, errCode := b.RoleResourceData.GetRoleOwnedResourceList(roleId)
	if errCode != wapper.Success {
		return nil, wapper.GetRoleResourceFailed
	}
	return roleOwnedResourceData, wapper.Success
}

// 批量删除角色拥有的资源
func (b *RoleResourceBiz) DelRoleOwnedResource(delRoleResourceInfo scheme.DelRoleOwnedResourceReq) wapper.ErrorCode {
	err := b.RoleResourceData.DelRoleOwnedResource(delRoleResourceInfo)
	if err != wapper.Success {
		return wapper.DeleteRoleOwnedResourceFailed
	}
	return wapper.Success
}
