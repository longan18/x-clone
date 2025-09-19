package storage

import (
	"auth-service/internal/module/permission/entity"
)

func (s *mysqlStorage) FirstPermissionByConditions(data map[string]interface{}) (*entity.Permission, error) {
	var permission *entity.Permission
	if err := s.db.Where(data).First(&permission).Error; err != nil {
		return nil, err
	}

	return permission, nil
}

func (s *mysqlStorage) GetPermissionByIds(ids []int) ([]entity.Permission, error) {
	var permissions []entity.Permission
	if err := s.db.Where("id IN ?", ids).Find(&permissions).Error; err != nil {
		return nil, err
	}

	return permissions, nil
}
