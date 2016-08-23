package answers

import (
	"net/http"

	"github.com/labstack/echo"
)

func New(e *echo.Group) {
	repository = NewRepository()
	group := e.Group("/:questionId/answers")
	group.Get("", getAll)
	group.Post("", create)
	group.Put("/:id", update)
	group.Get("/:id", getByID)
	group.Delete("/:id", delete)
}

func getAll(c echo.Context) error {
	var items, err = repository.GetAll(c.Param("questionId"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, items)
}

func getByID(c echo.Context) error {
	a, err := repository.GetByID(c.Param("questionId"), c.Param("id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, a)
}

func create(c echo.Context) error {
	a := new(Answer)
	if err := c.Bind(a); err != nil {
		return err
	}
	if err := repository.Create(c.Param("questionId"), a); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, a)
}

func update(c echo.Context) error {
	a := new(Answer)
	if err := c.Bind(a); err != nil {
		return err
	}
	a.ID = c.Param("id")
	if err := repository.Update(c.Param("questionId"), a); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, a)
}

func delete(c echo.Context) error {
	if err := repository.RemoveByID(c.Param("questionId"), c.Param("id")); err != nil {
		return err
	}
	return c.String(http.StatusOK, "")
}
