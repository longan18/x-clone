package storage

import (
	"auth-service/internal/module/permission/entity"
	"context"
)

func (s *mysqlStorage) UpdatePermission(ctx context.Context, id int, updates map[string]interface{}) error {
	return s.db.WithContext(ctx).Model(&entity.Permission{}).Where("id = ?", id).Updates(updates).Error
}
