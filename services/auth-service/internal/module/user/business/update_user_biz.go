package business

import (
	"auth-service/internal/common"
	"auth-service/internal/module/role/business"
	roleEntity "auth-service/internal/module/role/entity"
	"auth-service/internal/module/user/entity"
	"auth-service/internal/util"
	"context"
	"strings"
)

type UpdateUserStorage interface {
	FirstUserByConditions(data map[string]interface{}) (*entity.User, error)
	UpdateUser(ctx context.Context, id int, updates map[string]interface{}) error
	UpdateUserWithRoles(ctx context.Context, id int, updates map[string]interface{}, roles []roleEntity.Role) error
}

type updateUserBiz struct {
	biz        UpdateUserStorage
	getUserBiz *getUserBiz
	getRoleBiz business.GetRoleStorage
}

func NewUpdateUserBiz(biz UpdateUserStorage, bizRole business.GetRoleStorage) *updateUserBiz {
	return &updateUserBiz{
		biz:        biz,
		getUserBiz: NewGetUserBiz(biz),
		getRoleBiz: bizRole,
	}
}

func (ub *updateUserBiz) UpdateUser(ctx context.Context, id int, userReq *entity.UserUpdateRequest) (*entity.User, error) {
	if err := userReq.CheckValidation(); err != nil {
		return nil, common.ErrBadRequest.WithError(err.Error())
	}

	// Check if user exists
	existingUser, err := ub.getUserBiz.GetUserById(id)
	if err != nil {
		return nil, err
	}

	updates := make(map[string]interface{})

	// Check username uniqueness if being updated
	if userReq.UserName != "" && userReq.UserName != existingUser.UserName {
		checkUser, _ := ub.getUserBiz.GetUserByUsername(userReq.UserName)
		if checkUser != nil {
			return nil, common.ErrConflict.WithError(entity.ErrUsernameHasExisted.Error()).WithID(entity.ErrUsernameExisted)
		}
		updates["username"] = userReq.UserName
	}

	// Check email uniqueness if being updated
	if userReq.Email != "" && userReq.Email != existingUser.Email {
		checkUser, _ := ub.getUserBiz.GetUserByEmail(userReq.Email)
		if checkUser != nil {
			return nil, common.ErrConflict.WithError(entity.ErrEmailHasExisted.Error()).WithID(entity.ErrEmailExisted)
		}
		updates["email"] = userReq.Email
	}

	// Hash password if being updated
	if userReq.Password != "" {
		hashed, err := util.HashPassword(strings.TrimSpace(userReq.Password))
		if err != nil {
			return nil, common.ErrInternalServerError.WithTrace(err).WithID(entity.ErrHashPassword).WithReason("Failed to hash password")
		}
		updates["password_hash"] = hashed
	}

	// Handle roles update
	if len(userReq.Roles) > 0 {
		roles, err := ub.getRoleBiz.GetRoleByIds(userReq.Roles)
		if err != nil {
			return nil, common.ErrInternalServerError.WithTrace(err).WithID(entity.ErrCreateUser).WithReason("Failed to get roles from database")
		}
		if len(roles) == 0 {
			return nil, common.ErrBadRequest.WithError(entity.ErrRoleNotExists.Error()).WithID(entity.ErrRoleNotFound)
		}

		if err := ub.biz.UpdateUserWithRoles(ctx, id, updates, roles); err != nil {
			return nil, common.ErrInternalServerError.WithTrace(err).WithReason("Failed to update user in database")
		}
	} else if len(updates) > 0 {
		if err := ub.biz.UpdateUser(ctx, id, updates); err != nil {
			return nil, common.ErrInternalServerError.WithTrace(err).WithReason("Failed to update user in database")
		}
	}

	// Return updated user
	updatedUser, err := ub.getUserBiz.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
