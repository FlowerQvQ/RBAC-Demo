package data

import (
	"NewProject/app/scheme"
	"NewProject/models"
	"NewProject/pkg/wapper"
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
func (d *UserRoleData) AddUserRole(addUserRole scheme.GetUsernameReq) ([]models.UserRole, error) {
	var (
		addUserRoleData []models.UserRole
		err             error
	)
	for _, roleId := range addUserRole.RoleId {
		addUserRoleData = append(addUserRoleData, models.UserRole{
			UserId:    addUserRole.UserId,
			RoleId:    roleId,
			CreatedBy: addUserRole.CreatedBy,
		})
	}
	err = d.DB.DBClient.Model(&models.UserRole{}).CreateInBatches(addUserRoleData, 10).Error
	if err != nil {
		return nil, err
	}
	return addUserRoleData, nil
}

// 查询用户拥有的角色列表
func (D *UserRoleData) UserOwnedRole(userId scheme.GetUserOwnedRoleReq) ([]models.UserRole, wapper.ErrorCode) {
	var (
		userOwnedRoleData []models.UserRole
	)
	err := D.DB.DBClient.Model(&models.UserRole{}).Where("user_id = ?", userId.UserId).Find(&userOwnedRoleData).Error
	if err != nil {
		return nil, wapper.GetUserRoleFailed
	}
	return userOwnedRoleData, wapper.Success
}

// 查询用户拥有的资源列表
func (d *UserRoleData) UserOwnedResource(userId scheme.UserOwnedRoleReq) (scheme.UserOwnedResourceResp, error) {
	var (
		roleIDs           []int
		pathList          []string
		resourceList      []map[string]interface{}
		err               error
		userOwnedRoleData scheme.UserOwnedResourceResp
	)
	err = d.DB.DBClient.Model(&models.UserRole{}).Where("user_id = ?", userId.UserId).Pluck("role_id", &roleIDs).Error
	if err != nil {
		return scheme.UserOwnedResourceResp{}, err
	}
	if len(roleIDs) == 0 {
		return scheme.UserOwnedResourceResp{}, nil
	}

	// 查询用户通过角色关联的资源信息
	err = d.DB.DBClient.Model(&models.RoleResource{}).
		Select("role.id as role_id, role.name as role_name, role.description as role_description, "+
			"resource.id as resource_id, resource.name as resource_name, resource.description as resource_description, "+
			"resource.path ,  role_resource.created_at, role_resource.created_by").
		Joins("left join role on role_resource.role_id = role.id").
		Joins("left join resource on role_resource.resource_id = resource.id").
		Where("role_resource.role_id IN ? AND role_resource.status = ? AND resource.status = ?",
			roleIDs, scheme.StatusOK, scheme.StatusOK).
		Scan(&resourceList).Error
	if err != nil {
		return scheme.UserOwnedResourceResp{}, err
	}
	//遍历资源列表中的每个resource
	//检查resource中是否存在"path"字段
	//验证path值是否为字符串类型且非空
	//符合条件的路径被追加到pathList中
	for _, resource := range resourceList {
		if pathVal, exists := resource["path"]; exists {
			if path, ok := pathVal.(string); ok && path != "" {
				pathList = append(pathList, path)
			}
		}
	}

	userOwnedRoleData.Resources = resourceList
	userOwnedRoleData.Path = pathList
	return userOwnedRoleData, nil
}

// 删除用户拥有的角色
func (d *UserRoleData) DelUserRole(delId scheme.DelUserOwnedRoleReq) wapper.ErrorCode {
	if delId.UserId <= 0 || len(delId.RoleId) == 0 {
		return wapper.ParameterMissing
	}
	err := d.DB.DBClient.Model(&models.UserRole{}).
		Where("user_id = ? AND role_id IN ?", delId.UserId, delId.RoleId).
		Delete(&models.UserRole{}).Error
	if err != nil {
		return wapper.DelUserRoleFailed
	}
	return wapper.Success
}
