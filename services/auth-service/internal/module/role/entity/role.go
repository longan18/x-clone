package entity

import (
	"auth-service/internal/common"
	"time"
)

type Role struct {
	common.SQLModel
	RoleName    string `json:"role_name"`
	Description string `json:"description"`
	DeletedAt *time.Time `json:"deleted_at"`
}

const (
	RoleAdmin = iota + 1
	RoleUser
	RoleGuest
)
