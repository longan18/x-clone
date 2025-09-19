package business

import (
	"auth-service/internal/common"
	"auth-service/internal/module/user/entity"
	"context"
)

type DeleteUserStorage interface {
	FirstUserByConditions(data map[string]interface{}) (*entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	SoftDeleteUser(ctx context.Context, id int) error
}

type deleteUserBiz struct {
	biz        DeleteUserStorage
	getUserBiz *getUserBiz
}

func NewDeleteUserBiz(biz DeleteUserStorage) *deleteUserBiz {
	return &deleteUserBiz{
		biz:        biz,
		getUserBiz: NewGetUserBiz(biz),
	}
}

func (db *deleteUserBiz) DeleteUser(ctx context.Context, id int) error {
	// Check if user exists
	_, err := db.getUserBiz.GetUserById(id)
	if err != nil {
		return err
	}

	if err := db.biz.DeleteUser(ctx, id); err != nil {
		return common.ErrInternalServerError.WithTrace(err).WithReason("Failed to delete user from database")
	}

	return nil
}

func (db *deleteUserBiz) SoftDeleteUser(ctx context.Context, id int) error {
	// Check if user exists
	_, err := db.getUserBiz.GetUserById(id)
	if err != nil {
		return err
	}

	if err := db.biz.SoftDeleteUser(ctx, id); err != nil {
		return common.ErrInternalServerError.WithTrace(err).WithReason("Failed to soft delete user from database")
	}

	return nil
}
