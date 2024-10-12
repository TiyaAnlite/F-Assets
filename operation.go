package main

import (
	"github.com/TiyaAnlite/FocotServicesCommon/echox"
	"github.com/labstack/echo/v4"
	"net/http"
)

func getPosition(c echo.Context) error {
	return nil
}

func postPosition(c echo.Context) error {
	// 同时支持新增和修改
	return nil
}

func getAsset(c echo.Context) error {
	return nil
}

func postAsset(c echo.Context) error {
	// 同时支持新增和修改
	return nil
}

func action(c echo.Context) error {
	// 同时支持ID和Code
	act := c.Param("action")
	id := c.QueryParam("id")
	code := c.QueryParam("code")
	if id == "" && code == "" {
		return echox.NormalErrorResponse(c, http.StatusBadRequest, http.StatusBadRequest, "need id or code")
	}
	switch act {

	}
	return nil
}
