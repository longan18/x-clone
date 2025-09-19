package business

import (
	"auth-service/internal/common"
	"auth-service/internal/module/permission/entity"
	"context"
)

type CreatePermissionStorage interface {
	CreatePermission(ctx context.Context, permission *entity.Permission) error
	FirstPermissionByConditions(data map[string]interface{}) (*entity.Permission, error)
	GetPermissionByIds(ids []int) ([]entity.Permission, error)
}

type createPermissionBiz struct {
	biz              CreatePermissionStorage
	getPermissionBiz *getPermissionBiz
}

func NewCreatePermissionBiz(biz CreatePermissionStorage) *createPermissionBiz {
	return &createPermissionBiz{
		biz:              biz,
		getPermissionBiz: NewGetPermissionBiz(biz),
	}
}

func (cp *createPermissionBiz) CreateNewPermission(ctx context.Context, permissionReq *entity.PermissionRequest) (*entity.Permission, error) {
	if err := permissionReq.CheckValidation(); err != nil {
		return nil, common.ErrBadRequest.WithError(err.Error())
	}

	// Check if permission name already exists
	existingPermission, _ := cp.getPermissionBiz.GetPermissionByName(permissionReq.PermissionName)
	if existingPermission != nil {
		return nil, common.ErrConflict.WithError(entity.ErrPermissionNameHasExisted.Error()).WithID(entity.ErrPermissionNameExisted)
	}

	var permission entity.Permission
	permission.PermissionName = permissionReq.PermissionName
	permission.Description = permissionReq.Description

	if err := cp.biz.CreatePermission(ctx, &permission); err != nil {
		return nil, common.ErrInternalServerError.WithTrace(err).WithReason("Failed to create permission in database")
	}

	return &permission, nil
}
