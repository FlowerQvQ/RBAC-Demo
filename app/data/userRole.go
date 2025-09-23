package data

import (
	"NewProject/app/scheme"
	"NewProject/models"
	"NewProject/pkg/wapper"
	"context"
	"encoding/json"
	"fmt"
)

type UserRoleData struct {
	DB      *Data
	RedisDB *RedisData
}

func NewRoleUserData(mysqlDB *Data, redisDB *RedisData) *UserRoleData {
	return &UserRoleData{
		DB:      mysqlDB,
		RedisDB: redisDB,
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
func (d *UserRoleData) UserOwnedRole(userId scheme.GetUserOwnedRoleReq) ([]models.UserRole, wapper.ErrorCode) {
	var (
		userOwnedRoleData []models.UserRole
		err               error
		cached            string
	)
	ctx := context.Background()
	cacheKey := fmt.Sprintf("userOwnedRole:%d", userId.UserId)

	cached, err = d.RedisDB.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		//添加日志查看Redis中存储的数据
		fmt.Printf("Redis cached data: %s\n", cached)

		var parsedData []models.UserRole
		if json.Unmarshal([]byte(cached), &parsedData) == nil {
			return parsedData, wapper.Success
		}
	}
	//缓存中没有，则从数据库获取
	err = d.DB.DBClient.Model(&models.UserRole{}).Where("user_id = ?", userId.UserId).Find(&userOwnedRoleData).Error
	if err != nil {
		return nil, wapper.GetUserRoleFailed
	}
	//从数据库获取后，存入redis，并设置过期时间
	redisData, err := json.Marshal(userOwnedRoleData)
	if err == nil {
		d.RedisDB.RedisClient.Set(ctx, cacheKey, redisData, 0) //设置过期时间，0为永不过期，角色变更时手动删除
	}
	return userOwnedRoleData, wapper.Success
}

// 查询用户拥有的资源列表
func (d *UserRoleData) UserOwnedResource(userId scheme.UserOwnedRoleReq) (scheme.UserOwnedResourceResp, wapper.ErrorCode) {
	var (
		roleIDs           []int
		pathList          []string
		resourceList      []map[string]interface{}
		err               error
		userOwnedRoleData scheme.UserOwnedResourceResp
	)
	ctx := context.Background()
	cacheKey := fmt.Sprintf("userOwnedResource:%d", userId.UserId)
	////尝试从Redis获取
	cached, err := d.RedisDB.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var cachedResp scheme.UserOwnedResourceResp
		if json.Unmarshal([]byte(cached), &cachedResp) == nil {
			return cachedResp, wapper.Success
		}
	}
	//缓存中没有，则从数据库获取,并存入redis
	err = d.DB.DBClient.Model(&models.UserRole{}).Where("user_id = ?", userId.UserId).Pluck("role_id", &roleIDs).Error
	if err != nil {
		return scheme.UserOwnedResourceResp{}, wapper.GetUserIdFailed
	}
	if len(roleIDs) == 0 {
		return scheme.UserOwnedResourceResp{}, wapper.ThisUserHasNoRole
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
		return scheme.UserOwnedResourceResp{}, wapper.GetUserResourceFailed
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
	//存入Redis
	redisData, err := json.Marshal(resourceList)
	if err == nil {
		d.RedisDB.RedisClient.Set(ctx, cacheKey, redisData, 0) //设置过期时间，0为永不过期，角色变更时手动删除
	}

	return userOwnedRoleData, wapper.Success
}

// 删除用户拥有的角色
func (d *UserRoleData) DelUserRole(delId scheme.DelUserOwnedRoleReq) wapper.ErrorCode {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("user_role:%d", delId.UserId)
	//删除缓存
	d.RedisDB.RedisClient.Del(ctx, cacheKey)

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
