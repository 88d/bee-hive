package questions

import (
	"net/http"

	"github.com/labstack/echo"
)

var repository Repository

func New(e *echo.Group, server string, database string) {
	repository = NewRepository(server, database)
	group := e.Group("/questions")
	group.Get("", getAll)
	group.Post("", create)
	group.Put("/:id", update)
	group.Delete("/:id", delete)
}

func Close() {
	if repository.session != nil {
		repository.session.Close()
	}
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
