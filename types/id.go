package types

import (
	"errors"
	"gorm.io/gorm"
)

const (
	UnionIDType RequestIDType = "union"
	EAN13IDType RequestIDType = "ean13"
)

type RequestIDType string

// RequestID 用于同时支持ID和code的兼容支持
type RequestID struct {
	identifier string
	idType     RequestIDType
}

func NewRequestID(identifier string) (*RequestID, error) {
	switch len(identifier) {
	case 13:
		// EAN13 not checked
		return &RequestID{identifier: identifier, idType: EAN13IDType}, nil
	case 0:
		return nil, errors.New("no identifier")
	default:
		return &RequestID{identifier: identifier, idType: UnionIDType}, nil
	}
}

func (r *RequestID) Identifier() string {
	return r.identifier
}

func (r *RequestID) IDType() RequestIDType {
	return r.idType
}
func (r *RequestID) QueryScope(db *gorm.DB) *gorm.DB {
	switch r.idType {
	case UnionIDType:
		return db.Where("id = ?", r.identifier)
	case EAN13IDType:
		return db.Where("code = ?", r.identifier)
	default:
		return db
	}
}
