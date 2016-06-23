package questions

import (
	"net/http"

	"github.com/labstack/echo"
)

func New(e *echo.Group) {
	repository = NewRepository()
	group := e.Group("/questions")
	group.Get("", getAll)
	group.Post("", create)
	group.Put("/:id", update)
	group.Delete("/:id", delete)
}

func getAll(c echo.Context) error {
	var items, err = repository.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, items)
}

func create(c echo.Context) error {
	q := new(Question)
	if err := c.Bind(q); err != nil {
		return err
	}
	if err := repository.Create(q); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, q)
}

func update(c echo.Context) error {
	q := new(Question)
	if err := c.Bind(q); err != nil {
		return err
	}
	q.ID = c.Param("id")
	if err := repository.Update(q); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, q)
}

func delete(c echo.Context) error {
	if err := repository.Remove(c.Param("id")); err != nil {
		return err
	}
	return c.String(http.StatusOK, "")
}
