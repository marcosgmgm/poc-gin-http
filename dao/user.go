package dao

import (
	"errors"
	"github.com/marcosgmgm/poc-gin-http/api/entity"
)

var (
	InvalidUserErr = errors.New("invalid user to save")
	NotFoundErr = errors.New("user not found")
)

type User interface {
	Save(user *entity.User) error
	ById(id string) (entity.User, error)
}