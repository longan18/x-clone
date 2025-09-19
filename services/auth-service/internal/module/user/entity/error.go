package entity

import "errors"

const (
	ErrCreateUser      = "ErrCreateUser"
	ErrEmailExisted    = "ErrEmailExisted"
	ErrUsernameExisted = "ErrUsernameExisted"
	ErrHashPassword    = "ErrHashPassword"
	ErrRoleNotFound    = "ErrRoleNotFound"
)

var (
	ErrUsernameHasExisted = errors.New("username has existed")
	ErrUsernameIsNotValid = errors.New("username must have from 8 characters")
	ErrPasswordIsNotValid = errors.New("password must have from 8 to 30 characters")
	ErrEmailIsNotValid    = errors.New("email is not valid")
	ErrEmailHasExisted    = errors.New("email has existed")
	ErrLoginFailed        = errors.New("email and password are not valid")
	ErrCannotRegister     = errors.New("cannot register")
	ErrRoleNotExists      = errors.New("one or more roles do not exist")
)
