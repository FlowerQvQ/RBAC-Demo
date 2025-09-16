package scheme

// 增加角色资源
type AddRoleResourceReq struct {
	RoleId     int   `json:"role_id" binding:"required"`
	ResourceId []int `json:"resource_id" binding:"required" `
}

type AddRoleOwnedResourceInfo struct {
	RoleId     int    `json:"role_id" binding:"required"`
	ResourceId []int  `json:"resource_id" binding:"required"`
	CreatedBy  string `json:"created_by"`
}

// 查询角色拥有的资源列表
type RoleOwnedResourceListReq struct {
	RoleId int `json:"role_id" binding:"required"`
}

// 批量删除角色拥有的资源
type DelRoleOwnedResourceReq struct {
	RoleId     int   `json:"role_id" binding:"required"`
	ResourceId []int `json:"resource_id" binding:"required"`
}
