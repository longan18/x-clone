package storage

import (
	"auth-service/internal/module/role/entity"
	"context"
)

func (s *mysqlStorage) CreateRole(ctx context.Context, role *entity.Role) error {
	return s.db.WithContext(ctx).Create(role).Error
}
