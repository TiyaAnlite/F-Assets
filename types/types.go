package types

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type ItemOptRequest struct {
	Type     AssetType `json:"type" validate:"required"`
	Code     string    `json:"code"`
	Name     string    `json:"name" validate:"required"`
	Position string    `json:"position"`
	Pic      string    `json:"pic"`
}

type BookOptRequest struct {
	ItemOptRequest
	Author         string      `json:"author" validate:"required"`
	Publisher      string      `json:"publisher" validate:"required"`
	Specifications string      `json:"specifications" validate:"required"`
	Tag            string      `json:"tag"`
	Language       string      `json:"language" validate:"required"`
	PurchaseTime   pgtype.Date `json:"purchase_time"`
}
