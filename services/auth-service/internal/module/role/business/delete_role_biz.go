package business

import (
	"auth-service/internal/common"
	"auth-service/internal/module/role/entity"
	"context"
)

type DeleteRoleStorage interface {
	FirstRoleByConditions(data map[string]interface{}) (*entity.Role, error)
	DeleteRole(ctx context.Context, id int) error
	SoftDeleteRole(ctx context.Context, id int) error
	GetRoleByIds(ids []int) ([]entity.Role, error)
}

type deleteRoleBiz struct {
	biz        DeleteRoleStorage
	getRoleBiz *getRoleBiz
}

func NewDeleteRoleBiz(biz DeleteRoleStorage) *deleteRoleBiz {
	return &deleteRoleBiz{
		biz:        biz,
		getRoleBiz: NewGetRoleBiz(biz),
	}
}

func (db *deleteRoleBiz) DeleteRole(ctx context.Context, id int) error {
	// Check if role exists
	_, err := db.getRoleBiz.GetRoleById(id)
	if err != nil {
		return err
	}

	if err := db.biz.DeleteRole(ctx, id); err != nil {
		return common.ErrInternalServerError.WithTrace(err).WithReason("Failed to delete role from database")
	}

	return nil
}

func (db *deleteRoleBiz) SoftDeleteRole(ctx context.Context, id int) error {
	// Check if role exists
	_, err := db.getRoleBiz.GetRoleById(id)
	if err != nil {
		return err
	}

	if err := db.biz.SoftDeleteRole(ctx, id); err != nil {
		return common.ErrInternalServerError.WithTrace(err).WithReason("Failed to soft delete role from database")
	}

	return nil
}
