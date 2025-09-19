package storage

import (
	"auth-service/internal/module/permission/entity"
	"context"
)

func (s *mysqlStorage) DeletePermission(ctx context.Context, id int) error {
	return s.db.WithContext(ctx).Delete(&entity.Permission{}, id).Error
}

func (s *mysqlStorage) SoftDeletePermission(ctx context.Context, id int) error {
	return s.db.WithContext(ctx).Model(&entity.Permission{}).Where("id = ?", id).Update("deleted_at", "NOW()").Error
}
