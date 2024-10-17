package main

import "github.com/labstack/echo/v4"

func setupRoutes(e *echo.Echo) {
	e.GET("/position", listPosition)
	e.GET("/position/:id", getPosition)
	e.POST("/position", postPosition)

	e.GET("/asset", listAsset)
	e.GET("/asset/:id", getAsset)
	e.POST("/asset", postAsset)

	e.GET("/action/:id/:action", action)

	e.GET("/record/:id", getRecords)
}
