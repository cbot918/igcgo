package lib

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Dby struct{}

func NewDby() *Dby {
	return &Dby{}
}

func (d *Dby) GetConnStr(dbType string, user string, password string, host string, db string) string {
	return fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable", dbType, user, password, host, db)
}

func (d *Dby) MakeConn(dbType string, connStr string) *sql.DB {
	conn, err := sql.Open(dbType, connStr)
	if err != nil {
		panic(err)
	}
	return conn
}

func (d *Dby) Ping(conn *sql.DB) {
	err := conn.Ping()
	if err != nil {
		fmt.Println("dby:conn.Ping failed")
		panic(err)
	}
	fmt.Println("database ping success!")
}
