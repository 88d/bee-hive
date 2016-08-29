package main

import (
	"log"
	"net/http"

	"github.com/black-banana/bee-hive/hive/auth"
	"github.com/black-banana/bee-hive/hive/hub"
	"github.com/black-banana/bee-hive/hive/questions"
	"github.com/black-banana/bee-hive/hive/users"
	"github.com/black-banana/bee-hive/rethink"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	// LoadConfiguration from config.json
	loadConfiguration()

	rethink.StartMasterSession(globalConfig.db)
	defer rethink.StopMasterSession()

	// Echo instance
	e := echo.New()
	e.SetDebug(true)
	e.SetHTTPErrorHandler(JSONErrorHandler)

	// Middleware
	e.Use(middleware.Logger())

	// e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: globalConfig.auth.SigningKey,
	// 	Claims:     globalConfig.auth.Claims,
	// }))
	e.Use(middleware.Recover())

	api := e.Group("/api")
	questions.New(api)
	users.New(api)
	auth.New(api, globalConfig.auth)

	h := hub.NewHub()
	go h.Run()
	go questions.GetAllChanges(h)

	e.GET("/hub", standard.WrapHandler(http.HandlerFunc(h.ServeHub())))

	routes := e.Routes()
	for _, route := range routes {
		log.Println(route.Method, route.Path)
	}

	log.Println("Started with", globalConfig.listen)
	e.Run(standard.New(globalConfig.listen))
}
