package types

import (
	"errors"
	"gorm.io/gorm"
	"strings"
)

const (
	UnionIDType RequestIDType = "union"
	EAN13IDType RequestIDType = "ean13"
	OtherIDType RequestIDType = "other"
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
	case 18:
		if strings.HasPrefix(identifier, "63") {
			// len 18, 63 prefix
			return &RequestID{identifier: identifier, idType: UnionIDType}, nil
		}
		return &RequestID{identifier: identifier, idType: OtherIDType}, nil
	case 0:
		return nil, errors.New("no identifier")
	default:
		return &RequestID{identifier: identifier, idType: OtherIDType}, nil
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
	case OtherIDType:
		return db.Where("id = ?", r.identifier).Or("code = ?", r.identifier)
	default:
		return db
	}
}
