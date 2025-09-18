package business

import (
	"auth-service/internal/common"
	"auth-service/internal/model"
	"auth-service/internal/util"
	"context"
	"strings"
)

type StoreUserStorage interface {
	CreateUser(ctx context.Context, user *model.User) error
	FirstUserByConditions(data map[string]interface{}) (*model.User, error)
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

func (cu *createUserBiz) CreateNewUser(ctx context.Context, userReq model.UserRequest) (*model.User, error) {
	existingUser, _ := cu.getUserBiz.GetUserByEmail(userReq.Email)
	if existingUser != nil {
		return nil, common.ErrConflict.WithError(model.ErrEmailExistedMsg).WithID(model.ErrEmailExisted)
	}

	existingUser, _ = cu.getUserBiz.GetUserByUsername(userReq.UserName)
	if existingUser != nil {
		return nil, common.ErrConflict.WithError(model.ErrUsernameExistedMsg).WithID(model.ErrUsernameExisted)
	}

	hashed, err := util.HashPassword(strings.TrimSpace(userReq.Password))
	if err != nil {
		return nil, common.ErrInternalServerError.WithTrace(err).WithID(model.ErrHashPassword).WithReason("Failed to hash password")
	}

	var user model.User
	user.Email = userReq.Email
	user.UserName = userReq.UserName
	user.PasswordHash = hashed

	if err := cu.biz.CreateUser(ctx, &user); err != nil {
		return nil, common.ErrInternalServerError.WithTrace(err).WithID(model.ErrCreateUser).WithReason("Failed to create user in database")
	}

	return &user, nil
}
