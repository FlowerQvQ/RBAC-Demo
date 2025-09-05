package scheme

// 增加角色
type AddRoleReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
}
