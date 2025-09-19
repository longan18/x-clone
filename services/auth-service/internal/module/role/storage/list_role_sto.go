package storage

import (
	"auth-service/internal/module/role/entity"
)

func (s *mysqlStorage) ListRoles(limit, offset int) ([]entity.Role, error) {
	var roles []entity.Role
	if err := s.db.Limit(limit).Offset(offset).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (s *mysqlStorage) CountRoles() (int64, error) {
	var count int64
	if err := s.db.Model(&entity.Role{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
