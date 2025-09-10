package data

import (
	"NewProject/app/scheme"
	"NewProject/models"
	"NewProject/pkg/wapper"
)

type RoleResourceData struct {
	DB *Data
}

func NewRoleResourceData(data *Data) *RoleResourceData {
	return &RoleResourceData{
		DB: data,
	}
}

// 角色资源绑定
func (d *RoleResourceData) RoleResourceBind(AddRoleResourceBind scheme.AddRoleResourceReq) ([]models.RoleResource, error) {
	var (
		bindData []models.RoleResource
		err      error
	)
	for _, resourceId := range AddRoleResourceBind.ResourceId {
		bindData = append(bindData, models.RoleResource{
			RoleId:     AddRoleResourceBind.RoleId,
			ResourceId: resourceId,
			CreatedBy:  AddRoleResourceBind.CreatedBy,
			Status:     1,
		})
	}
	err = d.DB.DBClient.Model(&models.RoleResource{}).CreateInBatches(bindData, 10).Error
	if err != nil {
		return nil, err
	}
	return bindData, nil
}

// 查询角色拥有的资源列表
func (d *RoleResourceData) GetRoleOwnedResourceList(roleId scheme.RoleOwnedResourceListReq) ([]models.RoleResource, wapper.ErrorCode) {
	var roleOwnedResourceData []models.RoleResource
	err := d.DB.DBClient.Model(&models.RoleResource{}).Where("role_id = ?", roleId.RoleId).Find(&roleOwnedResourceData).Error
	if err != nil {
		return nil, wapper.GetRoleResourceFailed
	}
	return roleOwnedResourceData, wapper.Success
}
