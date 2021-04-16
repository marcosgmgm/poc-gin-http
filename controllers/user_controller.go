package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcosgmgm/poc-gin-http/api/request"
	"github.com/marcosgmgm/poc-gin-http/middleware"
	"github.com/marcosgmgm/poc-gin-http/services"
	"github.com/rs/zerolog/log"
	"net/http"
)

type usersController struct{
	service services.User
}

func NewUserController(s services.User) User {
	return usersController{
		service: s,
	}
}

func (controller usersController) Get(c *gin.Context) {
	org, _ := c.Get(middleware.X_ORG_CTX_KEY)
	log.Info().Msgf("receive x-org: %s", org)
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid user id"})
		return
	}

	user, err := controller.service.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user.ToResponse())
}

func (controller usersController) Save(c *gin.Context) {
	org, _ := c.Get(middleware.X_ORG_CTX_KEY)
	log.Info().Msgf("receive x-org: %s", org)
	var userRequest request.CreateUser
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid json body"})
		return
	}
	user := userRequest.GenerateEntity()
	if err := controller.service.Save(&user); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user.ToResponse())
}
