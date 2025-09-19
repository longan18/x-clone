package business

import (
	"auth-service/internal/common"
	"auth-service/internal/module/permission/entity"
	"context"
)

type UpdatePermissionStorage interface {
	FirstPermissionByConditions(data map[string]interface{}) (*entity.Permission, error)
	UpdatePermission(ctx context.Context, id int, updates map[string]interface{}) error
	GetPermissionByIds(ids []int) ([]entity.Permission, error)
}

type updatePermissionBiz struct {
	biz              UpdatePermissionStorage
	getPermissionBiz *getPermissionBiz
}

func NewUpdatePermissionBiz(biz UpdatePermissionStorage) *updatePermissionBiz {
	return &updatePermissionBiz{
		biz:              biz,
		getPermissionBiz: NewGetPermissionBiz(biz),
	}
}

func (up *updatePermissionBiz) UpdatePermission(ctx context.Context, id int, permissionReq *entity.PermissionUpdateRequest) (*entity.Permission, error) {
	if err := permissionReq.CheckValidation(); err != nil {
		return nil, common.ErrBadRequest.WithError(err.Error())
	}

	// Check if permission exists
	existingPermission, err := up.getPermissionBiz.GetPermissionById(id)
	if err != nil {
		return nil, err
	}

	updates := make(map[string]interface{})

	// Check permission name uniqueness if being updated
	if permissionReq.PermissionName != "" && permissionReq.PermissionName != existingPermission.PermissionName {
		checkPermission, _ := up.getPermissionBiz.GetPermissionByName(permissionReq.PermissionName)
		if checkPermission != nil {
			return nil, common.ErrConflict.WithError(entity.ErrPermissionNameHasExisted.Error()).WithID(entity.ErrPermissionNameExisted)
		}
		updates["permission_name"] = permissionReq.PermissionName
	}

	// Update description if provided
	if permissionReq.Description != "" {
		updates["description"] = permissionReq.Description
	}

	if len(updates) > 0 {
		if err := up.biz.UpdatePermission(ctx, id, updates); err != nil {
			return nil, common.ErrInternalServerError.WithTrace(err).WithReason("Failed to update permission in database")
		}
	}

	// Return updated permission
	updatedPermission, err := up.getPermissionBiz.GetPermissionById(id)
	if err != nil {
		return nil, err
	}

	return updatedPermission, nil
}
