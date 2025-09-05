package models

import (
	"time"
)

type User struct {
	Id           int64     `gorm:"column:id" json:"id"`
	Username     string    `gorm:"column:username" json:"username"`
	PasswordHash string    `gorm:"column:password_hash" json:"password"`
	Email        string    `gorm:"column:email" json:"email"`
	IsActive     int       `gorm:"column:is_active" json:"is_active"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy    string    `gorm:"column:created_by" json:"created_by"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy    string    `gorm:"column:updated_by" json:"updated_by"`
}

func (User) TableName() string {
	return "user"
}

type Resource struct {
	Pid         int       `gorm:"column:pid" json:"pid"`
	Id          int       `gorm:"column:id" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	Path        string    `gorm:"column:path" json:"path"`
	Type        int       `gorm:"column:type" json:"type"`
	Status      int       `gorm:"column:status" json:"status"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy   string    `gorm:"column:created_by" json:"created_by"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy   string    `gorm:"column:updated_by" json:"updated-by"`
}

func (Resource) TableName() string {
	return "resource"
}

type Role struct {
	Id          int    `gorm:"column:id" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	//IsSystem    int       `gorm:"column:is_system" json:"is_system"`
	CreatedBy string    `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy string    `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Role) TableName() string {
	return "role"
}

type RoleResource struct {
	Id         int       `gorm:"column:id" json:"id"`
	RoleId     int       `gorm:"column:role_id" json:"role_id"`
	ResourceId int       `gorm:"column:resource_id" json:"resource_id"`
	Status     int       `gorm:"column:status" json:"status"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy  string    `gorm:"column:created_by" json:"created_by"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy  string    `gorm:"column:updated_by" json:"updated_by"`
}

func (RoleResource) TableName() string {
	return "role_resource"
}

type UserRole struct {
	Id        int       `gorm:"column:id" json:"id"`
	UserId    int       `gorm:"column:user_id" json:"user_id"`
	RoleId    int       `gorm:"column:role_id" json:"role_id"`
	Status    int       `gorm:"column:status" json:"status"`
	CreatedBy string    `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy string    `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (UserRole) TableName() string {
	return "user_role"
}
