package storage

import (
	"auth-service/internal/module/permission/entity"
	"context"
)

func (s *mysqlStorage) CreatePermission(ctx context.Context, permission *entity.Permission) error {
	return s.db.WithContext(ctx).Create(permission).Error
}
