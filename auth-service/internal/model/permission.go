package model

import "time"

type Permission struct {
	ID             uint         `json:"id"`
	PermissionName string       `json:"permission_name"`
	Description    string       `json:"description"`
	Roles          []Permission `gorm:"many2many:role_permissions" json:"roles"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	DeletedAt      *time.Time   `json:"deleted_at"`
}
