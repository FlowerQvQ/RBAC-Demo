package biz

import (
	"NewProject/app/data"
	"NewProject/app/scheme"
	"NewProject/models"
	"NewProject/pkg/wapper"
)

type RoleBiz struct {
	RoleData *data.RoleData
}

func NewRoleBiz(roleData *data.RoleData) *RoleBiz {
	return &RoleBiz{
		RoleData: roleData,
	}
}

// 查询角色列表
func (b *RoleBiz) GetRoleList(roleListReq scheme.RoleListReq) (scheme.RoleListResp, wapper.ErrorCode) {
	roleListData, errCode := b.RoleData.GetRoleList(roleListReq)
	if errCode != wapper.Success {
		return scheme.RoleListResp{}, errCode
	}
	return roleListData, wapper.Success
}

// 查询角色
func (b *RoleBiz) GetRole(getRoleReq scheme.GetRoleReq) (models.Role, wapper.ErrorCode) {
	roleData, errCode := b.RoleData.GetRole(getRoleReq)
	if errCode != wapper.Success {
		return models.Role{}, wapper.GetRoleFailed
	}
	return roleData, wapper.Success
}

// 增加角色
func (b *RoleBiz) AddRole(roleInfo models.Role) (models.Role, wapper.ErrorCode) {
	roleData, err := b.RoleData.AddRole(roleInfo)
	if err != nil {
		return models.Role{}, wapper.AddRoleFailed
	}
	return roleData, wapper.Success
}

// 修改角色
func (b *RoleBiz) UpdateRole(updateRoleInfo models.Role) (models.Role, wapper.ErrorCode) {
	roleData, err := b.RoleData.UpdateRole(updateRoleInfo)
	if err != wapper.Success {
		return models.Role{}, wapper.UpdateRoleFailed
	}
	return roleData, wapper.Success
}

// 删除角色
func (b *RoleBiz) DelRole(delRoleReq scheme.RoleDelReq) wapper.ErrorCode {
	err := b.RoleData.DelRole(delRoleReq)
	if err != wapper.Success {
		return wapper.DelRoleFailed
	}
	return wapper.Success
}
