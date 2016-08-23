package users

import (
	"net/http"

	"github.com/labstack/echo"
)

var paramID = "id"

// New creates a new instance of the users Controller
func New(e *echo.Group) {
	repository = NewRepository()
	group := e.Group("/users")
	group.Get("", getAll)
	group.Post("", create)
	group.Put("/:id", update)
	group.Get("/:id", getByID)
	group.Delete("/:id", delete)
}

func getAll(c echo.Context) error {
	var items, err = repository.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, items)
}

func getByID(c echo.Context) error {
	q, err := repository.GetByID(c.Param(paramID))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, q)
}

func create(c echo.Context) error {
	q := new(User)
	if err := c.Bind(q); err != nil {
		return err
	}
	if err := repository.Create(q); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, q)
}

func update(c echo.Context) error {
	q := new(User)
	if err := c.Bind(q); err != nil {
		return err
	}
	q.ID = c.Param(paramID)
	if err := repository.Update(q); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, q)
}

func delete(c echo.Context) error {
	if err := repository.RemoveByID(c.Param(paramID)); err != nil {
		return err
	}
	return c.String(http.StatusOK, "")
}
