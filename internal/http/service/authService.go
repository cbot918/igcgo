package service

import (
	"database/sql"

	"github.com/cbot918/igcgo/internal/http/repository"
	"github.com/cbot918/igcgo/lib"
)

type AuthService struct {
	Db *sql.DB
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{
		Db: db,
	}
}

func (a *AuthService) GetJwtToken(id int, email string) string {
	jwty := lib.NewJwty()
	token, err := jwty.FastJwt(id, email)
	if err != nil {
		panic(err)
	}

	authRepo := repository.NewAuth(a.Db)
	authRepo.SaveUserToken()
	// unImplemented: set jwt token to database

	return token
}
