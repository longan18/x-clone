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

type RoleRequest struct {
	RoleName    string `json:"role_name"`
	Description string `json:"description"`
}

type RoleUpdateRequest struct {
	RoleName    string `json:"role_name,omitempty"`
	Description string `json:"description,omitempty"`
}

const (
	RoleAdmin = iota + 1
	RoleUser
	RoleGuest
)
