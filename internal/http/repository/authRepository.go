package repository

import (
	"database/sql"
	"fmt"
)

type AuthRepository struct {
	Db *sql.DB
}

func NewAuth(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		Db: db,
	}
}

func (a *AuthRepository) SaveUserToken() error {

	rows, err := a.Db.Query("select * from users;")
	if err != nil {
		return fmt.Errorf("Db.Query failed:%s", err)
	}
	defer rows.Close()
	type user struct {
		Name  string
		Job   string
		Email string
	}

	users := []user{}

	for rows.Next() {
		u := user{}
		err := rows.Scan(&u.Name, &u.Job, &u.Email)
		if err != nil {
			fmt.Println("in rows.Scan")
			panic(err)
		}
		users = append(users, u)
	}
	fmt.Println(users)

	return nil
}
