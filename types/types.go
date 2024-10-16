package types

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type PositionOptRequest struct {
	ID   string `json:"id"`
	Name string `json:"name" validate:"required"`
}

type ItemOptRequest struct {
	ID       string `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name" validate:"required"`
	Position string `json:"position"`
	Pic      string `json:"pic"`
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
