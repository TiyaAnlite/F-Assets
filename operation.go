package main

import (
	"errors"
	"fmt"
	. "github.com/TiyaAnlite/F-Assests/types"
	"github.com/TiyaAnlite/FocotServicesCommon/echox"
	"github.com/duke-git/lancet/v2/random"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"strconv"
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
			Scopes(id.QueryScope).
			Association("asset").
			Find(&asset); err != nil {
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
	insertRecord := func(asset Asset, operation AssetOperation, position Position) error {
		newRecord := Record{
			ID:        strconv.FormatInt(snowFlake.NextVal(), 10),
			AssetID:   asset.ID,
			Operation: operation,
			Position:  position,
		}
		q := db.DB()
		if position.ID == "" {
			q = q.Omit("position_id")
		}
		if err := q.Create(&newRecord).Error; err != nil {
			return err
		}
		return nil
	}
	var omitField []string
	switch assetType {
	case "":
		return BadRequest(c, errors.New("asset type is required"))
	case AssetBasicItemType:
		req, err := echox.CheckInput[ItemOptRequest](c)
		if err != nil {
			return BadRequest(c, err)
		}
		// code insert check
		if req.Code == "" {
			omitField = append(omitField, "code")
		}
		var pos Position
		if req.Position != "" {
			// find position first
			if err := db.DB().
				Where("id = ?", req.Position).
				Take(&pos).
				Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return NotFound(c, errors.New("position not found"))
				}
				return InternalError(c, err)
			}
		}
		// position insert check
		if pos.ID == "" {
			omitField = append(omitField, "position_id")
		}
		var newAsset Asset
		if req.ID == "" {
			// new
			newAsset.ID = strconv.FormatInt(snowFlake.NextVal(), 10)
		} else {
			// edit
			newAsset.ID = req.ID
		}
		newAsset.Type = AssetBasicItemType
		newAsset.Code = req.Code
		newAsset.Name = req.Name
		newAsset.Position = pos
		newAsset.Pic = req.Pic
		if req.ID == "" {
			if err := db.DB().
				Omit(omitField...).
				Create(&newAsset).
				Error; err != nil {
				return InternalError(c, err)
			}
			if err := insertRecord(newAsset, AssetOperationCreate, pos); err != nil {
				return InternalError(c, err)
			}
		} else {
			if err := db.DB().
				Updates(&newAsset).
				Error; err != nil {
				return InternalError(c, err)
			}
			if err := insertRecord(newAsset, AssetOperationPost, pos); err != nil {
				return InternalError(c, err)
			}
		}
		return echox.NormalResponse(c, &newAsset)
	case AssetBookType:
		req, err := echox.CheckInput[BookOptRequest](c)
		if err != nil {
			return BadRequest(c, err)
		}
		// code insert check
		if req.Code == "" {
			omitField = append(omitField, "Asset.code")
		}
		var pos Position
		if req.Position != "" {
			// find position first
			if err := db.DB().
				Where("id = ?", req.Position).
				Take(&pos).
				Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return NotFound(c, errors.New("position not found"))
				}
				return InternalError(c, err)
			}
		}
		// position insert check
		if pos.ID == "" {
			omitField = append(omitField, "Asset.position_id")
		}
		var newBook Book
		if req.ID == "" {
			// new
			newBook.Asset.ID = strconv.FormatInt(snowFlake.NextVal(), 10)
		} else {
			// edit
			newBook.Asset.ID = req.ID
		}
		newBook.Asset.Type = AssetBookType
		newBook.Asset.Code = req.Code
		newBook.Asset.Name = req.Name
		newBook.Asset.Position = pos
		newBook.Asset.Pic = req.Pic
		newBook.Author = req.Author
		newBook.Publisher = req.Publisher
		newBook.Specifications = req.Specifications
		newBook.Tag = req.Tag
		newBook.Language = req.Language
		newBook.PurchaseTime = req.PurchaseTime
		if req.ID == "" {
			if err := db.DB().
				Omit(omitField...).
				Create(&newBook).
				Error; err != nil {
				return InternalError(c, err)
			}
			if err := insertRecord(newBook.Asset, AssetOperationCreate, pos); err != nil {
				return InternalError(c, err)
			}
		} else {
			if err := db.DB().
				Session(&gorm.Session{FullSaveAssociations: true}).
				Omit(omitField...).
				Updates(&newBook).
				Error; err != nil {
				return InternalError(c, err)
			}
			if err := insertRecord(newBook.Asset, AssetOperationPost, pos); err != nil {
				return InternalError(c, err)
			}
		}
		return echox.NormalResponse(c, &newBook)
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
