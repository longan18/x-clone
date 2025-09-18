package storage

import (
	"auth-service/internal/model"
	"context"
)

func (s *mysqlStorage) CreateUser(ctx context.Context, user *model.User) error {
	if err := s.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}