package main

import (
	"errors"
	. "github.com/TiyaAnlite/F-Assests/types"
	"github.com/TiyaAnlite/FocotServicesCommon/echox"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func listPosition(c echo.Context) error {
	var pos []Position
	if err := db.DB().
		Order("create_time").
		Find(&pos).Error; err != nil {
		return InternalError(c, err)
	}
	return echox.NormalResponse(c, &pos)
}

func listAsset(c echo.Context) error {
	assetType := AssetType(c.QueryParam("type"))
	switch assetType {
	case "":
		var asset []Asset
		if err := db.DB().
			Find(&asset).
			Error; err != nil {
			return InternalError(c, err)
		}
		return echox.NormalResponse(c, &asset)
	case AssetBasicItemType:
		var asset []Asset
		if err := db.DB().
			Order("last_update desc").
			Where("type = ?", AssetBasicItemType).
			Find(&asset).
			Error; err != nil {
			return InternalError(c, err)
		}
		return echox.NormalResponse(c, &asset)
	case AssetBookType:
		var asset []Book
		if err := db.DB().
			Preload(clause.Associations).
			Preload("Asset.Position").
			Find(&asset).
			Error; err != nil {
			return InternalError(c, err)
		}
		slice.SortBy(asset, func(a, b Book) bool {
			return a.Asset.LastUpdate.Before(b.Asset.LastUpdate)
		})
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
	pos := Position{}
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
	id, err := NewRequestID(c.Param("id"))
	if err != nil {
		return BadRequest(c, err)
	}
	var asset Asset
	if err := db.DB().
		Scopes(id.QueryScope).
		Take(&asset, id).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return NotFound(c, errors.New("asset not found"))
		}
		return InternalError(c, err)
	}
	var records []Record
	if err := db.DB().
		Where("asset_id = ?", asset.ID).
		Find(&records).
		Error; err != nil {
		return InternalError(c, err)
	}
	return echox.NormalResponse(c, &records)
}
