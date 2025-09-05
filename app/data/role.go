package data

import (
	"NewProject/app/scheme"
	"NewProject/models"
)

type RoleData struct {
	DB *Data
}

func NewRoleData(data *Data) *RoleData {
	return &RoleData{
		DB: data,
	}
}

// 增加角色
func (d *RoleData) AddRole(roleInfo scheme.AddRoleReq) (models.Role, error) {
	var (
		roleData = models.Role{}
		err      error
	)
	roleData = models.Role{
		Name:        roleInfo.Name,
		Description: roleInfo.Description,
		CreatedBy:   roleInfo.CreatedBy,
	}
	err = d.DB.DBClient.Model(&models.Role{}).Create(&roleData).Error
	if err != nil {
		return models.Role{}, err
	}
	return roleData, nil
}
