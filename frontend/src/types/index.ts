export type AssetStatus = 'I' | 'O' | 'A'
export type AssetType = 'ITEM' | 'BOOK' | 'CD'
export type AssetOperation = 'C' | 'P' | 'E' | 'L' | 'D'
export type ActionMode = 'I' | 'O' | 'A' | null

export interface Position {
  id: string
  name: string
  create_time: string
  last_update: string
}

export interface Asset {
  id: string
  type: AssetType
  code: string
  name: string
  status: AssetStatus
  position_id: string
  position: Position
  last_update: string
  pic: string
}

export interface Book {
  asset_id: string
  asset: Asset
  author: string
  publisher: string
  specifications: string
  tag: string
  language: string
  purchase_time: string
  price: number
  purchase_price: number
  price_unit: string
  signed: boolean
}

export interface CD {
  asset_id: string
  asset: Asset
  author: string
  publisher: string
  year: number
  language: string
  track: number
  tag: string
  purchase_time: string
  price: number
  purchase_price: number
  price_unit: string
  signed: boolean
}

export interface AssetRecord {
  id: string
  asset_id: string
  operation: AssetOperation
  position_id: string
  position: Position
  time: string
}

export const OPERATION_LABELS: { [K in AssetOperation]: string } = {
  C: '新建',
  P: '修改',
  E: '入库',
  L: '出库',
  D: '销毁',
}

export const OPERATION_COLORS: { [K in AssetOperation]: string } = {
  C: '#059669',
  P: '#8B5CF6',
  E: '#2563EB',
  L: '#F59E0B',
  D: '#DC2626',
}

export const STATUS_LABELS: { [K in AssetStatus]: string } = {
  I: '在库',
  O: '出库',
  A: '已销毁',
}

export const STATUS_COLORS: { [K in AssetStatus]: string } = {
  I: '#2563EB',
  O: '#F59E0B',
  A: '#6B7280',
}

export const LANGUAGE_LABELS: { [key: string]: string } = {
  unknown: '未知',
  'zh-cn': '简体中文',
  'zh-tw': '繁体中文',
  jp: '日语',
}
