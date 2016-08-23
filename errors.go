package main

import (
	"net/http"

	"github.com/black-banana/bee-hive/rethink"
	"github.com/labstack/echo"
)

// JSONError is used to display the state and the error message
type JSONError struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}

// JSONErrorWrapper is used to create an error property
type JSONErrorWrapper struct {
	Error JSONError `json:"error"`
}

// JSONErrorHandler is used to display an json error instead of an string error
func JSONErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if err == rethink.ErrEmptyResult {
		code = 404
	}
	msg := http.StatusText(code)
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	}
	if c.Echo().Debug() {
		msg = err.Error()
	}
	if !c.Response().Committed() {
		if c.Request().Method() == echo.HEAD {
			c.NoContent(code)
		} else {
			c.JSON(code, JSONErrorWrapper{
				Error: JSONError{
					code,
					msg,
				},
			})
		}
	}
	c.Logger().Error(err)
}
