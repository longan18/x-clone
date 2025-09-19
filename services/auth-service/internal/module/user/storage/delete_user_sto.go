package storage

import (
	"auth-service/internal/module/user/entity"
	"context"
)

func (s *mysqlStorage) DeleteUser(ctx context.Context, id int) error {
	return s.db.WithContext(ctx).Delete(&entity.User{}, id).Error
}

func (s *mysqlStorage) SoftDeleteUser(ctx context.Context, id int) error {
	return s.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", id).Update("deleted_at", "NOW()").Error
}
