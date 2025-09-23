// app/data/cacheWarmUp.go
package data

import (
	"NewProject/app/scheme"
	"NewProject/models"
	"context"
	"encoding/json"
	"fmt"
)

// CacheWarmUp 缓存预热机制
func (d *Data) CacheWarmUp() error {
	ctx := context.Background()

	// 预热热门角色的资源列表
	if err := d.warmUpRoleResources(ctx); err != nil {
		return err
	}

	// 预热活跃用户的权限信息
	if err := d.warmUpUserRoles(ctx); err != nil {
		return err
	}

	// 可以添加更多预热逻辑
	return nil
}

// warmUpRoleResources 预热角色资源列表
func (d *Data) warmUpRoleResources(ctx context.Context) error {
	// 查询活跃的角色（可以根据需要调整条件）
	var roles []models.Role
	err := d.DBClient.Model(&models.Role{}).
		Where("status = ?", scheme.StatusOK).
		Limit(100). // 限制预热的角色数量
		Find(&roles).Error
	if err != nil {
		return err
	}

	// 为每个角色预热资源列表
	for _, role := range roles {
		// 查询角色拥有的资源
		var roleResources []models.RoleResource
		err := d.DBClient.Model(&models.RoleResource{}).
			Where("role_id = ? AND status = ?", role.Id, scheme.StatusOK).
			Find(&roleResources).Error
		if err != nil {
			continue // 跳过出错的角色
		}

		// 序列化并存入缓存
		if data, err := json.Marshal(roleResources); err == nil {
			cacheKey := fmt.Sprintf("role_resources:%d", role.Id)
			d.RedisDB.RedisClient.Set(ctx, cacheKey, data, 0) // 永久缓存或设置合适的过期时间
		}
	}

	return nil
}

// warmUpUserRoles 预热用户角色信息
func (d *Data) warmUpUserRoles(ctx context.Context) error {
	// 查询活跃用户（可以根据需要调整条件）
	var users []models.User
	err := d.DBClient.Model(&models.User{}).
		Where("is_active = ?", scheme.UserActive).
		Limit(1000). // 限制预热的用户数量
		Find(&users).Error
	if err != nil {
		return err
	}

	// 为每个用户预热角色信息
	for _, user := range users {
		// 查询用户拥有的角色
		var userRoles []models.UserRole
		err := d.DBClient.Model(&models.UserRole{}).
			Where("user_id = ? AND status = ?", user.Id, scheme.StatusOK).
			Find(&userRoles).Error
		if err != nil {
			continue // 跳过出错的用户
		}

		// 序列化并存入缓存
		if data, err := json.Marshal(userRoles); err == nil {
			cacheKey := fmt.Sprintf("userOwnedRole:%d", user.Id)
			d.RedisDB.RedisClient.Set(ctx, cacheKey, data, 0) // 永久缓存或设置合适的过期时间
		}
	}

	return nil
}
