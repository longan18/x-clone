package entity

import "errors"

const (
	ErrCreateUser      = "ErrCreateUser"
	ErrEmailExisted    = "ErrEmailExisted"
	ErrUsernameExisted = "ErrUsernameExisted"
	ErrHashPassword    = "ErrHashPassword"
)

var (
	ErrUsernameHasExisted = errors.New("Username has existed")
	ErrPasswordIsNotValid = errors.New("Password must have from 8 to 30 characters")
	ErrEmailIsNotValid    = errors.New("Email is not valid")
	ErrEmailHasExisted    = errors.New("Email has existed")
	ErrLoginFailed        = errors.New("Email and password are not valid")
	ErrCannotRegister     = errors.New("Cannot register")
)