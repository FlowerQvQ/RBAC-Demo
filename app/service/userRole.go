package service

import (
	"NewProject/app/biz"
	"NewProject/app/scheme"
	"NewProject/pkg/wapper"
	"github.com/gin-gonic/gin"
)

type UserRoleService struct {
	userRoleBiz *biz.UserRoleBiz
}

func NewUserRoleService(userRoleBiz *biz.UserRoleBiz) *UserRoleService {
	return &UserRoleService{
		userRoleBiz: userRoleBiz,
	}
}

// 给用户授权角色
func (s *UserRoleService) AddUserRole(c *gin.Context) {
	var addUserRole scheme.AddUserRoleReq
	err := c.ShouldBindJSON(&addUserRole)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
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
	addUserRoleInfo := scheme.GetUsernameReq{
		UserId:    addUserRole.UserId,
		RoleId:    addUserRole.RoleId,
		CreatedBy: username,
	}
	addUserRoleData, errCode := s.userRoleBiz.AddUserRole(addUserRoleInfo)
	if errCode != wapper.Success {
		wapper.ResError(c, wapper.AddUserRoleFailed)
		return
	}
	wapper.ResSuccess(c, addUserRoleData)
}

// 查询用户拥有的角色列表
func (s *UserRoleService) UserOwnedRole(c *gin.Context) {
	var (
		userId scheme.GetUserOwnedRoleReq
		err    error
	)
	err = c.ShouldBindJSON(&userId)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}
	userOwnedRoleData, errCode := s.userRoleBiz.UserOwnedRole(userId)
	if errCode != wapper.Success {
		wapper.ResError(c, errCode)
		return
	}
	wapper.ResSuccess(c, userOwnedRoleData)
}

// 查询用户拥有的资源列表
func (s *UserRoleService) UserOwnedResource(c *gin.Context) {
	var (
		userId scheme.UserOwnedRoleReq
		err    error
	)
	err = c.ShouldBindJSON(&userId)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}
	UserOwnedResourceData, errCode := s.userRoleBiz.UserOwnedResource(userId)
	if errCode != wapper.Success {
		wapper.ResError(c, wapper.GrtUserResourceFailed)
		return
	}
	wapper.ResSuccess(c, UserOwnedResourceData)
}

// 删除用户拥有的角色
func (s *UserRoleService) DelUserRole(c *gin.Context) {
	var (
		delId scheme.DelUserOwnedRoleReq
		err   error
	)
	err = c.ShouldBindJSON(&delId)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
	}
	errCode := s.userRoleBiz.DelUserRole(delId)
	if errCode != wapper.Success {
		wapper.ResError(c, wapper.DelUserRoleFailed)
		return
	}
	wapper.ResSuccess(c, "删除成功")
}
