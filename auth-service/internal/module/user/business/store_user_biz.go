package business

import (
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
		return nil, util.NewDuplicateError(model.ErrEmailExistedMsg, model.ErrEmailExisted)
	}

	existingUser, _ = cu.getUserBiz.GetUserByUsername(userReq.UserName)
	if existingUser != nil {
		return nil, util.NewDuplicateError(model.ErrUsernameExistedMsg, model.ErrUsernameExisted)
	}

	hashed, err := util.HashPassword(strings.TrimSpace(userReq.Password))
	if err != nil {
		return nil, util.NewCreateError(err, model.ErrHashPassword)
	}

	var user model.User
	user.Email = userReq.Email
	user.UserName = userReq.UserName
	user.PasswordHash = hashed

	if err := cu.biz.CreateUser(ctx, &user); err != nil {
		return nil, util.NewCreateError(err, model.ErrCreateUser)
	}

	return &user, nil
}
