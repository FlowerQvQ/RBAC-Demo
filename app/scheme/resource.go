package scheme

import "NewProject/models"

// 获取资源列表
type ResourceListReq struct {
	Name     string `json:"name" biding:"omitempty"`
	IsStatus int    `json:"is_status" binding:"omitempty"`
	Path     string `json:"path" biding:"omitempty"`
	Type     int    `json:"type" biding:"omitempty"`
	Pagination
}
type ResourceListResp struct {
	Pagination
	ResourceList []models.Resource `json:"resourceList"`
	Total        int               `json:"total"` //总页数
}

// 查询资源
type ResourceGetReq struct {
	Id int `json:"id" binding:"required"`
}

// 增加资源
type ResourceCreateReq struct {
	Pid         int    `json:"pid"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Path        string `json:"path"`
	Type        int    `json:"type"`
	Status      int    `json:"status"`
}

// 更新资源
type ResourceUpdateReq struct {
	Id          int    `json:"id" binding:"required"`
	Pid         int    `json:"pid"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Path        string `json:"path" binding:"required"`
	Type        int    `json:"type"`
	Status      int    `json:"status"`
	UpdateBy    string `json:"update_by"`
}

// 删除资源
type ResourceDelReq struct {
	Id int `json:"id" binding:"required"`
}
