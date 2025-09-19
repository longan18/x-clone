package storage

import (
	roleEntity "auth-service/internal/module/role/entity"
	"auth-service/internal/module/user/entity"
	"context"
)

func (s *mysqlStorage) UpdateUser(ctx context.Context, id int, updates map[string]interface{}) error {
	return s.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", id).Updates(updates).Error
}

func (s *mysqlStorage) UpdateUserWithRoles(ctx context.Context, id int, updates map[string]interface{}, roles []roleEntity.Role) error {
	tx := s.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// Update user fields
	if len(updates) > 0 {
		if err := tx.Model(&entity.User{}).Where("id = ?", id).Updates(updates).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Update user roles
	var user entity.User
	if err := tx.Where("id = ?", id).First(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&user).Association("Roles").Replace(&roles); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
