package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/marcosgmgm/poc-gin-http/api/entity"
	"github.com/marcosgmgm/poc-gin-http/api/request"
	"github.com/marcosgmgm/poc-gin-http/api/response"
	"github.com/marcosgmgm/poc-gin-http/middleware"
	"github.com/marcosgmgm/poc-gin-http/services"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSuccess(t *testing.T) {
	controller := NewUserController(services.UserServiceCustomMock{
		GetMock: func(id string) (*entity.User, error) {
			return &entity.User{
				Id:    id,
				Name:  "Guima",
				Email: "guima@guima.com",
			}, nil
		},
	})
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.Organization.Handler)
	router.GET("/users/:id", controller.Get)
	r, _ := http.NewRequest(http.MethodGet, "/users/12345678", nil)
	r.Header.Add(middleware.X_ORG_CTX_KEY, "guima")
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	var got response.User
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	want := response.User{
		Id:    "12345678",
		Name:  "Guima",
		Email: "guima@guima.com",
	}
	assert.Equal(t, want, got)
}

func TestGetBadRequest(t *testing.T) {
	controller := NewUserController(services.UserServiceCustomMock{})
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.Organization.Handler)
	router.GET("/users", controller.Get)
	r, _ := http.NewRequest(http.MethodGet, "/users", nil)
	r.Header.Add(middleware.X_ORG_CTX_KEY, "guima")
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	var got map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{"message": "invalid user id"}
	assert.Equal(t, want, got)
}

func TestGetNotFound(t *testing.T) {
	wantErr := errors.New("want error")
	controller := NewUserController(services.UserServiceCustomMock{
		GetMock: func(id string) (*entity.User, error) {
			return nil, wantErr
		},
	})
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.Organization.Handler)
	router.GET("/users/:id", controller.Get)
	r, _ := http.NewRequest(http.MethodGet, "/users/12345678", nil)
	r.Header.Add(middleware.X_ORG_CTX_KEY, "guima")
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNotFound, w.Code)
	var got map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{"message": "want error"}
	assert.Equal(t, want, got)
}

func TestGetNoHeaderXOrg(t *testing.T) {
	controller := NewUserController(services.UserServiceCustomMock{})
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.Organization.Handler)
	router.GET("/users/:id", controller.Get)
	r, _ := http.NewRequest(http.MethodGet, "/users/12345678", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	var got map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{"x-org": "header x-org is required"}
	assert.Equal(t, want, got)
}

//Save tests
func TestSaveSuccess(t *testing.T) {
	want := response.User{
		Id:    uuid.New().String(),
		Name:  "Guima",
		Email: "guima@guima.com",
	}
	controller := NewUserController(services.UserServiceCustomMock{
		SaveMock: func(user *entity.User) error {
			user.Id = want.Id
			return nil
		},
	})
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.Organization.Handler)
	router.POST("/users", controller.Save)
	userCreate := request.CreateUser{
		Name:  "Guima",
		Email: "guima@guima.com",
	}
	body, _ := json.Marshal(userCreate)
	r, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	r.Header.Add(middleware.X_ORG_CTX_KEY, "guima")
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusCreated, w.Code)
	var got response.User
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, want, got)
}

func TestSaveInvalidBody(t *testing.T) {
	controller := NewUserController(services.UserServiceCustomMock{})
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.Organization.Handler)
	router.POST("/users", controller.Save)
	type invalid struct {
		Name int64 `json:"name"`
	}
	request := invalid{
		Name: 35,
	}
	body, _ := json.Marshal(request)
	r, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	r.Header.Add(middleware.X_ORG_CTX_KEY, "guima")
 	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	var got map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{"message": "invalid json body"}
	assert.Equal(t, want, got)
}

func TestSaveInternalError(t *testing.T) {
	wantErr := errors.New("want error")
	controller := NewUserController(services.UserServiceCustomMock{
		SaveMock: func(user *entity.User) error {
			return wantErr
		},
	})
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.Organization.Handler)
	router.POST("/users", controller.Save)
	userCreate := request.CreateUser{
		Name:  "Guima",
		Email: "guima@guima.com",
	}
	body, _ := json.Marshal(userCreate)
	r, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	r.Header.Add(middleware.X_ORG_CTX_KEY, "guima")
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	var got map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{"message": "want error"}
	assert.Equal(t, want, got)
}

func TestSaveNoHeaderXOrg(t *testing.T) {
	controller := NewUserController(services.UserServiceCustomMock{})
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.Organization.Handler)
	router.POST("/users", controller.Save)
	userCreate := request.CreateUser{
		Name:  "Guima",
		Email: "guima@guima.com",
	}
	body, _ := json.Marshal(userCreate)
	r, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	var got map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{"x-org": "header x-org is required"}
	assert.Equal(t, want, got)
}
