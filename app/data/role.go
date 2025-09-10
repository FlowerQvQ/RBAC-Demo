package data

import (
	"NewProject/app/scheme"
	"NewProject/models"
	"NewProject/pkg/wapper"
)

type RoleData struct {
	DB *Data
}

func NewRoleData(data *Data) *RoleData {
	return &RoleData{
		DB: data,
	}
}

// 查询角色列表
func (d *RoleData) GetRoleList(roleListReq scheme.RoleListReq) (scheme.RoleListResp, wapper.ErrorCode) {
	var (
		//roleListData []models.Role
		role       scheme.RoleListResp
		totalRoles int64
		total      int
		err        error
		roleData   = d.DB.DBClient.Model(&models.Role{}) //创建一个基于 models.Role{} 模型的数据库查询构建器实例，并将其赋值给 roleListData 变量。
		//1. d.DB.DBClient - 这是 GORM 数据库连接实例
		//2. Model(&models.Role{}) - 指定查询的模型为 Role 表
		//3. roleListData - 接收查询构建器的变量
	)
	if roleListReq.Id != 0 {
		roleData = roleData.Where("id = ?", roleListReq.Id)
	}
	if roleListReq.Name != "" {
		roleData = roleData.Where("name = ?", "%"+roleListReq.Name+"%")
	}
	if roleListReq.Description != "" {
		roleData = roleData.Where("description = ?", "%"+roleListReq.Description+"%")
	}
	if roleListReq.Status != 0 {
		roleData = roleData.Where("status = ?", roleListReq.Status)
	}
	err = roleData.Count(&totalRoles).Error
	if err != nil {
		return scheme.RoleListResp{}, wapper.GetTotalRecordsFailed
	}
	total = int((totalRoles + int64(roleListReq.Limit) - 1) / int64(roleListReq.Limit))
	if roleListReq.Limit > 0 && roleListReq.Page > 0 {
		//计算偏移量，跳过前n条记录，从第n+1条记录开始取
		offset := (roleListReq.Page - 1) * roleListReq.Limit
		//设置分页查询的偏移量和限制。Offset:跳过前n条记录，Limit:限制返回的记录数
		roleData = roleData.Offset(offset).Limit(roleListReq.Limit)
	}
	err = roleData.Find(&role.RoleList).Error
	if err != nil {
		return scheme.RoleListResp{}, wapper.GetRoleListFailed
	}
	role = scheme.RoleListResp{
		RoleList: role.RoleList,
		Total:    total,
		Pagination: scheme.Pagination{
			Limit: roleListReq.Limit,
			Page:  roleListReq.Page,
		},
	}

	return role, wapper.Success
}

// 查询角色
func (d *RoleData) GetRole(getRoleReq scheme.GetRoleReq) (models.Role, wapper.ErrorCode) {
	var (
		roleData models.Role
		err      error
	)
	err = d.DB.DBClient.Model(&models.Role{}).Where("id = ?", getRoleReq.Id).First(&roleData, getRoleReq.Id).Error
	if err != nil {
		return models.Role{}, wapper.DataNotFound
	}
	return roleData, wapper.Success
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

// 更新角色信息
func (d *RoleData) UpdateRole(updateRoleReq scheme.UpdateRoleReq) (models.Role, wapper.ErrorCode) {
	updatedRoleData := models.Role{
		Id:          updateRoleReq.Id,
		Name:        updateRoleReq.Name,
		Description: updateRoleReq.Description,
		Status:      updateRoleReq.Status,
		UpdatedBy:   updateRoleReq.UpdateBy,
	}
	err := d.DB.DBClient.Model(&models.Role{}).Where("id = ?", updateRoleReq.Id).Updates(&updatedRoleData).Error
	if err != nil {
		return models.Role{}, wapper.UpdateRoleFailed
	}
	var updatedRole models.Role
	err = d.DB.DBClient.Model(&models.Role{}).Where("id = ?", updateRoleReq.Id).First(&updatedRole).Error
	if err != nil {
		return models.Role{}, wapper.DataNotFound
	}
	return updatedRoleData, wapper.Success
}

// 角色删除
func (d *RoleData) DelRole(delRoleReq scheme.RoleDelReq) wapper.ErrorCode {
	err := d.DB.DBClient.Model(&models.Role{}).Where("id = ?", delRoleReq.Id).Update("status", scheme.StatusOff).Error
	if err != nil {
		return wapper.DelRoleFailed
	}
	return wapper.Success
}
