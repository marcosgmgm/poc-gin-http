package entity

import "github.com/marcosgmgm/poc-gin-http/api/response"

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u User) ToResponse() response.User {
	return response.User{
		Id:    u.Id,
		Name:  u.Name,
		Email: u.Email,
	}
}
