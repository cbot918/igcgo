package main

import (
	"fmt"
	"os"

	"github.com/cbot918/igcgo/internal/config"
	"github.com/cbot918/igcgo/internal/http"
	"github.com/cbot918/igcgo/lib"
)

func main() {

	// Config Init
	cfg, err := config.New()
	if err != nil {
		fmt.Println("config.New error")
		os.Exit(1)
	}

	// Database Init
	dby := lib.NewDby()
	connStr := dby.GetConnStr(
		cfg.Connection.Postgresql.DbType,
		cfg.Connection.Postgresql.User,
		cfg.Connection.Postgresql.Password,
		cfg.Connection.Postgresql.Host,
		cfg.Connection.Postgresql.Db,
	)
	conn := dby.MakeConn(cfg.Connection.Postgresql.DbType, connStr)
	defer conn.Close()
	dby.Ping(conn)

	// HttpServer Init
	server := http.NewHttpServer(conn)
	server.Static("/", cfg.Web.StaticPath) // serve spa
	err = server.Listen(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		panic(err)
	}
}
