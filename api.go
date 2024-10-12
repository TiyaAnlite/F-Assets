package main

import "github.com/labstack/echo/v4"

func setupRoutes(e *echo.Echo) {
	e.GET("/position/:id", getPosition)
	e.POST("/position/:id", postPosition)

	e.GET("/asset/:id", getAsset)
	e.POST("/asset", postAsset)

	e.GET("/action/:action", action)
}
