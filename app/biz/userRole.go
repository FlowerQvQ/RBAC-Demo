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
func (b *UserRoleBiz) AddUserRole(addUserRole scheme.GetUsernameReq) ([]models.UserRole, wapper.ErrorCode) {
	addUserRoleData, err := b.userRoleData.AddUserRole(addUserRole)
	if err != nil {
		return nil, wapper.AddUserRoleFailed
	}
	return addUserRoleData, wapper.Success
}

// 查询用户拥有的角色列表
func (b *UserRoleBiz) UserOwnedRole(userId scheme.GetUserOwnedRoleReq) ([]models.UserRole, wapper.ErrorCode) {
	userOwnedRoleData, errCode := b.userRoleData.UserOwnedRole(userId)
	if errCode != wapper.Success {
		return nil, errCode
	}
	return userOwnedRoleData, wapper.Success
}

// 查询用户拥有的资源列表
func (b *UserRoleBiz) UserOwnedResource(userId scheme.UserOwnedRoleReq) (scheme.UserOwnedResourceResp, wapper.ErrorCode) {
	UserOwnedResourceData, err := b.userRoleData.UserOwnedResource(userId)
	if err != nil {
		return scheme.UserOwnedResourceResp{}, wapper.GetResourceListFailed
	}
	return UserOwnedResourceData, wapper.Success
}

// 删除用户拥有的角色
func (b *UserRoleBiz) DelUserRole(delId scheme.DelUserOwnedRoleReq) wapper.ErrorCode {
	err := b.userRoleData.DelUserRole(delId)
	if err != wapper.Success {
		return wapper.DelUserRoleFailed
	}
	return wapper.Success
}
