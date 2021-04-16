package services

import "github.com/marcosgmgm/poc-gin-http/api/entity"

type UserServiceCustomMock struct {
	GetMock  func(id string) (*entity.User, error)
	SaveMock func(user *entity.User) error
}

func (u UserServiceCustomMock) Get(id string) (*entity.User, error) {
	return u.GetMock(id)
}

func (u UserServiceCustomMock) Save(user *entity.User) error {
	return u.SaveMock(user)
}
