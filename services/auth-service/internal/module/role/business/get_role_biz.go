package business

import (
	"auth-service/internal/common"
	"auth-service/internal/module/role/entity"
	"errors"
)

type GetRoleStorage interface {
	FirstRoleByConditions(data map[string]interface{}) (*entity.Role, error)
	GetRoleByIds(ids []int) ([]entity.Role, error)
}

type getRoleBiz struct {
	biz GetRoleStorage
}

func NewGetRoleBiz(biz GetRoleStorage) *getRoleBiz {
	return &getRoleBiz{biz: biz}
}

func (gu *getRoleBiz) GetRoleById(id int) (*entity.Role, error) {
	role, err := gu.biz.FirstRoleByConditions(map[string]interface{}{"id": id})

	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.ErrNotFound.WithTrace(err).WithReason("Role not found")
		}
		return nil, common.ErrInternalServerError.WithTrace(err).WithReason("Failed to get user from database")
	}

	return role, nil
}

func (gu *getRoleBiz) GetRoleByIds(ids []int) ([]entity.Role, error) {
	roles, err := gu.biz.GetRoleByIds(ids)

	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.ErrNotFound.WithTrace(err).WithReason("Role not found")
		}
		return nil, common.ErrInternalServerError.WithTrace(err).WithReason("Failed to get roles from database")
	}

	return roles, nil
}

func (gu *getRoleBiz) GetRoleByName(roleName string) (*entity.Role, error) {
	role, err := gu.biz.FirstRoleByConditions(map[string]interface{}{"role_name": roleName})

	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.ErrNotFound.WithTrace(err).WithReason("Role with this name not found")
		}
		return nil, common.ErrInternalServerError.WithTrace(err).WithReason("Failed to get role by name from database")
	}

	role.Mask(role.Id)

	return role, nil
}
