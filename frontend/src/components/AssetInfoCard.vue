<template>
  <!-- Empty state -->
  <div
    v-if="!data"
    style="min-height: 400px; display: flex; flex-direction: column; align-items: center; justify-content: center; gap: 16px; color: #64748B;"
  >
    <PackageSearch :size="64" stroke-width="1.5" />
    <p style="font-size: 16px; margin: 0;">等待查询</p>
  </div>

  <!-- Data state -->
  <div v-else style="padding: 0 32px 32px;">
    <div
      style="background: #1E293B; border: 1px solid #334155; border-radius: 12px; padding: 24px;"
    >
      <!-- Card header -->
      <div style="display: flex; align-items: center; justify-content: space-between; margin-bottom: 20px;">
        <h3 style="font-size: 16px; font-weight: 600; color: #F8FAFC; margin: 0;">资产信息</h3>
        <StatusBadge :status="baseAsset.status" />
      </div>

      <!-- Fields grid -->
      <div style="display: flex; gap: 40px; margin-bottom: 20px;">
        <!-- Left column -->
        <div style="flex: 1; display: flex; flex-direction: column; gap: 12px;">
          <FieldRow label="主键" :value="baseAsset.id" />
          <FieldRow label="资产码" :value="baseAsset.code || '—'" />
          <FieldRow label="名称" :value="baseAsset.name" />
          <template v-if="tab === 'general'">
            <FieldRow label="类型" :value="baseAsset.type" />
          </template>
          <template v-if="tab === 'book'">
            <FieldRow label="作者" :value="(data as Book).author" />
            <FieldRow label="出版方" :value="(data as Book).publisher" />
          </template>
          <template v-if="tab === 'cd'">
            <FieldRow label="艺术家" :value="(data as CD).author" />
            <FieldRow label="出品方" :value="(data as CD).publisher" />
          </template>
        </div>

        <!-- Divider -->
        <div style="width: 1px; background: #334155; flex-shrink: 0;" />

        <!-- Right column -->
        <div style="flex: 1; display: flex; flex-direction: column; gap: 12px;">
          <FieldRow label="位置" :value="baseAsset.position?.name || baseAsset.position_id || '—'" />
          <template v-if="tab === 'general'">
            <FieldRow label="更新时间" :value="formatTime(baseAsset.last_update)" />
          </template>
          <template v-if="tab === 'book'">
            <FieldRow label="规格" :value="(data as Book).specifications || '—'" />
            <FieldRow label="语言" :value="formatLanguage((data as Book).language)" />
            <FieldRow label="购入时间" :value="formatDate((data as Book).purchase_time)" />
            <FieldRow label="购入价格" :value="formatPrice((data as Book).purchase_price, (data as Book).price_unit)" />
          </template>
          <template v-if="tab === 'cd'">
            <FieldRow label="年份" :value="String((data as CD).year || '—')" />
            <FieldRow label="语言" :value="formatLanguage((data as CD).language)" />
            <FieldRow label="曲目数" :value="String((data as CD).track || '—')" />
            <FieldRow label="购入时间" :value="formatDate((data as CD).purchase_time)" />
            <FieldRow label="购入价格" :value="formatPrice((data as CD).purchase_price, (data as CD).price_unit)" />
          </template>
        </div>
      </div>

      <!-- Divider -->
      <div style="height: 1px; background: #334155; margin-bottom: 16px;" />

      <!-- Action bar -->
      <div style="display: flex; align-items: center; justify-content: flex-end; gap: 8px;">
        <a-button
          style="background: #7C3AED; border-color: #7C3AED; color: #fff;"
          @click="handleEdit"
        >
          <template #icon><Edit3 :size="14" /></template>
          编辑
        </a-button>
        <a-button
          :disabled="baseAsset.status !== 'I'"
          :style="baseAsset.status !== 'I' ? { opacity: 0.5, cursor: 'not-allowed' } : {}"
          style="background: #F59E0B; border-color: #F59E0B; color: #fff;"
          @click="emit('checkout')"
        >
          <template #icon><LogOut :size="14" /></template>
          出库
        </a-button>
        <a-button
          style="background: #2563EB; border-color: #2563EB; color: #fff;"
          @click="emit('viewRecords')"
        >
          <template #icon><ClipboardList :size="14" /></template>
          操作记录
        </a-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { PackageSearch, Edit3, LogOut, ClipboardList } from 'lucide-vue-next'
import { Message } from '@arco-design/web-vue'
import StatusBadge from './StatusBadge.vue'
import FieldRow from './FieldRow.vue'
import type { Asset, Book, CD } from '@/types'
import { LANGUAGE_LABELS } from '@/types'

const props = defineProps<{
  data: Asset | Book | CD | null
  tab: 'general' | 'book' | 'cd'
}>()

const emit = defineEmits<{
  checkout: []
  viewRecords: []
}>()

const baseAsset = computed<Asset>(() => {
  if (!props.data) return {} as Asset
  const d = props.data as any
  return d.asset ?? d
})

function formatTime(t: string | undefined): string {
  if (!t) return '—'
  try {
    return new Date(t).toLocaleString('zh-CN')
  } catch {
    return t
  }
}

function formatDate(d: string | undefined): string {
  if (!d) return '—'
  // pgtype.Date serialises as { "Time": "...", "InfinityModifier": 0, "Valid": true }
  if (typeof d === 'object') {
    const obj = d as any
    if (!obj.Valid) return '—'
    return obj.Time ? new Date(obj.Time).toLocaleDateString('zh-CN') : '—'
  }
  try {
    return new Date(d).toLocaleDateString('zh-CN')
  } catch {
    return d
  }
}

function formatPrice(value: number | undefined, unit?: string): string {
  if (value === undefined || value === null) return '—'
  if (value === 0) return '—'
  const yuanValue = (value / 100).toFixed(2)
  return `¥ ${yuanValue}${unit ? ` (${unit})` : ''}`
}

function formatLanguage(lang: string | undefined): string {
  if (!lang) return '—'
  return LANGUAGE_LABELS[lang] ?? lang
}

function handleEdit() {
  Message.info('功能暂未实现')
}
</script>
