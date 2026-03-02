package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed frontend/dist/*
var static embed.FS

func setupRoutes(e *echo.Echo) {
	dist, _ := fs.Sub(static, "frontend/dist")
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       ".",
		HTML5:      true,
		Filesystem: http.FS(dist),
	}))

	api := e.Group("/api")
	api.GET("/position", listPosition)
	api.GET("/position/:id", getPosition)
	api.POST("/position", postPosition)

	api.GET("/asset", listAsset)
	api.GET("/asset/:id", getAsset)
	api.POST("/asset", postAsset)

	api.GET("/action/:id/:action", action)

	api.GET("/record/:id", getRecords)
}
