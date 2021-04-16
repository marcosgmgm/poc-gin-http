package app

import (
	"github.com/marcosgmgm/poc-gin-http/controllers"
	"github.com/marcosgmgm/poc-gin-http/dao"
	"github.com/marcosgmgm/poc-gin-http/services"
)

func mapUrlsUser() {
	userService := services.NewUserService(dao.NewUserDao())
	controller := controllers.NewUserController(userService)
	router.GET("/users/:id", controller.Get)
	router.POST("/users", controller.Save)
}