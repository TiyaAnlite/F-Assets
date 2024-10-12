package main

import (
	"errors"
	"github.com/TiyaAnlite/F-Assests/types"
	"github.com/TiyaAnlite/FocotServicesCommon/echox"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func getItemContext(c echo.Context) (*types.ItemOptRequest, *types.Asset, error) {
	req, err := echox.CheckInput[types.ItemOptRequest](c)
	if err != nil {
		_ = echox.NormalErrorResponse(c, http.StatusBadRequest, http.StatusBadRequest, err.Error())
		return nil, nil, err
	}
	var position types.Position
	if req.Position != "" {
		// find position
		if err := db.DB().
			Where("id = ?", req.Position).
			Take(&position).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return req, nil, errors.New("position not found")
			}
		}
	}
	asset := &types.Asset{
		ID:       snowFlake.NextVal(),
		Type:     req.Type,
		Code:     req.Code,
		Name:     req.Name,
		Position: position,
		Pic:      req.Pic,
	}
	return req, asset, nil
}

func getBookContext(c echo.Context) (*types.BookOptRequest, *types.Book, error) {
	req, err := echox.CheckInput[types.BookOptRequest](c)
	if err != nil {
		_ = echox.NormalErrorResponse(c, http.StatusBadRequest, http.StatusBadRequest, err.Error())
		return nil, nil, err
	}
	var position types.Position
	if req.Position != "" {
		// find position
		if err := db.DB().
			Where("id = ?", req.Position).
			Take(&position).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return req, nil, errors.New("position not found")
			}
		}
	}
	asset := types.Asset{
		ID:       snowFlake.NextVal(),
		Type:     req.Type,
		Code:     req.Code,
		Name:     req.Name,
		Position: position,
		Pic:      req.Pic,
	}
	book := &types.Book{
		Asset:          asset,
		Author:         req.Author,
		Publisher:      req.Publisher,
		Specifications: req.Specifications,
		Tag:            req.Tag,
		Language:       req.Language,
		PurchaseTime:   req.PurchaseTime,
	}
	return req, book, nil
}
