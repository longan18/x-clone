package model

import (
	"auth-service/internal/common"
	"time"
)

type Permission struct {
	common.SQLModel
	PermissionName string     `json:"permission_name"`
	Description    string     `json:"description"`
	Roles          []Role     `gorm:"many2many:role_permissions" json:"roles"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

const (
	PremissionViewUser = iota + 1
	PremissionEditUser
	PremissionDeleteUser
	PremissionViewPost
	PremissionCreatePost
	PremissionEditPost
	PremissionDeletePost
)
