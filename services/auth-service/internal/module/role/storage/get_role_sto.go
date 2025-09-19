package storage

import (
	"auth-service/internal/module/role/entity"
)

func (s *mysqlStorage) FirstRoleByConditions(data map[string]interface{}) (*entity.Role, error) {
	var user *entity.Role
	if err := s.db.Where(data).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *mysqlStorage) GetRoleByIds(ids []int) (*entity.Role, error) {
	var user *entity.Role
	if err := s.db.Where("id IN ?", ids).Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

