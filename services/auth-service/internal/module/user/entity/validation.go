package entity

import (
	"net/mail"
	"strings"
)

func (req *UserRequest) CheckValidation() error {
	if !req.checkUserName() {
		return ErrUsernameIsNotValid
	}
	if !req.checkEmail() {
		return ErrEmailIsNotValid
	}
	if !req.checkPassword() {
		return ErrPasswordIsNotValid
	}

	return nil
}

func (req *UserRequest) checkEmail() bool {
	req.Email = strings.TrimSpace(req.Email)
	_, err := mail.ParseAddress(req.Email)
	return err == nil
}

func (req *UserRequest) checkPassword() bool {
	req.Password = strings.TrimSpace(req.Password)
	if len(req.Password) < 8 || len(req.Password) > 30 {
		return false
	}
	return true
}

func (req *UserRequest) checkUserName() bool {
	req.UserName = strings.TrimSpace(req.UserName)
	return len(req.UserName) >= 8
}
