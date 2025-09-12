package service

import (
	"NewProject/app/biz"
	"NewProject/app/scheme"
	"NewProject/models"
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
		addRoleReq scheme.AddRoleReq
		roleInfo   models.Role
		err        error
	)
	err = c.ShouldBindJSON(&addRoleReq)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
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
	roleInfo = models.Role{
		Name:        addRoleReq.Name,
		Description: addRoleReq.Description,
		CreatedBy:   username,
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
		updateReq   scheme.UpdateRoleReq
		err         error
		updatedInfo models.Role
	)
	err = c.ShouldBindJSON(&updateReq)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}
	usernameInfo, exists := c.Get("username")
	if !exists {
		wapper.ResError(c, wapper.GetUserNameFailed)
		return
	}
	username, ok := usernameInfo.(string)
	if !ok {
		wapper.ResError(c, wapper.TypeAssertionFailed)
		return
	}
	updatedInfo = models.Role{
		Id:          updateReq.Id,
		Name:        updateReq.Name,
		Description: updateReq.Description,
		Status:      updateReq.Status,
		UpdatedBy:   username,
	}
	updatedData, errCode := s.RoleBiz.UpdateRole(updatedInfo)
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
