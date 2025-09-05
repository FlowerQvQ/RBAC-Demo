package biz

import (
	"NewProject/app/data"
	"NewProject/app/scheme"
	"NewProject/models"
	"NewProject/pkg/wapper"
)

type UserRoleBiz struct {
	userRoleData *data.UserRoleData
}

func NewUserRoleBiz(userRoleData *data.UserRoleData) *UserRoleBiz {
	return &UserRoleBiz{
		userRoleData: userRoleData,
	}
}

// 给用户授权角色
func (b *UserRoleBiz) AddUserRole(addUserRole scheme.AddUserRoleReq) ([]models.UserRole, wapper.ErrorCode) {
	addUserRoleData, err := b.userRoleData.AddUserRole(addUserRole)
	if err != nil {
		return nil, wapper.AddUserRoleFailed
	}
	return addUserRoleData, wapper.Success
}
