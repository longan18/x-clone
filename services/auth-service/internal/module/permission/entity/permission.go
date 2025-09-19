package entity

import (
	"auth-service/internal/common"
	"time"
)

type Permission struct {
	common.SQLModel
	PermissionName string     `json:"permission_name"`
	Description    string     `json:"description"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
}

type PermissionRequest struct {
	PermissionName string `json:"permission_name"`
	Description    string `json:"description"`
}

type PermissionUpdateRequest struct {
	PermissionName string `json:"permission_name,omitempty"`
	Description    string `json:"description,omitempty"`
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
