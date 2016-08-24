package auth

import (
	"net/http"

	"github.com/labstack/echo"
)

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
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "opi" && password == "opi" {
		t := GenerateToken("hahaha", []string{"admin"})
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}
