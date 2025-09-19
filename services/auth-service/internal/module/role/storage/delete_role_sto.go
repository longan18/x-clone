package storage

import (
	"auth-service/internal/module/role/entity"
	"context"
)

func (s *mysqlStorage) DeleteRole(ctx context.Context, id int) error {
	return s.db.WithContext(ctx).Delete(&entity.Role{}, id).Error
}

func (s *mysqlStorage) SoftDeleteRole(ctx context.Context, id int) error {
	return s.db.WithContext(ctx).Model(&entity.Role{}).Where("id = ?", id).Update("deleted_at", "NOW()").Error
}
