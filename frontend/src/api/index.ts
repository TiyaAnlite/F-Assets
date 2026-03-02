import axios from 'axios'
import type { Position, Asset, Book, CD, AssetRecord } from '@/types'

const http = axios.create({
  baseURL: '/api',
})

// 响应拦截器：直接返回 NormalResponse 的 data 字段
http.interceptors.response.use(
  (response) => {
    const body = response.data
    if (body && typeof body === 'object' && 'data' in body) {
      return body.data
    }
    return response.data
  },
  (error) => {
    return Promise.reject(error)
  },
)

export const api = {
  getPosition: (id: string): Promise<Position> =>
    http.get(`/position/${id}`),

  listPositions: (): Promise<Position[]> =>
    http.get('/position'),

  getAsset: (id: string, type?: 'BOOK' | 'CD'): Promise<Asset | Book | CD> =>
    http.get(`/asset/${id}`, { params: type ? { type } : {} }),

  performAction: (
    id: string,
    action: 'I' | 'O' | 'A',
    position: string,
  ): Promise<Asset> =>
    http.get(`/action/${id}/${action}`, { params: { position } }),

  getRecords: (id: string): Promise<AssetRecord[]> =>
    http.get(`/record/${id}`),
}
