package scheme

const (
	StatusOK  = 1
	StatusOff = 2
)

type AddUserRoleReq struct {
	UserId int   `json:"user_Id" binding:"required"`
	RoleId []int `json:"role_Id" binding:"required"`
}

// 自动获取用户名
type GetUsernameReq struct {
	UserId    int    `json:"user_Id" binding:"required"`
	RoleId    []int  `json:"role_Id" binding:"required"`
	CreatedBy string `json:"created_by"`
}

// 查询用户拥有的资源，查询用户拥有的角色都用这个结构体中的UserId查询
type GetUserOwnedRoleReq struct {
	UserId int `json:"user_id" binding:"required"`
}
type UserOwnedRoleReq struct {
	UserId int `json:"user_id" binding:"required"`
}

type UserOwnedResourceResp struct {
	Resources []map[string]interface{} `json:"resources"`
	Path      []string                 `json:"path"`
}

type DelUserOwnedRoleReq struct {
	UserId int   `json:"user_id" binding:"required"`
	RoleId []int `json:"role_id" binding:"required"`
}
