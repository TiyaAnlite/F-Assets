<template>
  <div style="min-height: 100vh; background: #020617;">
    <AppHeader records-variant="outline" @click-records="router.back()" />

    <div style="padding: 24px 32px;">
      <!-- Page title -->
      <div style="margin-bottom: 20px;">
        <h2 style="font-size: 20px; font-weight: 600; color: #F8FAFC; margin: 0 0 4px;">操作记录</h2>
        <p style="color: #64748B; font-size: 14px; margin: 0;">
          {{ assetName || props.id }}
          <span v-if="assetCode" style="color: #475569; margin-left: 8px;">{{ assetCode }}</span>
        </p>
      </div>

      <!-- Filter -->
      <div style="display: flex; align-items: center; gap: 12px; margin-bottom: 16px;">
        <span style="color: #94A3B8; font-size: 14px;">操作类型</span>
        <a-select
          v-model="filterOperation"
          style="width: 140px;"
          @change="currentPage = 1"
        >
          <a-option value="">全部</a-option>
          <a-option v-for="op in operationOptions" :key="op.value" :value="op.value">
            {{ op.label }}
          </a-option>
        </a-select>
      </div>

      <!-- Table -->
      <div
        style="background: #1E293B; border: 1px solid #334155; border-radius: 12px; overflow: hidden;"
      >
        <!-- Loading -->
        <div
          v-if="loading"
          style="display: flex; justify-content: center; align-items: center; height: 200px; color: #64748B;"
        >
          <a-spin />
        </div>

        <!-- Error -->
        <div
          v-else-if="fetchError"
          style="display: flex; justify-content: center; align-items: center; height: 200px; color: #DC2626; font-size: 14px;"
        >
          {{ fetchError }}
        </div>

        <!-- Table content -->
        <template v-else>
          <!-- Header -->
          <div
            style="display: grid; grid-template-columns: 200px 80px 1fr 180px; padding: 12px 20px; border-bottom: 1px solid #334155;"
          >
            <span style="color: #94A3B8; font-size: 13px; font-weight: 500;">记录 ID</span>
            <span style="color: #94A3B8; font-size: 13px; font-weight: 500;">操作</span>
            <span style="color: #94A3B8; font-size: 13px; font-weight: 500;">位置</span>
            <span style="color: #94A3B8; font-size: 13px; font-weight: 500;">时间</span>
          </div>

          <!-- Rows -->
          <div
            v-for="record in paginatedRecords"
            :key="record.id"
            style="display: grid; grid-template-columns: 200px 80px 1fr 180px; padding: 14px 20px; border-bottom: 1px solid #1E293B; transition: background 0.1s;"
            @mouseenter="(e) => ((e.currentTarget as HTMLElement).style.background = '#263248')"
            @mouseleave="(e) => ((e.currentTarget as HTMLElement).style.background = 'transparent')"
          >
            <span style="color: #94A3B8; font-size: 13px; font-family: monospace; word-break: break-all;">
              {{ record.id }}
            </span>
            <span>
              <OperationBadge :operation="record.operation" />
            </span>
            <span style="color: #F8FAFC; font-size: 13px;">
              {{ record.position?.name || record.position_id || '—' }}
            </span>
            <span style="color: #64748B; font-size: 13px;">
              {{ formatTime(record.time) }}
            </span>
          </div>

          <!-- Empty state -->
          <div
            v-if="filteredRecords.length === 0"
            style="display: flex; justify-content: center; align-items: center; height: 160px; color: #64748B; font-size: 14px;"
          >
            暂无记录
          </div>
        </template>
      </div>

      <!-- Pagination -->
      <div
        v-if="filteredRecords.length > pageSize"
        style="display: flex; justify-content: flex-end; margin-top: 16px;"
      >
        <a-pagination
          v-model:current="currentPage"
          :total="filteredRecords.length"
          :page-size="pageSize"
          :show-total="true"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import AppHeader from '@/components/AppHeader.vue'
import OperationBadge from '@/components/OperationBadge.vue'
import { api } from '@/api'
import type { AssetRecord } from '@/types'
import { OPERATION_LABELS } from '@/types'

const props = defineProps<{
  id: string
}>()

const router = useRouter()
const route = useRoute()

const assetName = computed(() => route.query.assetName as string | undefined)
const assetCode = computed(() => route.query.assetCode as string | undefined)

const records = ref<AssetRecord[]>([])
const loading = ref(true)
const fetchError = ref('')
const filterOperation = ref<string>('')
const currentPage = ref(1)
const pageSize = 10

const operationOptions = Object.entries(OPERATION_LABELS).map(([value, label]) => ({
  value,
  label,
}))

const filteredRecords = computed(() => {
  if (!filterOperation.value) return records.value
  return records.value.filter((r) => r.operation === filterOperation.value)
})

const paginatedRecords = computed(() => {
  const start = (currentPage.value - 1) * pageSize
  return filteredRecords.value.slice(start, start + pageSize)
})

function formatTime(t: string): string {
  try {
    return new Date(t).toLocaleString('zh-CN')
  } catch {
    return t
  }
}

onMounted(async () => {
  try {
    records.value = await api.getRecords(props.id)
  } catch (err: any) {
    fetchError.value =
      err?.response?.data?.message ||
      err?.response?.data?.msg ||
      err?.message ||
      '获取记录失败'
  } finally {
    loading.value = false
  }
})
</script>
