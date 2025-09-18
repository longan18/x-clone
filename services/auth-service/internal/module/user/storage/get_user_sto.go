package storage

import (
	"auth-service/internal/model"
)

func (s *mysqlStorage) FirstUserByConditions(data map[string]interface{}) (*model.User, error) {
	var user *model.User
	if err := s.db.Where(data).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}