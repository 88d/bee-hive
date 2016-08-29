package auth

import (
	"net/http"

	"github.com/labstack/echo"
)

type LoginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	config *Config
)

// New creates a new instance of the users Controller
func New(e *echo.Group, c *Config) {
	config = c
	group := e.Group("/auth")
	group.Post("/login", login)
}

func login(c echo.Context) error {
	login := new(LoginModel)
	if err := c.Bind(&login); err != nil {
		return err
	}

	if login.Username == "opi" && login.Password == "opi" {
		t := GenerateToken("hahaha", []string{"admin"})
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}
