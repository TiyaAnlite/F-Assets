package types

import "C"
import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const (
	AssetBasicItemType AssetType = "ITEM"
	AssetBookType      AssetType = "BOOK"
	AssetCDType        AssetType = "CD"
)

const (
	AssetStatusOutbound AssetStatusType = "O" // 出库
	AssetStatusInbound  AssetStatusType = "I" // 在库
	AssetStatusAbandon  AssetStatusType = "A" // 已销毁
)

const (
	AssetOperationCreate  AssetOperation = "C" // 新建
	AssetOperationPost    AssetOperation = "P" // 修改
	AssetOperationEnter   AssetOperation = "E" // 入库
	AssetOperationLeave   AssetOperation = "L" // 出库
	AssetOperationDestroy AssetOperation = "D" // 销毁
)

type AssetType string
type AssetStatusType string
type AssetOperation string

// Asset 资产表，是所有资产的基表
type Asset struct {
	ID         string          `json:"id" gorm:"column:id;primary_key;type:bigint;comment:唯一主键"`                                        // 唯一主键
	Type       AssetType       `json:"type" gorm:"column:type;type:varchar(255);not null;comment:资产类型"`                                 // 资产类型
	Code       string          `json:"code" gorm:"column:code;type:varchar(255);uniqueIndex:idx_code,type:hash;comment:资产码，一物并不一定会有一码"` // 资产码，一物并不一定会有一码
	Name       string          `json:"name" gorm:"column:name;type:varchar(255);not null;comment:资产名"`                                  // 资产名
	Status     AssetStatusType `json:"status" gorm:"column:status;type:char(1);not null;default:O;comment:状态"`                          // 状态
	PositionID string          `json:"position_id" gorm:"column:position_id;type:char(5);comment:位置"`                                   // 位置
	Position   Position        `json:"position"`
	LastUpdate time.Time       `json:"last_update" gorm:"column:last_update;autoUpdateTime;not null;comment:更新时间"` // 更新时间
	Pic        string          `json:"pic" gorm:"column:pic;type:varchar(255);comment:实物图片"`                       // 实物图片
}

func (a *Asset) GetAsset() *Asset {
	return a
}

// Record 资产操作记录
type Record struct {
	ID         string         `json:"id" gorm:"column:id;primary_key;type:bigint;comment:记录主键"`                                    // 记录主键
	AssetID    string         `json:"asset_id" gorm:"column:asset_id;type:bigint;index:idx_record_asset,type:hash;comment:资产唯一主键"` // 对应资产
	Operation  AssetOperation `json:"operation" gorm:"column:operation;type:char(1);not null;comment:操作"`                          // 操作
	PositionID string         `json:"position_id" gorm:"column:position_id;type:char(5);comment:位置"`                               // 位置
	Position   Position       `json:"position" `
	Time       time.Time      `json:"time" gorm:"column:time;autoCreateTime;not null;comment:操作时间"` // 操作时间
}

// Position 资产位置
type Position struct {
	ID         string    `json:"id" gorm:"column:id;primary_key;type:char(5);comment:位置代码(5位纯数字)"`            // 位置代码(5位纯数字)
	Name       string    `json:"name" gorm:"column:name;type:varchar(255);not null;comment:位置名称"`             // 位置名称
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime;not null;comment:创建时间;"` // 创建时间
	LastUpdate time.Time `json:"last_update" gorm:"column:last_update;autoUpdateTime;not null;comment:更新时间"`  // 更新时间
}

// Book 图书类型资产
type Book struct {
	AssetID        string       `json:"asset_id" gorm:"column:asset_id;type:bigint;primary_key;comment:资产唯一主键"` // 资产唯一主键
	Asset          Asset        `json:"asset"`
	Author         string       `json:"author" gorm:"column:author;type:varchar(255);not null;comment:作者"`              // 作者
	Publisher      string       `json:"publisher" gorm:"column:publisher;type:varchar(255);not null;comment:出版/出品/销售方"` // 出版/出品/销售方
	Specifications string       `json:"specifications" gorm:"column:specifications;type:varchar(255);comment:规格(开本)"`   // 规格(开本)
	Tag            string       `json:"tag" gorm:"column:tag;type:varchar(255);comment:标签"`                             // 标签
	Language       LanguageType `json:"language" gorm:"column:language;type:varchar(255);not null;comment:语言"`          // 语言
	PurchaseTime   pgtype.Date  `json:"purchase_time" gorm:"column:purchase_time;comment:购入时间"`                         // 购入时间
	Price          uint32       `json:"price" gorm:"column:price;comment:标价"`                                           // 标价
	PurchasePrice  uint32       `json:"purchase_price" gorm:"column:purchase_price;comment:购入价格"`                       // 购入价格
	PriceUnit      string       `json:"price_unit" gorm:"column:price_unit;type:varchar(255);comment:价格单位"`             // 价格单位
	Signed         bool         `json:"signed" gorm:"column:signed;comment:有无签名"`                                       // 有无签名
}

func (b *Book) GetAsset() *Asset {
	return &b.Asset
}

// CD 专辑类型资产
type CD struct {
	AssetID       string       `json:"asset_id" gorm:"column:asset_id;type:bigint;primary_key;comment:资产唯一主键"` // 资产唯一主键
	Asset         Asset        `json:"asset"`
	Author        string       `json:"author" gorm:"column:author;type:varchar(255);not null;comment:作者"`                 // 作者
	Publisher     string       `json:"publisher" gorm:"column:publisher;type:varchar(255);not null;comment:出版/出品/社团/销售方"` // 出版/出品/社团/销售方
	Year          uint32       `json:"year" gorm:"column:year;not null;comment:年份"`                                       // 年份
	Language      LanguageType `json:"language" gorm:"column:language;type:varchar(255);not null;comment:语言"`             // 语言
	Track         uint32       `json:"track" gorm:"column:track;not null;comment:曲目数"`                                    // 曲目数
	Tag           string       `json:"tag" gorm:"column:tag;type:varchar(255);comment:标签"`                                // 标签
	PurchaseTime  pgtype.Date  `json:"purchase_time" gorm:"column:purchase_time;comment:购入时间"`                            // 购入时间
	Price         uint32       `json:"price" gorm:"column:price;comment:标价"`                                              // 标价
	PurchasePrice uint32       `json:"purchase_price" gorm:"column:purchase_price;comment:购入价格"`
	PriceUnit     string       `json:"price_unit" gorm:"column:price_unit;type:varchar(255);comment:价格单位"`
	Signed        bool         `json:"signed" gorm:"column:signed;comment:有无签名"` // 有无签名
}

func (c *CD) GetAsset() *Asset {
	return &c.Asset
}

type Meta struct {
	AssetID    string    `json:"asset_id" gorm:"column:asset_id;type:bigint;primary_key;comment:资产唯一主键"`     // 资产唯一主键
	Data       string    `json:"data" gorm:"column:data;type:text;not null;comment:元数据"`                     // 元数据
	LastUpdate time.Time `json:"last_update" gorm:"column:last_update;autoUpdateTime;not null;comment:更新时间"` // 更新时间
}
