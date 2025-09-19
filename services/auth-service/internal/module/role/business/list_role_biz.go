package business

import (
	"auth-service/internal/common"
	"auth-service/internal/module/role/entity"
)

type ListRoleStorage interface {
	ListRoles(limit, offset int) ([]entity.Role, error)
	CountRoles() (int64, error)
}

type listRoleBiz struct {
	biz ListRoleStorage
}

func NewListRoleBiz(biz ListRoleStorage) *listRoleBiz {
	return &listRoleBiz{biz: biz}
}

func (lb *listRoleBiz) ListRoles(limit, offset int) ([]entity.Role, int64, error) {
	roles, err := lb.biz.ListRoles(limit, offset)
	if err != nil {
		return nil, 0, common.ErrInternalServerError.WithTrace(err).WithReason("Failed to get roles from database")
	}

	count, err := lb.biz.CountRoles()
	if err != nil {
		return nil, 0, common.ErrInternalServerError.WithTrace(err).WithReason("Failed to count roles from database")
	}

	// Mask all roles
	for i, role := range roles {
		role.Mask(role.Id)
		roles[i] = role
	}

	return roles, count, nil
}
