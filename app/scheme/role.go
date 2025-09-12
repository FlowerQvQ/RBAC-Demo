package scheme

import "NewProject/models"

// 查询角色列表
type RoleListReq struct {
	Id          int    `json:"id" binding:"omitempty,min=1"`
	Name        string `json:"name" binding:"omitempty"`
	Description string `json:"description" binding:"omitempty"`
	Status      int    `json:"status" binding:"omitempty,oneof=0 1 2"`
	Pagination
}

type RoleListResp struct {
	RoleList []models.Role `json:"list"`
	Total    int           `json:"total"`
	Pagination
}

// 查询角色信息
type GetRoleReq struct {
	Id int `json:"id" binding:"required"`
}

// 增加角色
type AddRoleReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// 更新角色信息
type UpdateRoleReq struct {
	Id          int    `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

// 删除角色信息
type RoleDelReq struct {
	Id int `json:"id" binding:"required"`
}
