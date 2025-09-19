package entity

import "strings"

func (req *PermissionRequest) CheckValidation() error {
	if !req.checkPermissionName() {
		return ErrPermissionNameIsNotValid
	}
	return nil
}

func (req *PermissionRequest) checkPermissionName() bool {
	req.PermissionName = strings.TrimSpace(req.PermissionName)
	return len(req.PermissionName) >= 3 && len(req.PermissionName) <= 50
}

func (req *PermissionUpdateRequest) CheckValidation() error {
	if req.PermissionName != "" && !req.checkPermissionName() {
		return ErrPermissionNameIsNotValid
	}
	return nil
}

func (req *PermissionUpdateRequest) checkPermissionName() bool {
	req.PermissionName = strings.TrimSpace(req.PermissionName)
	return len(req.PermissionName) >= 3 && len(req.PermissionName) <= 50
}
