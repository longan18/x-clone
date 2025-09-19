package business

import (
	"auth-service/internal/common"
	roleEntity "auth-service/internal/module/role/entity"
	"auth-service/internal/module/user/entity"
	"auth-service/internal/util"
	"context"
	"strings"
)

type StoreUserStorage interface {
	CreateUser(ctx context.Context, user *entity.User) error
	FirstUserByConditions(data map[string]interface{}) (*entity.User, error)
}

type GetRoleStorage interface {
	GetRoleByIds(ids []int) (*roleEntity.Role, error)
}

type createUserBiz struct {
	biz        StoreUserStorage
	getUserBiz *getUserBiz
}

func NewCreateUserBiz(biz StoreUserStorage) *createUserBiz {
	return &createUserBiz{
		biz:        biz,
		getUserBiz: NewGetUserBiz(biz),
	}
}

func (cu *createUserBiz) CreateNewUser(ctx context.Context, userReq *entity.UserRequest) (*entity.User, error) {
	existingUser, _ := cu.getUserBiz.GetUserByUsername(userReq.UserName)
	if existingUser != nil {
		return nil, common.ErrConflict.WithError(entity.ErrUsernameHasExisted.Error()).WithID(entity.ErrUsernameExisted)
	}

	existingUser, _ = cu.getUserBiz.GetUserByEmail(userReq.Email)
	if existingUser != nil {
		return nil, common.ErrConflict.WithError(entity.ErrEmailHasExisted.Error()).WithID(entity.ErrEmailExisted)
	}

	hashed, err := util.HashPassword(strings.TrimSpace(userReq.Password))
	if err != nil {
		return nil, common.ErrInternalServerError.WithTrace(err).WithID(entity.ErrHashPassword).WithReason("Failed to hash password")
	}

	var user entity.User
	user.Email = userReq.Email
	user.UserName = userReq.UserName
	user.PasswordHash = hashed

	if err := cu.biz.CreateUser(ctx, &user); err != nil {
		return nil, common.ErrInternalServerError.WithTrace(err).WithID(entity.ErrCreateUser).WithReason("Failed to create user in database")
	}

	return &user, nil
}
