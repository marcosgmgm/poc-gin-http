package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/marcosgmgm/poc-gin-http/api/entity"
	"github.com/marcosgmgm/poc-gin-http/dao"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveSuccess(t *testing.T) {
	id := uuid.New().String()
	service := NewUserService(dao.UserDaoCustomMock{
		SaveMock: func(user *entity.User) error {
			user.Id = id
			return nil
		},
	})
	userToSave := entity.User{
		Name:  "Guima",
		Email: "guima@guima.com",
	}
	err := service.Save(&userToSave)
	assert.Nil(t, err)
	assert.Equal(t, id, userToSave.Id)
}

func TestSaveError(t *testing.T) {
	wantErr := errors.New("error mock")
	service := NewUserService(dao.UserDaoCustomMock{
		SaveMock: func(user *entity.User) error {
			return wantErr
		},
	})
	userToSave := entity.User{
		Name:  "Guima",
		Email: "guima@guima.com",
	}
	err := service.Save(&userToSave)
	assert.Equal(t, wantErr, err)
	assert.Empty(t, userToSave.Id)
}

func TestGetSuccess(t *testing.T) {
	mockId := uuid.New().String()
	service := NewUserService(dao.UserDaoCustomMock{
		ByIdMock: func(id string) (entity.User, error) {
			return entity.User{
				Id:    mockId,
				Name:  "Guima",
				Email: "guima@guima.com",
			}, nil
		},
	})
	userFind, err := service.Get(mockId)
	assert.Nil(t, err)
	assert.Equal(t, mockId, userFind.Id)
}

func TestGetError(t *testing.T) {
	wantErr := errors.New("error mock")
	service := NewUserService(dao.UserDaoCustomMock{
		ByIdMock: func(id string) (entity.User, error) {
			return entity.User{}, wantErr
		},
	})
	userFind, err := service.Get(uuid.New().String())
	assert.Equal(t, wantErr, err)
	assert.Nil(t, userFind)
}
