package questions

import (
	"net/http"
	"time"

	"github.com/black-banana/bee-hive/hive/answers"
	"github.com/labstack/echo"
)

var (
	paramID    = "id"
	repository *Repository
)

// New creates an new instance of the questions controller
func New(e *echo.Group) {
	repository = NewRepository()
	group := e.Group("/questions")
	group.Get("", getAll)
	group.Post("", create)
	group.Put("/:id", update)
	group.Get("/:id", getByID)
	group.Delete("/:id", delete)
	answers.New(group)
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
	q := new(Question)
	q.CreatedAt = time.Now().UTC()
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
	q.ID = c.Param(paramID)
	q.UpdatedAt = time.Now().UTC()
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
