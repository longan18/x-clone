package business

import (
	"auth-service/internal/common"
	"auth-service/internal/module/permission/entity"
	"errors"
)

type GetPermissionStorage interface {
	FirstPermissionByConditions(data map[string]interface{}) (*entity.Permission, error)
	GetPermissionByIds(ids []int) ([]entity.Permission, error)
}

type getPermissionBiz struct {
	biz GetPermissionStorage
}

func NewGetPermissionBiz(biz GetPermissionStorage) *getPermissionBiz {
	return &getPermissionBiz{biz: biz}
}

func (gp *getPermissionBiz) GetPermissionById(id int) (*entity.Permission, error) {
	permission, err := gp.biz.FirstPermissionByConditions(map[string]interface{}{"id": id})

	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.ErrNotFound.WithTrace(err).WithReason("Permission not found")
		}
		return nil, common.ErrInternalServerError.WithTrace(err).WithReason("Failed to get permission from database")
	}

	permission.Mask(permission.Id)

	return permission, nil
}

func (gp *getPermissionBiz) GetPermissionByIds(ids []int) ([]entity.Permission, error) {
	permissions, err := gp.biz.GetPermissionByIds(ids)

	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.ErrNotFound.WithTrace(err).WithReason("Permission not found")
		}
		return nil, common.ErrInternalServerError.WithTrace(err).WithReason("Failed to get permissions from database")
	}

	return permissions, nil
}

func (gp *getPermissionBiz) GetPermissionByName(permissionName string) (*entity.Permission, error) {
	permission, err := gp.biz.FirstPermissionByConditions(map[string]interface{}{"permission_name": permissionName})

	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.ErrNotFound.WithTrace(err).WithReason("Permission with this name not found")
		}
		return nil, common.ErrInternalServerError.WithTrace(err).WithReason("Failed to get permission by name from database")
	}

	permission.Mask(permission.Id)

	return permission, nil
}
