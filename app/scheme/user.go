package scheme

import (
	"NewProject/models"
	"time"
)

const (
	UserActive   = 1
	UserInActive = 2
)

// 分页
type Pagination struct {
	Page  int `json:"page" binding:"min=1"`  //请求的页码
	Limit int `json:"limit" binding:"min=1"` //每页的数量
}

// 用户注册
type UserRegisterReq struct {
	Username     string `json:"username" binding:"required"`
	Email        string `json:"email" binding:"required"`
	PasswordHash string `json:"password_hash" binding:"required,min=8,max=16"`
}

// 登录
type UserLoginReq struct {
	LoginName    string `json:"login_name" binding:"required"` //账号和密码--2选1
	PasswordHash string `json:"password_hash" binding:"required,min=8,max=16"`
}

// 查询用户列表
type UserListReq struct {
	QueryName string `json:"queryName" biding:"omitempty"`
	IsActive  int    `json:"is_active" binding:"omitempty"`
	Pagination
}
type UserListResp struct {
	UserList []models.User `json:"userList"`
	Total    int           `json:"total"`
	Pagination
}

// 用户查询
type GetUserInfoReq struct {
	Id int64 `json:"id" binding:"required"`
}

// 修改密码和邮箱和修改状态（传了修改，不传不改，bind一个什么东西，修改完密码要重新加密）
// 修改用户信息
type UserUpdateReq struct {
	Id           int64  `json:"id" binding:"required"`
	Email        string `json:"email" binding:"omitempty"`
	Username     string `json:"username" binding:"omitempty"`
	PasswordHash string `json:"password_hash" binding:"omitempty"`
	IsActive     int    `json:"is_active" binding:"omitempty" common:"1开启/2禁用"`
}
type UserUpdateResp struct {
	Id           int64     `json:"id" binding:"required"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"password_hash"`
	IsActive     int       `json:"is_active"`
	CreatedBy    string    `json:"created_by"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedBy    string    `json:"updated_by"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// 删除用户
type DelUserReq struct {
	Id int64 `json:"id" binding:"required"`
}
