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

// 增加角色
func (b *RoleBiz) AddRole(roleInfo scheme.AddRoleReq) (models.Role, wapper.ErrorCode) {
	roleData, err := b.RoleData.AddRole(roleInfo)
	if err != nil {
		return models.Role{}, wapper.AddRoleFailed
	}
	return roleData, wapper.Success
}
