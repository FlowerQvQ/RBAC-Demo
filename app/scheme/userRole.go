package scheme

type AddUserRoleReq struct {
	UserId    int    `json:"user_Id" binding:"required"`
	RoleId    []int  `json:"role_Id" binding:"required"`
	Status    int    `json:"status" `
	CreatedBy string `json:"created_By" binding:"required"`
}
