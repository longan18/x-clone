package business

import (
	"auth-service/internal/common"
	"auth-service/internal/module/permission/entity"
	"context"
)

type DeletePermissionStorage interface {
	FirstPermissionByConditions(data map[string]interface{}) (*entity.Permission, error)
	DeletePermission(ctx context.Context, id int) error
	SoftDeletePermission(ctx context.Context, id int) error
	GetPermissionByIds(ids []int) ([]entity.Permission, error)
}

type deletePermissionBiz struct {
	biz              DeletePermissionStorage
	getPermissionBiz *getPermissionBiz
}

func NewDeletePermissionBiz(biz DeletePermissionStorage) *deletePermissionBiz {
	return &deletePermissionBiz{
		biz:              biz,
		getPermissionBiz: NewGetPermissionBiz(biz),
	}
}

func (dp *deletePermissionBiz) DeletePermission(ctx context.Context, id int) error {
	// Check if permission exists
	_, err := dp.getPermissionBiz.GetPermissionById(id)
	if err != nil {
		return err
	}

	if err := dp.biz.DeletePermission(ctx, id); err != nil {
		return common.ErrInternalServerError.WithTrace(err).WithReason("Failed to delete permission from database")
	}

	return nil
}

func (dp *deletePermissionBiz) SoftDeletePermission(ctx context.Context, id int) error {
	// Check if permission exists
	_, err := dp.getPermissionBiz.GetPermissionById(id)
	if err != nil {
		return err
	}

	if err := dp.biz.SoftDeletePermission(ctx, id); err != nil {
		return common.ErrInternalServerError.WithTrace(err).WithReason("Failed to soft delete permission from database")
	}

	return nil
}
