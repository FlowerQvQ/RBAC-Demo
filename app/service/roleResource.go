package service

import (
	"NewProject/app/biz"
	"NewProject/app/scheme"
	"NewProject/pkg/wapper"
	"github.com/gin-gonic/gin"
)

type RoleResourceService struct {
	RoleResourceBiz *biz.RoleResourceBiz
}

func NewRoleResourceService(roleResourceService *biz.RoleResourceBiz) *RoleResourceService {
	return &RoleResourceService{
		RoleResourceBiz: roleResourceService,
	}
}

func (s *RoleResourceService) RoleResourceBind(c *gin.Context) {
	var bindInfo scheme.AddRoleResourceReq
	err := c.ShouldBind(&bindInfo)
	if err != nil {
		wapper.ResError(c, wapper.AddResourceFailed)
		return
	}
	bindData, errCode := s.RoleResourceBiz.RoleResourceBind(bindInfo)
	if errCode != wapper.Success {
		wapper.ResError(c, wapper.AddResourceFailed)
		return
	}
	wapper.ResSuccess(c, bindData)
}
