package main

import (
	"errors"
	"net/http"
	"strconv"

	. "github.com/TiyaAnlite/F-Assests/types"
	"github.com/TiyaAnlite/FocotServicesCommon/echox"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

func commonReqCheck(c echo.Context, req *ItemOptRequest, omitField []string, assetType AssetType, subAssetPrefix string) (assent Asset, isNew bool, updatedOmitField []string, shouldReturn bool, err error) {
	if assetType == "" {
		assetType = AssetBasicItemType
	}
	// code insert check
	if req.Code == "" {
		if subAssetPrefix != "" {
			omitField = append(omitField, subAssetPrefix+".code")
		} else {
			omitField = append(omitField, "code")
		}
	}
	var pos Position
	if req.Position != "" {
		// find position first
		if err := db.DB().
			Where("id = ?", req.Position).
			Take(&pos).
			Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return Asset{}, false, omitField, true, NotFound(c, errors.New("position not found"))
			}
			return Asset{}, false, omitField, true, InternalError(c, err)
		}
	}
	// position insert check
	if pos.ID == "" {
		if subAssetPrefix != "" {
			omitField = append(omitField, subAssetPrefix+".position_id")
		} else {
			omitField = append(omitField, "position_id")
		}
	}
	if req.ID == "" {
		// new
		isNew = true
		assent.ID = strconv.FormatInt(snowFlake.NextVal(), 10)
	} else {
		// edit
		isNew = false
		assent.ID = req.ID
	}
	assent.Type = assetType
	req.UpdateModel(&assent)
	assent.Position = pos
	assent.Pic = req.Pic
	return assent, isNew, omitField, false, nil
}

func createOrUpdate(c echo.Context, asset AssetStorer, isNew bool, omitField []string) (shouldReturn bool, err error) {
	insertRecord := func(asset *Asset, operation AssetOperation, position Position) error {
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

	if isNew {
		if err := db.DB().
			Omit(omitField...).
			Create(asset).
			Error; err != nil {
			return true, InternalError(c, err)
		}
		if err := insertRecord(asset.GetAsset(), AssetOperationCreate, asset.GetAsset().Position); err != nil {
			return true, InternalError(c, err)
		}
	} else {
		if err := db.DB().
			Session(&gorm.Session{FullSaveAssociations: true}).
			Omit(omitField...).
			Updates(asset).
			Error; err != nil {
			return true, InternalError(c, err)
		}
		if err := insertRecord(asset.GetAsset(), AssetOperationPost, asset.GetAsset().Position); err != nil {
			return true, InternalError(c, err)
		}
	}
	return false, nil
}
