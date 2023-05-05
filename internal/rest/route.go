package rest

import (
	"github.com/cbot918/health-helper/internal/rest/controller"
	"github.com/gin-gonic/gin"
)

func RegistRoute(server *gin.Engine) *gin.Engine {
	ctr := controller.New()

	server.GET("/", ctr.Hello)

	return server
}
