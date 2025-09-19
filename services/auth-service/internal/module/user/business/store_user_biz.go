package business

import (
	"auth-service/internal/common"
	"auth-service/internal/module/role/business"
	"auth-service/internal/module/user/entity"
	"auth-service/internal/util"
	"context"
	"strings"
)

type StoreUserStorage interface {
	CreateUser(ctx context.Context, user *entity.User) error
	FirstUserByConditions(data map[string]interface{}) (*entity.User, error)
}

type createUserBiz struct {
	biz        StoreUserStorage
	getUserBiz *getUserBiz
	getRoleBiz business.GetRoleStorage
}

func NewCreateUserBiz(biz StoreUserStorage, bizRole business.GetRoleStorage) *createUserBiz {
	return &createUserBiz{
		biz:        biz,
		getUserBiz: NewGetUserBiz(biz),
		getRoleBiz: bizRole,
	}
}

func (cu *createUserBiz) CreateNewUser(ctx context.Context, userReq *entity.UserRequest) (*entity.User, error) {
	if err := userReq.CheckValidation(); err != nil {
		return nil, common.ErrBadRequest.WithError(err.Error())
	}

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

	roles, err := cu.getRoleBiz.GetRoleByIds(userReq.Roles)
	if err != nil {
		return nil, common.ErrInternalServerError.WithTrace(err).WithID(entity.ErrCreateUser).WithReason("Failed to get role in database")
	}
	if len(roles) == 0 {
		return nil, common.ErrBadRequest.WithError(entity.ErrRoleNotExists.Error()).WithID(entity.ErrRoleNotFound)
	}

	var user entity.User
	user.Email = userReq.Email
	user.UserName = userReq.UserName
	user.PasswordHash = hashed
	user.Roles = roles

	if err := cu.biz.CreateUser(ctx, &user); err != nil {
		return nil, common.ErrInternalServerError.WithTrace(err).WithID(entity.ErrCreateUser).WithReason("Failed to create user in database")
	}

	return &user, nil
}
