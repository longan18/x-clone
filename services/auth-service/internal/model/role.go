package model

import (
	"auth-service/internal/common"
	"time"
)

type Role struct {
	common.SQLModel
	RoleName    string       `json:"role_name"`
	Description string       `json:"description"`
	Permissions []Permission `gorm:"many2many:role_permissions" json:"permissions"`
	DeletedAt   *time.Time   `json:"deleted_at"`
}

const (
	RoleAdmin = iota + 1
	RoleUser
	RoleGuest
)