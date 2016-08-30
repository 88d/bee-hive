package answers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

var (
	paramQuestion = "questionId"
	paramID       = "id"
	repository    *Repository
)

// New creates a new instance of the answers controller
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
	var items, err = repository.GetAll(c.Param(paramQuestion))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, items)
}

func getByID(c echo.Context) error {
	a, err := repository.GetByID(c.Param(paramQuestion), c.Param(paramID))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, a)
}

func create(c echo.Context) error {
	a := new(Answer)
	a.CreatedAt = time.Now().UTC()
	if err := c.Bind(a); err != nil {
		return err
	}
	if err := repository.Create(c.Param(paramQuestion), a); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, a)
}

func update(c echo.Context) error {
	a := new(Answer)
	a.UpdatedAt = time.Now().UTC()
	if err := c.Bind(a); err != nil {
		return err
	}
	a.ID = c.Param(paramID)
	if err := repository.Update(c.Param(paramQuestion), a); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, a)
}

func delete(c echo.Context) error {
	if err := repository.RemoveByID(c.Param(paramQuestion), c.Param(paramID)); err != nil {
		return err
	}
	return c.String(http.StatusOK, "")
}
