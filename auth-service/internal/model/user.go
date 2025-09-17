package model

import "time"

type User struct {
	ID           uint       `gorm:"column:id" json:"id"`
	Roles        []Role     `gorm:"many2many:user_roles" json:"roles"`
	UserName     string     `gorm:"column:username" json:"username"`
	Email        string     `gorm:"column:email" json:"email"`
	PasswordHash string     `gorm:"column:password_hash" json:"password_hash"`
	CreatedAt    time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

type UserRequest struct {
	Roles    []uint `json:"roles"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

const (
	ErrCreateUser      = "ErrCreateUser"
	ErrEmailExisted    = "ErrEmailExisted"
	ErrUsernameExisted = "ErrUsernameExisted"
	ErrHashPassword    = "ErrHashPassword"
)

const (
	ErrEmailExistedMsg    = "Email already exists, please check again"
	ErrUsernameExistedMsg = "User name already exists, please check again"
)
