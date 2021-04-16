package controllers

import "github.com/gin-gonic/gin"

type User interface {
	Get(c *gin.Context)
	Save(c *gin.Context)
}
