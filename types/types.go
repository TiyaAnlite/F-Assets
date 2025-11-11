package types

import (
	"github.com/duke-git/lancet/v2/maputil"
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

func (r *ItemOptRequest) UpdateModel(m *Asset) {
	m.Code = r.Code
	m.Name = r.Name
}

type BookOptRequest struct {
	ItemOptRequest
	Author         string      `json:"author" validate:"required"`
	Publisher      string      `json:"publisher" validate:"required"`
	Specifications string      `json:"specifications"`
	Tag            string      `json:"tag"`
	Language       string      `json:"language" validate:"required"`
	PurchaseTime   pgtype.Date `json:"purchase_time"`
	Signed         bool        `json:"signed"`
}

func (r *BookOptRequest) UpdateModel(m *Book) {
	m.Author = r.Author
	m.Publisher = r.Publisher
	m.Specifications = r.Specifications
	m.Tag = r.Tag
	m.Language = maputil.GetOrDefault(LanguageStringMapping, r.Language, UnknownLanguage)
	m.PurchaseTime = r.PurchaseTime
	m.Signed = r.Signed
}

type CDOptRequest struct {
	ItemOptRequest
	Author       string      `json:"author" validate:"required"`
	Publisher    string      `json:"publisher" validate:"required"`
	Year         uint32      `json:"year" validate:"required"`
	Language     string      `json:"language" validate:"required"`
	Track        uint32      `json:"track"`
	Tag          string      `json:"tag"`
	PurchaseTime pgtype.Date `json:"purchase_time"`
	Signed       bool        `json:"signed"`
}

func (r *CDOptRequest) UpdateModel(m *CD) {
	m.Author = r.Author
	m.Publisher = r.Publisher
	m.Year = r.Year
	m.Language = maputil.GetOrDefault(LanguageStringMapping, r.Language, UnknownLanguage)
	m.Track = r.Track
	m.Tag = r.Tag
	m.PurchaseTime = r.PurchaseTime
	m.Signed = r.Signed
}

type LanguageType string

const (
	UnknownLanguage            LanguageType = "unknown"
	SimplifiedChineseLanguage  LanguageType = "zh-cn"
	TraditionalChineseLanguage LanguageType = "zh-tw"
	JapaneseLanguage           LanguageType = "jp"
)

var LanguageStringMapping = map[string]LanguageType{
	"unknown": UnknownLanguage,
	"zh-cn":   SimplifiedChineseLanguage,
	"zh-tw":   TraditionalChineseLanguage,
	"jp":      JapaneseLanguage,
}

type AssetStorer interface {
	GetAsset() *Asset
}
