package business

import "auth-service/internal/model"

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
		return nil, err
	}

	return user, nil
}

func (gu *getUserBiz) GetUserByEmail(email string) (*model.User, error) {
	user, err := gu.biz.FirstUserByConditions(map[string]interface{}{"email": email})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (gu *getUserBiz) GetUserByUsername(username string) (*model.User, error) {
	user, err := gu.biz.FirstUserByConditions(map[string]interface{}{"username": username})

	if err != nil {
		return nil, err
	}

	return user, nil
}
