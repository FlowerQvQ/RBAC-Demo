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

// 查询用户拥有的资源列表
func (b *UserRoleBiz) UserOwnedResource(userId scheme.UserOwnedRoleReq) (scheme.UserOwnedResourceResp, wapper.ErrorCode) {
	UserOwnedResourceData, err := b.userRoleData.UserOwnedResource(userId)
	if err != nil {
		return scheme.UserOwnedResourceResp{}, wapper.GetResourceListFailed
	}
	return UserOwnedResourceData, wapper.Success
}
