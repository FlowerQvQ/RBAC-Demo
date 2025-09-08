package scheme

type AddUserRoleReq struct {
	UserId    int    `json:"user_Id" binding:"required"`
	RoleId    []int  `json:"role_Id" binding:"required"`
	Status    int    `json:"status" `
	CreatedBy string `json:"created_By" binding:"required"`
}

const (
	StatusOK  = 1
	StatusOff = 2
)

type UserOwnedRoleReq struct {
	UserId int `json:"user_id" binding:"required"`
}
type UserOwnedResourceResp struct {
	Resources []map[string]interface{} `json:"role_id"`
	Path      []string                 `json:"path"`
}
