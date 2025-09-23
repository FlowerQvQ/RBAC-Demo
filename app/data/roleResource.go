package data

import (
	"NewProject/app/scheme"
	"NewProject/models"
	"NewProject/pkg/wapper"
	"context"
	"encoding/json"
	"fmt"
)

type RoleResourceData struct {
	DB      *Data
	RedisDB *RedisData
}

func NewRoleResourceData(data *Data, redisDB *RedisData) *RoleResourceData {
	return &RoleResourceData{
		DB:      data,
		RedisDB: redisDB,
	}
}

// 角色资源绑定   (角色资源绑定时清除缓存--保证数据一致性)
func (d *RoleResourceData) RoleResourceBind(AddRoleResourceBind scheme.AddRoleOwnedResourceInfo) ([]models.RoleResource, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("role_resources:%d", AddRoleResourceBind.RoleId)
	var (
		bindData []models.RoleResource
		err      error
	)
	for _, resourceId := range AddRoleResourceBind.ResourceId {
		bindData = append(bindData, models.RoleResource{
			RoleId:     AddRoleResourceBind.RoleId,
			ResourceId: resourceId,
			CreatedBy:  AddRoleResourceBind.CreatedBy,
		})
	}
	err = d.DB.DBClient.Model(&models.RoleResource{}).CreateInBatches(bindData, 10).Error
	if err != nil {
		return nil, err
	}
	// 清除缓存
	d.RedisDB.RedisClient.Del(ctx, cacheKey)
	return bindData, nil
}

// 查询角色拥有的资源列表
func (d *RoleResourceData) GetRoleOwnedResourceList(roleId scheme.RoleOwnedResourceListReq) ([]models.RoleResource, wapper.ErrorCode) {
	//这种设计模式叫做"缓存优先策略"（Cache-Aside Pattern），可以显著提高系统的响应速度，减少数据库压力，
	//特别是在数据不经常变化的场景下非常有效。如果缓存中没有数据或数据解析失败，程序将继续执行数据库查询逻辑，
	//然后将查询结果存入缓存供下次使用。

	//创建一个空的上下文（context）,用于控制后续 Redis 操作的超时和取消。
	ctx := context.Background()
	//使用 fmt.Sprintf 构造一个 Redis 缓存键。键的格式为 role_resources:{roleId}，
	//其中 {roleId} 是具体的角色 ID,这样可以确保每个角色的资源列表都有唯一的缓存键。
	cacheKey := fmt.Sprintf("role_resources:%d", roleId.RoleId)
	//尝试从Redis获取

	//使用 Redis 客户端的 Get 方法尝试从 Redis 中获取指定键的值
	//d.RedisDB.RedisClient 是 Redis 客户端实例
	//Get(ctx, cacheKey) 执行 GET 命令获取缓存数据
	//.Result() 获取命令执行结果(string)
	//cached 变量存储获取到的缓存数据（string格式）
	cached, err := d.RedisDB.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var roleOwnedResourceData []models.RoleResource
		//尝试将从 Redis 获取到的缓存数据（字符串格式）解析为 []models.RoleResource 类型：
		//json.Unmarshal() 是 Go 标准库函数，用于将 JSON 格式的字节数据解析为 Go 数据结构。
		//[]byte(cached) 将字符串转换为字节切片。
		//&roleOwnedResourceData 传递变量的地址，以便解析结果能存储到该变量中。
		//== nil 检查解析是否成功（返回 nil 表示成功）。
		if json.Unmarshal([]byte(cached), &roleOwnedResourceData) == nil {
			return roleOwnedResourceData, wapper.Success
		}
	}

	var roleOwnedResourceData []models.RoleResource
	err = d.DB.DBClient.Model(&models.RoleResource{}).Where("role_id = ?", roleId.RoleId).Find(&roleOwnedResourceData).Error
	if err != nil {
		return nil, wapper.GetRoleResourceFailed
	}
	// 存入Redis
	//json.Marshal(roleOwnedResourceData) 将 Go 数据结构 roleOwnedResourceData 序列化为 JSON 格式的字节切片
	if data, err := json.Marshal(roleOwnedResourceData); err == nil {
		d.RedisDB.RedisClient.Set(ctx, cacheKey, data, 0)
	}
	return roleOwnedResourceData, wapper.Success
}

// 批量删除角色拥有的资源    	批量删除角色拥有的资源时清除缓存
func (d *RoleResourceData) DelRoleOwnedResource(delRoleResourceInfo scheme.DelRoleOwnedResourceReq) wapper.ErrorCode {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("role_resources:%d", delRoleResourceInfo.RoleId)

	if delRoleResourceInfo.RoleId <= 0 || len(delRoleResourceInfo.ResourceId) == 0 {
		return wapper.DeleteDataFailed
	}
	err := d.DB.DBClient.Model(&models.RoleResource{}).
		Where("role_id = ? AND resource_id IN ?", delRoleResourceInfo.RoleId, delRoleResourceInfo.ResourceId).
		Delete(&models.RoleResource{}).Error
	if err != nil {
		return wapper.DeleteRoleOwnedResourceFailed
	}
	// 清除缓存
	d.RedisDB.RedisClient.Del(ctx, cacheKey)
	return wapper.Success
}
