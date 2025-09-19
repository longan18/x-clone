package storage

import (
	"auth-service/internal/module/user/entity"
)

func (s *mysqlStorage) FirstUserByConditions(data map[string]interface{}) (*entity.User, error) {
	var user *entity.User
	if err := s.db.Where(data).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
