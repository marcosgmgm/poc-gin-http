package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	Organization = organization{}
)

const X_ORG_CTX_KEY = "x-org"

type organization struct{}

func (org organization) Handler(c *gin.Context) {
	xOrg := c.GetHeader("x-org")
	if len(xOrg) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"x-org": "header x-org is required"})
	}
	c.Set(X_ORG_CTX_KEY, xOrg)
}
