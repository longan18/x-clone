package entity

import "strings"

func (req *RoleRequest) CheckValidation() error {
	if !req.checkRoleName() {
		return ErrRoleNameIsNotValid
	}
	return nil
}

func (req *RoleRequest) checkRoleName() bool {
	req.RoleName = strings.TrimSpace(req.RoleName)
	return len(req.RoleName) >= 3 && len(req.RoleName) <= 50
}

func (req *RoleUpdateRequest) CheckValidation() error {
	if req.RoleName != "" && !req.checkRoleName() {
		return ErrRoleNameIsNotValid
	}
	return nil
}

func (req *RoleUpdateRequest) checkRoleName() bool {
	req.RoleName = strings.TrimSpace(req.RoleName)
	return len(req.RoleName) >= 3 && len(req.RoleName) <= 50
}
