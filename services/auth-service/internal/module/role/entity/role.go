package entity

import (
	"auth-service/internal/common"

	"gorm.io/gorm"
)

type Role struct {
	common.SQLModel
	RoleName    string          `json:"role_name"`
	Description string          `json:"description"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

const (
	RoleAdmin = iota + 1
	RoleUser
	RoleGuest
)
