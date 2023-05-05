package rest

import (
	"fmt"
	"os"

	"github.com/cbot918/health-helper/internal/config"

	"github.com/gin-gonic/gin"
)

type Rest struct{}

func New() *Rest {
	return &Rest{}
}

func (r *Rest) Run() {
	// setup config
	cfg, err := config.New()
	if err != nil {
		fmt.Println("config.New error")
		os.Exit(1)
	}

	// setup server
	server := gin.Default()

	server = RegistRoute(server)

	server.Run(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port))
}
