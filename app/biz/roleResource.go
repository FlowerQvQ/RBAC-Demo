package biz

import (
	"NewProject/app/data"
	"NewProject/app/scheme"
	"NewProject/models"
	"NewProject/pkg/wapper"
)

type RoleResourceBiz struct {
	RoleResourceData *data.RoleResourceData
}

func NewRoleResourceBiz(roleResourceData *data.RoleResourceData) *RoleResourceBiz {
	return &RoleResourceBiz{
		RoleResourceData: roleResourceData,
	}
}

func (b *RoleResourceBiz) RoleResourceBind(AddRoleResourceBind scheme.AddRoleResourceReq) ([]models.RoleResource, wapper.ErrorCode) {
	bindData, errCode := b.RoleResourceData.RoleResourceBind(AddRoleResourceBind)
	if errCode != nil {
		return nil, wapper.RoleResourceBindFailed
	}
	return bindData, wapper.Success
}
