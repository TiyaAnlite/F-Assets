package main

import (
	"errors"
	"fmt"
	"strconv"

	. "github.com/TiyaAnlite/F-Assests/types"
	"github.com/TiyaAnlite/FocotServicesCommon/echox"
	"github.com/duke-git/lancet/v2/random"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func postPosition(c echo.Context) error {
	// 同时支持新增和修改
	req, err := echox.CheckInput[PositionOptRequest](c)
	if err != nil {
		return BadRequest(c, err)
	}
	var newPos Position
	if req.ID == "" {
		// new
		newPos.ID = random.RandNumeral(5)
		newPos.Name = req.Name
		if err := db.DB().
			Create(&newPos).
			Error; err != nil {
			return InternalError(c, err)
		}
	} else {
		// fetch first
		if err := db.DB().
			Where("id = ?", req.ID).
			Take(&newPos).
			Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return NotFound(c, err)
			}
			return InternalError(c, err)
		}
		newPos.Name = req.Name
		if q := db.DB().
			Updates(&newPos); q.Error != nil {
			return InternalError(c, err)
		}
	}
	return echox.NormalResponse(c, &newPos)
}

func getAsset(c echo.Context) error {
	// 同时支持id和code
	id, err := NewRequestID(c.Param("id"))
	if err != nil {
		return BadRequest(c, err)
	}
	assetType := AssetType(c.QueryParam("type"))
	switch assetType {
	case "":
		var asset Asset
		if err := db.DB().
			Scopes(id.QueryScope).
			Preload(clause.Associations).
			Take(&asset).
			Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return NotFound(c, errors.New("not found"))
			}
			return InternalError(c, err)
		}
		return echox.NormalResponse(c, asset)
	case AssetBasicItemType:
		var asset Asset
		if err := db.DB().
			Scopes(id.QueryScope).
			Where("type = ?", AssetBasicItemType).
			Preload(clause.Associations).
			Take(&asset).
			Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return NotFound(c, errors.New("not found"))
			}
			return InternalError(c, err)
		}
		return echox.NormalResponse(c, asset)
	case AssetBookType:
		var asset []Book
		if err := db.DB().
			Joins("LEFT JOIN asset ON asset.id = book.asset_id").
			Scopes(id.QueryScope).
			Preload(clause.Associations).
			Preload("Asset.Position").
			Find(&asset).Error; err != nil {
			return InternalError(c, err)
		}
		if len(asset) == 0 {
			return NotFound(c, errors.New("not found"))
		}
		return echox.NormalResponse(c, asset[0])
	default:
		return BadRequest(c, errors.New("asset type is not supported"))
	}
}

func postAsset(c echo.Context) error {
	// 同时支持新增和修改
	assetType := AssetType(c.QueryParam("type"))
	var omitField []string
	switch assetType {
	case "":
		return BadRequest(c, errors.New("asset type is required"))
	case AssetBasicItemType:
		req, err := echox.CheckInput[ItemOptRequest](c)
		if err != nil {
			return BadRequest(c, err)
		}
		asset, isNew, updatedOmitField, shouldReturn, err := commonReqCheck(c, req, omitField, AssetBasicItemType, "")
		if err != nil {
			return InternalError(c, err)
		}
		if shouldReturn {
			return nil
		}
		omitField = updatedOmitField
		newAsset := asset
		shouldReturn, err = createOrUpdate(c, &newAsset, isNew, omitField)
		if err != nil {
			return InternalError(c, err)
		}
		if shouldReturn {
			return nil
		}
		return echox.NormalResponse(c, &newAsset)
	case AssetBookType:
		req, err := echox.CheckInput[BookOptRequest](c)
		if err != nil {
			return BadRequest(c, err)
		}
		asset, isNew, updatedOmitField, shouldReturn, err := commonReqCheck(c, &req.ItemOptRequest, omitField, AssetBookType, "Asset")
		if err != nil {
			return InternalError(c, err)
		}
		if shouldReturn {
			return nil
		}
		omitField = updatedOmitField
		newBook := Book{Asset: asset}
		req.UpdateModel(&newBook)
		shouldReturn, err = createOrUpdate(c, &newBook, isNew, omitField)
		if err != nil {
			return InternalError(c, err)
		}
		if shouldReturn {
			return nil
		}
		return echox.NormalResponse(c, &newBook)
	case AssetCDType:
		req, err := echox.CheckInput[CDOptRequest](c)
		if err != nil {
			return BadRequest(c, err)
		}
		asset, isNew, updatedOmitField, shouldReturn, err := commonReqCheck(c, &req.ItemOptRequest, omitField, AssetCDType, "Asset")
		if err != nil {
			return InternalError(c, err)
		}
		if shouldReturn {
			return nil
		}
		omitField = updatedOmitField
		newCD := CD{Asset: asset}
		req.UpdateModel(&newCD)
		shouldReturn, err = createOrUpdate(c, &newCD, isNew, omitField)
		if err != nil {
			return InternalError(c, err)
		}
		if shouldReturn {
			return nil
		}
		return echox.NormalResponse(c, &newCD)
	default:
		return BadRequest(c, errors.New(fmt.Sprintf("unsupported asset type: %s", assetType)))
	}
}

func action(c echo.Context) error {
	id, err := NewRequestID(c.Param("id"))
	if err != nil {
		return BadRequest(c, err)
	}
	act := AssetStatusType(c.Param("action"))
	position := c.QueryParam("position")
	if position == "" {
		return BadRequest(c, errors.New("query param position is required"))
	}
	// find first
	var asset Asset
	if err := db.DB().
		Scopes(id.QueryScope).
		Take(&asset).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return NotFound(c, errors.New("asset not found"))
		}
		return InternalError(c, err)
	}
	var pos Position
	if err := db.DB().
		Where("id = ?", position).
		Take(&pos).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return NotFound(c, errors.New("position not found"))
		}
		return InternalError(c, err)
	}
	appendRecord := Record{
		ID:       strconv.FormatInt(snowFlake.NextVal(), 10),
		Position: pos,
	}
	switch act {
	case AssetStatusOutbound:
		if asset.Status != AssetStatusInbound {
			return BadRequest(c, errors.New(fmt.Sprintf("invalid asset status: %s", asset.Status)))
		}
		appendRecord.Operation = AssetOperationLeave
	case AssetStatusInbound:
		if asset.Status != AssetStatusOutbound {
			return BadRequest(c, errors.New(fmt.Sprintf("invalid asset status: %s", asset.Status)))
		}
		appendRecord.Operation = AssetOperationEnter
	case AssetStatusAbandon:
		if asset.Status == AssetStatusAbandon {
			return BadRequest(c, errors.New(fmt.Sprintf("invalid asset status: %s", asset.Status)))
		}
		appendRecord.Operation = AssetOperationDestroy
	}
	appendRecord.AssetID = asset.ID
	asset.Status = act
	asset.PositionID = pos.ID
	asset.Position = pos
	if err := db.DB().
		Select("status", "position_id").
		Updates(&asset).
		Error; err != nil {
		return InternalError(c, err)
	}
	if err := db.DB().
		Create(&appendRecord).
		Error; err != nil {
		return InternalError(c, err)
	}
	return echox.NormalResponse(c, &asset)
}
