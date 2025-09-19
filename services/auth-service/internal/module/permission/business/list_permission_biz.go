package business

import (
	"auth-service/internal/common"
	"auth-service/internal/module/permission/entity"
)

type ListPermissionStorage interface {
	ListPermissions(limit, offset int) ([]entity.Permission, error)
	CountPermissions() (int64, error)
}

type listPermissionBiz struct {
	biz ListPermissionStorage
}

func NewListPermissionBiz(biz ListPermissionStorage) *listPermissionBiz {
	return &listPermissionBiz{biz: biz}
}

func (lp *listPermissionBiz) ListPermissions(limit, offset int) ([]entity.Permission, int64, error) {
	permissions, err := lp.biz.ListPermissions(limit, offset)
	if err != nil {
		return nil, 0, common.ErrInternalServerError.WithTrace(err).WithReason("Failed to get permissions from database")
	}

	count, err := lp.biz.CountPermissions()
	if err != nil {
		return nil, 0, common.ErrInternalServerError.WithTrace(err).WithReason("Failed to count permissions from database")
	}

	// Mask all permissions
	for i, permission := range permissions {
		permission.Mask(permission.Id)
		permissions[i] = permission
	}

	return permissions, count, nil
}
