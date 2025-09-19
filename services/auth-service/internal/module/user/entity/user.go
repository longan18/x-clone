package entity

import (
	"auth-service/internal/common"
	roleEntity "auth-service/internal/module/role/entity"
	"time"
)

type User struct {
	common.SQLModel
	Roles        []roleEntity.Role `gorm:"many2many:user_roles" json:"roles"`
	UserName     string            `gorm:"column:username" json:"username"`
	Email        string            `gorm:"column:email" json:"email"`
	PasswordHash string            `gorm:"column:password_hash" json:"password_hash"`
	DeletedAt    *time.Time        `gorm:"column:deleted_at" json:"deleted_at"`
}

type UserRequest struct {
	Roles    []int `json:"roles"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
