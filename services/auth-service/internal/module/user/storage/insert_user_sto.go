package storage

import (
	"auth-service/internal/module/user/entity"
	"context"
)

func (s *mysqlStorage) CreateUser(ctx context.Context, user *entity.User) error {
	if err := s.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}
