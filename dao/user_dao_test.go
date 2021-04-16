package dao

import (
	"github.com/marcosgmgm/poc-gin-http/api/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveSuccess(t *testing.T) {
	dao := NewUserDao()
	userToSave := entity.User{
		Name:  "Guima",
		Email: "guima@guima.com",
	}
	err := dao.Save(&userToSave)
	assert.Nil(t, err)
	assert.NotEmpty(t, userToSave.Id)
}

func TestSaveNil(t *testing.T) {
	dao := NewUserDao()
	err := dao.Save(nil)
	assert.Equal(t, InvalidUserErr, err)
}

func TestByIdSuccess(t *testing.T) {
	dao := NewUserDao()
	userToSave := entity.User{
		Name:  "Guima",
		Email: "guima@guima.com",
	}
	dao.Save(&userToSave)
	userFind, err := dao.ById(userToSave.Id)
	assert.Nil(t, err)
	assert.Equal(t, userToSave.Id, userFind.Id)
}

func TestByIdNotFound(t *testing.T) {
	dao := NewUserDao()
	userFind, err := dao.ById("not")
	assert.Equal(t, NotFoundErr, err)
	assert.Empty(t, userFind.Id)
}

