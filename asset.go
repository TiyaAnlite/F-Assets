package main

import (
	"errors"
	"github.com/TiyaAnlite/F-Assests/types"
	"github.com/TiyaAnlite/FocotServicesCommon/echox"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func listPosition(c echo.Context) error {
	var pos []types.Position
	if err := db.DB().Find(&pos).Error; err != nil {
		return InternalError(c, err)
	}
	return echox.NormalResponse(c, &pos)
}

func listAsset(c echo.Context) error {
	assetType := types.AssetType(c.QueryParam("type"))
	switch assetType {
	case "":
		var asset []types.Asset
		if err := db.DB().
			Find(&asset).
			Error; err != nil {
			return InternalError(c, err)
		}
		return echox.NormalResponse(c, &asset)
	case types.AssetBasicItemType:
		var asset []types.Asset
		if err := db.DB().
			Where("type = ?", types.AssetBasicItemType).
			Find(&asset).
			Error; err != nil {
			return InternalError(c, err)
		}
		return echox.NormalResponse(c, &asset)
	case types.AssetBookType:
		var asset []types.Book
		if err := db.DB().
			Preload(clause.Associations).
			Preload("Asset.Position").
			Find(&asset).
			Error; err != nil {
			return InternalError(c, err)
		}
		return echox.NormalResponse(c, &asset)
	default:
		return BadRequest(c, errors.New("asset type is not supported"))
	}
}

func getPosition(c echo.Context) error {
	// 仅支持pos的唯一id
	id := c.Param("id")
	if id == "" {
		return BadRequest(c, errors.New("id is required"))
	}
	pos := types.Position{}
	if err := db.DB().
		Where("id = ?", id).
		Take(&pos).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return NotFound(c, errors.New("position not found"))
		}
		return InternalError(c, err)
	}
	return echox.NormalResponse(c, &pos)
}

func getRecords(c echo.Context) error {
	id, err := types.NewRequestID(c.Param("id"))
	if err != nil {
		return BadRequest(c, err)
	}
	assetId := id.Identifier()
	if id.IDType() == types.EAN13IDType {
		// take id from asset
		var asset types.Asset
		if err := db.DB().
			Scopes(id.QueryScope).
			Take(&asset).
			Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return NotFound(c, errors.New("asset not found"))
			}
			return InternalError(c, err)
		}
		assetId = asset.ID
	}
	var records []types.Record
	if err := db.DB().
		Where("asset_id = ?", assetId).
		Order("id desc").
		Find(&records).
		Error; err != nil {
		return InternalError(c, err)
	}
	return echox.NormalResponse(c, &records)
}
