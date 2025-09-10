package service

import (
	"NewProject/app/biz"
	"NewProject/app/scheme"
	"NewProject/pkg/wapper"
	"github.com/gin-gonic/gin"
)

type RoleService struct {
	RoleBiz *biz.RoleBiz
}

func NewRoleService(roleBiz *biz.RoleBiz) *RoleService {
	return &RoleService{
		RoleBiz: roleBiz,
	}
}

// 查询角色列表
func (s *RoleService) GetRoleList(c *gin.Context) {
	var (
		roleListReq  scheme.RoleListReq
		roleListData scheme.RoleListResp
		errCode      wapper.ErrorCode
		err          error
	)
	err = c.ShouldBindJSON(&roleListReq)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}
	roleListData, errCode = s.RoleBiz.GetRoleList(roleListReq)
	if errCode != wapper.Success {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}
	wapper.ResSuccess(c, roleListData)
}

// 查询角色
func (s *RoleService) GetRole(c *gin.Context) {
	var (
		getRoleReq scheme.GetRoleReq
		err        error
	)
	err = c.ShouldBindJSON(&getRoleReq)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}
	roleData, errCode := s.RoleBiz.GetRole(getRoleReq)
	if errCode != wapper.Success {
		wapper.ResError(c, wapper.GetRoleFailed)
		return
	}
	wapper.ResSuccess(c, roleData)
}

// 增加角色
func (s *RoleService) AddRole(c *gin.Context) {
	var (
		roleInfo scheme.AddRoleReq
		err      error
	)
	err = c.ShouldBindJSON(&roleInfo)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
	}
	roleData, errCode := s.RoleBiz.AddRole(roleInfo)
	if errCode != wapper.Success {
		wapper.ResError(c, wapper.AddRoleResourceFailed)
		return
	}
	wapper.ResSuccess(c, roleData)
}

// 修改角色
func (s *RoleService) UpdateRole(c *gin.Context) {
	var (
		updateReq scheme.UpdateRoleReq
		err       error
	)
	err = c.ShouldBindJSON(&updateReq)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}
	updatedData, errCode := s.RoleBiz.UpdateRole(updateReq)
	if errCode != wapper.Success {
		wapper.ResError(c, wapper.UpdateRoleFailed)
		return
	}
	wapper.ResSuccess(c, updatedData)
}

// 删除角色
func (s *RoleService) DelRole(c *gin.Context) {
	var (
		delRoleId scheme.RoleDelReq
		err       error
	)
	err = c.ShouldBindJSON(&delRoleId)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}
	errCode := s.RoleBiz.DelRole(delRoleId)
	if errCode != wapper.Success {
		wapper.ResError(c, wapper.DelRoleFailed)
		return
	}
	wapper.ResSuccess(c, "删除成功")
}
