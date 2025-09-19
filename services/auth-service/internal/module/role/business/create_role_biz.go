package business

import (
	"auth-service/internal/common"
	"auth-service/internal/module/role/entity"
	"context"
)

type CreateRoleStorage interface {
	CreateRole(ctx context.Context, role *entity.Role) error
	FirstRoleByConditions(data map[string]interface{}) (*entity.Role, error)
	GetRoleByIds(ids []int) ([]entity.Role, error)
}

type createRoleBiz struct {
	biz        CreateRoleStorage
	getRoleBiz *getRoleBiz
}

func NewCreateRoleBiz(biz CreateRoleStorage) *createRoleBiz {
	return &createRoleBiz{
		biz:        biz,
		getRoleBiz: NewGetRoleBiz(biz),
	}
}

func (cr *createRoleBiz) CreateNewRole(ctx context.Context, roleReq *entity.RoleRequest) (*entity.Role, error) {
	if err := roleReq.CheckValidation(); err != nil {
		return nil, common.ErrBadRequest.WithError(err.Error())
	}

	// Check if role name already exists
	existingRole, _ := cr.getRoleBiz.GetRoleByName(roleReq.RoleName)
	if existingRole != nil {
		return nil, common.ErrConflict.WithError(entity.ErrRoleNameHasExisted.Error()).WithID(entity.ErrRoleNameExisted)
	}

	var role entity.Role
	role.RoleName = roleReq.RoleName
	role.Description = roleReq.Description

	if err := cr.biz.CreateRole(ctx, &role); err != nil {
		return nil, common.ErrInternalServerError.WithTrace(err).WithReason("Failed to create role in database")
	}

	return &role, nil
}
