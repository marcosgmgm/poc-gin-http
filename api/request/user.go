package request

import (
	"github.com/marcosgmgm/poc-gin-http/api/entity"
)

type CreateUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (e CreateUser) GenerateEntity() entity.User {
	return entity.User{
		Name:  e.Name,
		Email: e.Email,
	}
}
