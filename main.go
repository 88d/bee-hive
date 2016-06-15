package main

import (
	"fmt"

	"github.com/black-banana/bee-hive/questions"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	// LoadConfiguration from config.json
	LoadConfiguration()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")
	questions.New(api, globalConfig.dbServer, globalConfig.dbName)
	defer questions.Close()
	fmt.Println(e.Routes())
	fmt.Println("listening On ", globalConfig.listen)
	e.Run(standard.New(globalConfig.listen))
}
