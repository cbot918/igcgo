package service

import (
	"github.com/cbot918/liby/jwty"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) GetJwtToken(id int, email string) string {
	token, err := jwty.New().FastJwt(id, email)
	if err != nil {
		panic(err)
	}
	return token
}
