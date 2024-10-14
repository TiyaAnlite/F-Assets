package main

import (
	"github.com/TiyaAnlite/FocotServicesCommon/echox"
	"github.com/labstack/echo/v4"
	"net/http"
)

func BadRequest(c echo.Context, err error) error {
	return echox.NormalErrorResponse(c, http.StatusBadRequest, http.StatusBadRequest, err.Error())
}

func NotFound(c echo.Context, err error) error {
	return echox.NormalErrorResponse(c, http.StatusNotFound, http.StatusNotFound, err.Error())
}

func InternalError(c echo.Context, err error) error {
	return echox.NormalErrorResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, err.Error())
}
