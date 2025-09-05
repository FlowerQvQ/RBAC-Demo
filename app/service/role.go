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
