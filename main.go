package main

import (
	"log"
	"net/http"

	"github.com/black-banana/bee-hive/hub"
	"github.com/black-banana/bee-hive/questions"
	"github.com/black-banana/bee-hive/rethink"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	// LoadConfiguration from config.json
	LoadConfiguration()

	rethink.StartMasterSession(globalConfig.db)
	defer rethink.StopMasterSession()

	// Echo instance
	e := echo.New()

	e.SetHTTPErrorHandler(func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if err == rethink.ErrEmptyResult {
			code = 404
		}
		msg := http.StatusText(code)
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			msg = he.Message
		}
		if e.Debug() {
			msg = err.Error()
		}
		if !c.Response().Committed() {
			if c.Request().Method() == echo.HEAD { // Issue #608
				c.NoContent(code)
			} else {
				c.JSON(code, JsonErrorWrapper{
					Error: JsonError{
						code,
						msg,
					},
				})
			}
		}
		e.Logger().Error(err)
	})

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")
	questions.New(api)

	go hub.Run()

	e.GET("/hub", standard.WrapHandler(http.HandlerFunc(hub.ServeHub())))

	routes := e.Routes()
	for _, route := range routes {
		log.Println(route.Method, route.Path)
	}

	log.Println("Started with", globalConfig.listen)
	e.Run(standard.New(globalConfig.listen))
}

type JsonError struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}

type JsonErrorWrapper struct {
	Error JsonError `json:"error"`
}
