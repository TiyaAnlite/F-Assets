<template>
  <div class="min-h-screen bg-darker">
    <AppHeader records-variant="outline" @click-records="router.back()" />

    <div class="px-4 md:px-8 py-6">
      <!-- Page title -->
      <div class="mb-5">
        <h2 class="text-lg md:text-xl font-semibold text-primary mb-1">操作记录</h2>
        <p class="text-muted text-sm md:text-base m-0">
          {{ assetName || props.id }}
          <span v-if="assetCode" class="text-gray-600 ml-2">{{ assetCode }}</span>
        </p>
      </div>

      <!-- Filter -->
      <div class="flex flex-col sm:flex-row sm:items-center gap-3 sm:gap-3 mb-4">
        <span class="text-secondary text-sm">操作类型</span>
        <a-select
          v-model="filterOperation"
          class="w-full sm:w-36"
          @change="currentPage = 1"
        >
          <a-option value="">全部</a-option>
          <a-option v-for="op in operationOptions" :key="op.value" :value="op.value">
            {{ op.label }}
          </a-option>
        </a-select>
      </div>

      <!-- Table container with horizontal scroll on mobile -->
      <div
        class="bg-card border border-color rounded-xl md:rounded-2xl overflow-hidden"
      >
        <!-- Loading -->
        <div
          v-if="loading"
          class="flex justify-center items-center h-48 md:h-52 text-muted"
        >
          <a-spin />
        </div>

        <!-- Error -->
        <div
          v-else-if="fetchError"
          class="flex justify-center items-center h-48 md:h-52 text-red-500 text-sm"
        >
          {{ fetchError }}
        </div>

        <!-- Table content -->
        <template v-else>
          <!-- Scrollable wrapper for mobile -->
          <div class="overflow-x-auto">
            <!-- Table -->
            <div class="min-w-[500px] md:min-w-0">
              <!-- Header -->
              <div
                class="grid grid-cols-[1fr_70px_1fr_140px] md:grid-cols-[200px_80px_1fr_180px] px-4 md:px-5 py-3 border-b border-color"
              >
                <span class="text-secondary text-xs md:text-sm font-medium">记录 ID</span>
                <span class="text-secondary text-xs md:text-sm font-medium">操作</span>
                <span class="text-secondary text-xs md:text-sm font-medium">位置</span>
                <span class="text-secondary text-xs md:text-sm font-medium">时间</span>
              </div>

              <!-- Rows -->
              <div
                v-for="record in paginatedRecords"
                :key="record.id"
                class="grid grid-cols-[1fr_70px_1fr_140px] md:grid-cols-[200px_80px_1fr_180px] px-4 md:px-5 py-3 md:py-3.5 border-b border-darker transition-colors duration-100 hover:bg-slate-700/50"
              >
                <span class="text-secondary text-xs md:text-sm font-mono break-all pr-2">
                  {{ record.id }}
                </span>
                <span>
                  <OperationBadge :operation="record.operation" />
                </span>
                <span class="text-primary text-xs md:text-sm truncate pr-2">
                  {{ record.position?.name || record.position_id || '—' }}
                </span>
                <span class="text-muted text-xs md:text-sm">
                  {{ formatTime(record.time) }}
                </span>
              </div>
            </div>
          </div>

          <!-- Empty state -->
          <div
            v-if="filteredRecords.length === 0"
            class="flex justify-center items-center h-40 text-muted text-sm"
          >
            暂无记录
          </div>
        </template>
      </div>

      <!-- Pagination -->
      <div
        v-if="filteredRecords.length > pageSize"
        class="flex justify-end mt-4"
      >
        <a-pagination
          v-model:current="currentPage"
          :total="filteredRecords.length"
          :page-size="pageSize"
          :show-total="true"
          :simple="isMobile"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
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
const isMobile = ref(false)

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

function checkMobile() {
  isMobile.value = window.innerWidth < 640
}

onMounted(() => {
  checkMobile()
  window.addEventListener('resize', checkMobile)
})

onUnmounted(() => {
  window.removeEventListener('resize', checkMobile)
})

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