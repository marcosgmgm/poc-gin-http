package dao

import (
	"github.com/google/uuid"
	"github.com/marcosgmgm/poc-gin-http/api/entity"
)

var (
	base = make(map[string]entity.User)
)

type userDao struct {}

func NewUserDao() User {
	return userDao{}
}

func (u userDao) Save(user *entity.User) error {
	if user == nil {
		return InvalidUserErr
	}
	user.Id = uuid.New().String()
	base[user.Id] = *user
	return nil
}

func (u userDao) ById(id string) (entity.User, error) {
	userBase := base[id]
	if len(userBase.Id) == 0 {
		return userBase, NotFoundErr
	}
	return userBase, nil
}
