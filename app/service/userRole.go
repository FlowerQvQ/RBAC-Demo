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
	addUserRoleData, errCode := s.userRoleBiz.AddUserRole(addUserRole)
	if errCode != wapper.Success {
		wapper.ResError(c, wapper.AddUserRoleFailed)
		return
	}
	wapper.ResSuccess(c, addUserRoleData)
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
