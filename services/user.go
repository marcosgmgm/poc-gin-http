package services

import (
	"github.com/marcosgmgm/poc-gin-http/api/entity"
)

type User interface {
	Get(id string) (*entity.User, error)
	Save(user *entity.User) error
}
