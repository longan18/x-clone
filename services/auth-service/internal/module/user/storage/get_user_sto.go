package storage

import (
	"auth-service/internal/module/user/entity"

	"gorm.io/gorm"
)

func (s *mysqlStorage) FirstUserByConditions(data map[string]interface{}) (*entity.User, error) {
	var user *entity.User
	if err := s.db.Preload("Roles", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "role_name", "description")
	}).Where(data).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
