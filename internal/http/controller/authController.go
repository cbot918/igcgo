package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/cbot918/igcgo/internal/http/service"
	"github.com/cbot918/igcgo/internal/model"
	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	Db *sql.DB
}

func NewAuthController(db *sql.DB) *Auth {
	return &Auth{
		Db: db,
	}
}

func (a *Auth) Auth(c *fiber.Ctx) error {
	// fmt.Println("in auth")
	// return c.SendString("got")
	user := model.Users{}
	if err := c.BodyParser(&user); err != nil {
		panic(err)
	}

	authService := service.NewAuthService(a.Db)
	token := authService.GetJwtToken(int(user.Id), user.Email)

	res := struct {
		Token string `json:"token"`
		User  string `json:"user"`
	}{}
	res.Token = token
	res.User = strings.Trim(regexp.MustCompile(".*@").FindString(user.Email), "@")
	fmt.Printf("user: %s", res.User)
	resp, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	return c.SendString(string(resp))
}
