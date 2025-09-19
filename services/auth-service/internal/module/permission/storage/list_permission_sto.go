package storage

import (
	"auth-service/internal/module/permission/entity"
)

func (s *mysqlStorage) ListPermissions(limit, offset int) ([]entity.Permission, error) {
	var permissions []entity.Permission
	if err := s.db.Limit(limit).Offset(offset).Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func (s *mysqlStorage) CountPermissions() (int64, error) {
	var count int64
	if err := s.db.Model(&entity.Permission{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
