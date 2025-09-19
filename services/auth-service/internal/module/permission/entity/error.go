package entity

import "errors"

var (
	ErrPermissionNameIsNotValid = errors.New("permission name is not valid")
	ErrPermissionNameHasExisted = errors.New("permission name has existed")
	ErrPermissionNotExists      = errors.New("permission not exists")
	ErrCreatePermission         = errors.New("create permission failed")
	ErrUpdatePermission         = errors.New("update permission failed")
	ErrDeletePermission         = errors.New("delete permission failed")
)

const (
	ErrPermissionNameExisted = "permission_name_existed"
	ErrPermissionNotFound    = "permission_not_found"
	ErrPermissionInvalid     = "permission_invalid"
)
