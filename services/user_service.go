package services

import (
	"github.com/marcosgmgm/poc-gin-http/api/entity"
	"github.com/marcosgmgm/poc-gin-http/dao"
)

type userService struct{
	dao dao.User
}

func NewUserService(u dao.User) User {
	return userService{
		dao: u,
	}
}

func (u userService) Get(id string) (*entity.User, error) {
	user, err := u.dao.ById(id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userService) Save(user *entity.User) error {
	if err := u.dao.Save(user); err != nil {
		return err
	}
	return nil
}
