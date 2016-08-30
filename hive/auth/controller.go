package auth

import (
	"net/http"

	"github.com/black-banana/bee-hive/hive/users"
	"github.com/labstack/echo"
)

type LoginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	config     *Config
	repository *Repository
	userRepo   *users.Repository
)

// New creates a new instance of the users Controller
func New(e *echo.Group, c *Config) {
	config = c
	repository = NewRepository()
	userRepo = users.NewRepository()
	group := e.Group("/auth")
	group.Post("/login", login)
}

func login(c echo.Context) error {
	login := new(LoginModel)
	if err := c.Bind(&login); err != nil {
		return err
	}
	if login.Username == "opi" && login.Password == "opi" {
		user, err := userRepo.GetByName(login.Username)
		if err != nil {
			return err
		}
		userScope, err := repository.GetUserScopes(user.ID)
		if err != nil {
			return err
		}
		t := GenerateToken(userScope.UserID, userScope.Scopes)
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
	return echo.ErrUnauthorized
}

func info(c echo.Context) error {
	return nil
}
