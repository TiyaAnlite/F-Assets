package main

import (
	"errors"
	"fmt"
	"github.com/TiyaAnlite/F-Assests/types"
	"github.com/TiyaAnlite/FocotServicesCommon/echox"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
		// edit, id not checked
		if q := db.DB().
			Where("id = ?", req.ID).
			Update("name = ?", req.Name); q.Error != nil {
			return InternalError(c, err)
		}
	}
	return echox.NormalEmptyResponse(c)
}

func getAsset(c echo.Context) error {
	// 同时支持id和code
	id, err := types.NewRequestID(c.Param("id"))
	if err != nil {
		return BadRequest(c, err)
	}
	assetType := types.AssetType(c.QueryParam("type"))
	switch assetType {
	case "":
		var asset types.Asset
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
	case types.AssetBasicItemType:
		var asset types.Asset
		if err := db.DB().
			Scopes(id.QueryScope).
			Where("type = ?", types.AssetBasicItemType).
			Take(&asset).
			Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return NotFound(c, errors.New("not found"))
			}
			return InternalError(c, err)
		}
		return echox.NormalResponse(c, asset)
	case types.AssetBookType:
		var asset []types.Book
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
	assetType := types.AssetType(c.QueryParam("type"))
	switch assetType {
	case "":
		return BadRequest(c, errors.New("asset type is required"))
	case types.AssetBasicItemType:
		req, err := echox.CheckInput[types.ItemOptRequest](c)
		if err != nil {
			return BadRequest(c, err)
		}
		var pos types.Position
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
		var newAsset types.Asset
		if req.ID == "" {
			// new
			newAsset.ID = snowFlake.NextVal()
		} else {
			// edit
			newAsset.ID, err = strconv.ParseInt(req.ID, 10, 64)
			if err != nil {
				return BadRequest(c, err)
			}
		}
		newAsset.Type = types.AssetBasicItemType
		newAsset.Code = req.Code
		newAsset.Name = req.Name
		newAsset.Position = pos
		newAsset.Pic = req.Pic
		if req.ID == "" {
			if err := db.DB().
				Create(&newAsset).
				Error; err != nil {
				return InternalError(c, err)
			}
		} else {
			if err := db.DB().
				Updates(&newAsset).
				Error; err != nil {
				return InternalError(c, err)
			}
		}
		return echox.NormalResponse(c, &newAsset)
	case types.AssetBookType:
		req, err := echox.CheckInput[types.BookOptRequest](c)
		if err != nil {
			return BadRequest(c, err)
		}
		var pos types.Position
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
		var newBook types.Book
		if req.ID == "" {
			// new
			newBook.Asset.ID = snowFlake.NextVal()
		} else {
			// edit
			newBook.Asset.ID, err = strconv.ParseInt(req.ID, 10, 64)
			if err != nil {
				return BadRequest(c, err)
			}
		}
		newBook.Asset.Type = types.AssetBookType
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
				Create(&newBook).
				Error; err != nil {
				return InternalError(c, err)
			}
		} else {
			if err := db.DB().
				Session(&gorm.Session{FullSaveAssociations: true}).
				Updates(&newBook).
				Error; err != nil {
				return InternalError(c, err)
			}
		}
		return echox.NormalResponse(c, &newBook)
	default:
		return BadRequest(c, errors.New(fmt.Sprintf("unsupported asset type: %s", assetType)))
	}
}

func action(c echo.Context) error {
	id, err := types.NewRequestID(c.Param("id"))
	if err != nil {
		return BadRequest(c, err)
	}
	act := types.AssetStatusType(c.Param("action"))
	position := c.QueryParam("position")
	if position == "" {
		return BadRequest(c, errors.New("query param position is required"))
	}
	// find first
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
	var pos types.Position
	if err := db.DB().
		Where("id = ?", position).
		Take(&pos).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return NotFound(c, errors.New("position not found"))
		}
		return InternalError(c, err)
	}
	appendRecord := types.Record{
		ID:       snowFlake.NextVal(),
		Position: pos,
	}
	switch act {
	case types.AssetStatusOutbound:
		if asset.Status != types.AssetStatusInbound {
			return BadRequest(c, errors.New(fmt.Sprintf("invalid asset status: %s", asset.Type)))
		}
		appendRecord.Operation = types.AssetOperationLeave
	case types.AssetStatusInbound:
		if asset.Status != types.AssetStatusOutbound {
			return BadRequest(c, errors.New(fmt.Sprintf("invalid asset status: %s", asset.Type)))
		}
		appendRecord.Operation = types.AssetOperationEnter
	case types.AssetStatusAbandon:
		if asset.Status == types.AssetStatusAbandon {
			return BadRequest(c, errors.New(fmt.Sprintf("invalid asset status: %s", asset.Type)))
		}
		appendRecord.Operation = types.AssetOperationDestroy
	}
	appendRecord.Asset = asset
	if err := db.DB().
		Model(&asset).
		Update("status", act).
		Error; err != nil {
		return InternalError(c, err)
	}
	if err := db.DB().
		Create(&appendRecord).
		Error; err != nil {
		return InternalError(c, err)
	}
	return echox.NormalEmptyResponse(c)
}
