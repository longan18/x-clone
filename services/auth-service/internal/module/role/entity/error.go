package entity

import "errors"

var (
	ErrRoleNameIsNotValid = errors.New("role name is not valid")
	ErrRoleNameHasExisted = errors.New("role name has existed")
	ErrRoleNotExists      = errors.New("role not exists")
	ErrCreateRole         = errors.New("create role failed")
	ErrUpdateRole         = errors.New("update role failed")
	ErrDeleteRole         = errors.New("delete role failed")
)

const (
	ErrRoleNameExisted = "role_name_existed"
	ErrRoleNotFound    = "role_not_found"
	ErrRoleInvalid     = "role_invalid"
)
