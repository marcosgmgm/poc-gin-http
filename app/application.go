package app

import (
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/marcosgmgm/poc-gin-http/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var (
	router = gin.Default()
)

func StartApplication() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

	router.Use(middleware.Organization.Handler)
	router.Use(logger.SetLogger())
	mapUrlsUser()
	router.Run(":3000")
}
