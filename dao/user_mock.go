package dao

import "github.com/marcosgmgm/poc-gin-http/api/entity"

type UserDaoCustomMock struct {
	SaveMock func(user *entity.User) error
	ByIdMock func(id string) (entity.User, error)
}

func (u UserDaoCustomMock) Save(user *entity.User) error {
	return u.SaveMock(user)
}

func (u UserDaoCustomMock) ById(id string) (entity.User, error) {
	return u.ByIdMock(id)
}

