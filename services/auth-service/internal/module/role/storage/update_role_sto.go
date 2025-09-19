package storage

import (
	"auth-service/internal/module/role/entity"
	"context"
)

func (s *mysqlStorage) UpdateRole(ctx context.Context, id int, updates map[string]interface{}) error {
	return s.db.WithContext(ctx).Model(&entity.Role{}).Where("id = ?", id).Updates(updates).Error
}
