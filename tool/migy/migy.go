package main

import (
	"fmt"
	"os"

	"github.com/cbot918/igcgo/internal/config"
	"github.com/cbot918/igcgo/lib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var scheme = `
CREATE TABLE users (
	id serial PRIMARY KEY,
	email text NOT NULL,
	password text NOT NULL,
	name text
);

CREATE TABLE post (
	title text,
	content text
)
`

var data = `
INSERT INTO users (email, password, name) VALUES ('yale918@gmail.com','12345', 'yale');
INSERT INTO users (email, password, name) VALUES ('yale918@gmail.com','12345', 'yale');
INSERT INTO users (email, password, name) VALUES ('yale918@gmail.com','12345', 'yale');
INSERT INTO users (email, password, name) VALUES ('yale918@gmail.com','12345', 'yale');
INSERT INTO users (email, password, name) VALUES ('yale918@gmail.com','12345', 'yale');
`

func main() {

	// setup config
	cfg, err := config.New()
	if err != nil {
		fmt.Println("config.New failed", err)
		os.Exit(1)
	}

	// establish connection
	dby := lib.NewDby()
	connStr := dby.GetConnStr(
		cfg.Connection.Postgresql.DbType,
		cfg.Connection.Postgresql.User,
		cfg.Connection.Postgresql.Password,
		cfg.Connection.Postgresql.Host,
		cfg.Connection.Postgresql.Db,
	)
	db, err := sqlx.Connect(cfg.Connection.Postgresql.DbType, connStr)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(scheme)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(data)
	if err != nil {
		panic(err)
	}
}
