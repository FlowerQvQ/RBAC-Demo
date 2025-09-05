package data

import (
	"NewProject/app/scheme"
	"NewProject/models"
)

type UserRoleData struct {
	DB *Data
}

func NewRoleUserData(data *Data) *UserRoleData {
	return &UserRoleData{
		DB: data,
	}
}

// 给用户授权角色
func (d *UserRoleData) AddUserRole(addUserRole scheme.AddUserRoleReq) ([]models.UserRole, error) {
	var (
		addUserRoleData []models.UserRole
		err             error
	)
	for _, roleId := range addUserRole.RoleId {
		addUserRoleData = append(addUserRoleData, models.UserRole{
			UserId:    addUserRole.UserId,
			RoleId:    roleId,
			CreatedBy: addUserRole.CreatedBy,
			Status:    addUserRole.Status,
		})
	}
	err = d.DB.DBClient.Model(&models.UserRole{}).CreateInBatches(addUserRoleData, 10).Error
	if err != nil {
		return nil, err
	}
	return addUserRoleData, nil
}
