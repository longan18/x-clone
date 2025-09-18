package business

import (
	"auth-service/internal/common"
	"auth-service/internal/model"
	"errors"
)

type GetUserStorage interface {
	FirstUserByConditions(data map[string]interface{}) (*model.User, error)
}

type getUserBiz struct {
	biz GetUserStorage
}

func NewGetUserBiz(biz GetUserStorage) *getUserBiz {
	return &getUserBiz{biz: biz}
}

func (gu *getUserBiz) GetUserById(id int) (*model.User, error) {
	user, err := gu.biz.FirstUserByConditions(map[string]interface{}{"id": id})

	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.ErrNotFound.WithTrace(err).WithReason("User not found")
		}
		return nil, common.ErrInternalServerError.WithTrace(err).WithReason("Failed to get user from database")
	}

	return user, nil
}

func (gu *getUserBiz) GetUserByEmail(email string) (*model.User, error) {
	user, err := gu.biz.FirstUserByConditions(map[string]interface{}{"email": email})

	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.ErrNotFound.WithTrace(err).WithReason("User with this email not found")
		}
		return nil, common.ErrInternalServerError.WithTrace(err).WithReason("Failed to get user by email from database")
	}

	return user, nil
}

func (gu *getUserBiz) GetUserByUsername(username string) (*model.User, error) {
	user, err := gu.biz.FirstUserByConditions(map[string]interface{}{"username": username})

	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.ErrNotFound.WithTrace(err).WithReason("User with this username not found")
		}
		return nil, common.ErrInternalServerError.WithTrace(err).WithReason("Failed to get user by username from database")
	}

	return user, nil
}
