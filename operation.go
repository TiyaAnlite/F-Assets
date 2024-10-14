package main

import (
	"errors"
	"github.com/TiyaAnlite/F-Assests/types"
	"github.com/TiyaAnlite/FocotServicesCommon/echox"
	"github.com/labstack/echo/v4"
	"strconv"
)

func postPosition(c echo.Context) error {
	// 同时支持新增和修改
	req, err := echox.CheckInput[types.PositionOptRequest](c)
	if err != nil {
		return BadRequest(c, err)
	}
	if req.ID == "" {
		// new
		newPos := types.Position{
			ID:   strconv.FormatInt(snowFlake.NextVal(), 10),
			Name: req.Name,
		}
		if err := db.DB().
			Create(&newPos).
			Error; err != nil {
			return InternalError(c, err)
		}
	} else {
		// edit
		if q := db.DB().
			Where("id = ?", req.ID).
			Update("name = ?", req.Name); q.Error != nil {
			return InternalError(c, err)
		} else if q.RowsAffected == 0 {
			return NotFound(c, errors.New("id not found"))
		}
	}
	return echox.NormalEmptyResponse(c)
}

func listAsset(c echo.Context) error {
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
	id := c.Param("id")
	act := c.Param("action")
	assetType := c.Param("type")
	_ = act
	_ = id
	switch assetType {

	}
	return nil
}
