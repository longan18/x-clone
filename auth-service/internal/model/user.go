package model

import "time"

type User struct {
	ID           uint       `json:"id"`
	Roles        []Role     `gorm:"many2many:user_roles" json:"roles"`
	UserName     string     `json:"username"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"password_hash"`
	Salt         string     `json:"salt"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

type UserRequest struct {
	Roles    []uint `json:"roles"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
