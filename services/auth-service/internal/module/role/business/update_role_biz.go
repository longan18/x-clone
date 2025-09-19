package business

import (
	"auth-service/internal/common"
	"auth-service/internal/module/role/entity"
	"context"
)

type UpdateRoleStorage interface {
	FirstRoleByConditions(data map[string]interface{}) (*entity.Role, error)
	UpdateRole(ctx context.Context, id int, updates map[string]interface{}) error
	GetRoleByIds(ids []int) ([]entity.Role, error)
}

type updateRoleBiz struct {
	biz        UpdateRoleStorage
	getRoleBiz *getRoleBiz
}

func NewUpdateRoleBiz(biz UpdateRoleStorage) *updateRoleBiz {
	return &updateRoleBiz{
		biz:        biz,
		getRoleBiz: NewGetRoleBiz(biz),
	}
}

func (ub *updateRoleBiz) UpdateRole(ctx context.Context, id int, roleReq *entity.RoleUpdateRequest) (*entity.Role, error) {
	if err := roleReq.CheckValidation(); err != nil {
		return nil, common.ErrBadRequest.WithError(err.Error())
	}

	// Check if role exists
	existingRole, err := ub.getRoleBiz.GetRoleById(id)
	if err != nil {
		return nil, err
	}

	updates := make(map[string]interface{})

	// Check role name uniqueness if being updated
	if roleReq.RoleName != "" && roleReq.RoleName != existingRole.RoleName {
		checkRole, _ := ub.getRoleBiz.GetRoleByName(roleReq.RoleName)
		if checkRole != nil {
			return nil, common.ErrConflict.WithError(entity.ErrRoleNameHasExisted.Error()).WithID(entity.ErrRoleNameExisted)
		}
		updates["role_name"] = roleReq.RoleName
	}

	// Update description if provided
	if roleReq.Description != "" {
		updates["description"] = roleReq.Description
	}

	if len(updates) > 0 {
		if err := ub.biz.UpdateRole(ctx, id, updates); err != nil {
			return nil, common.ErrInternalServerError.WithTrace(err).WithReason("Failed to update role in database")
		}
	}

	// Return updated role
	updatedRole, err := ub.getRoleBiz.GetRoleById(id)
	if err != nil {
		return nil, err
	}

	return updatedRole, nil
}
