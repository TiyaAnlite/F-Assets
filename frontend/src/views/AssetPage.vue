<template>
  <div class="min-h-screen bg-darker">
    <AppHeader @click-records="handleGlobalRecords" />
    <NavTabs v-model="activeTab" />
    <SearchSection
      ref="searchSectionRef"
      v-model="searchInput"
      :active-mode="activeMode"
      :position-name="positionName"
      @search="handleSearch"
      @mode-click="handleModeClick"
    />

    <div v-if="errorMsg" class="px-4 md:px-8 pb-4">
      <div
        class="bg-red-900/30 border border-red-500 rounded-lg md:rounded-xl px-4 py-3 text-red-500 text-sm"
      >
        {{ errorMsg }}
      </div>
    </div>

    <AssetInfoCard
      :data="currentAsset"
      :tab="activeTab"
      @checkout="handleCheckout"
      @view-records="handleViewRecords"
    />

    <PositionModal
      v-model="showPositionModal"
      @confirm="handlePositionConfirm"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, watch, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { Message } from '@arco-design/web-vue'
import AppHeader from '@/components/AppHeader.vue'
import NavTabs from '@/components/NavTabs.vue'
import SearchSection from '@/components/SearchSection.vue'
import AssetInfoCard from '@/components/AssetInfoCard.vue'
import PositionModal from '@/components/PositionModal.vue'
import { useAssetMode } from '@/composables/useAssetMode'
import { api } from '@/api'
import type { Asset, Book, CD } from '@/types'

type TabValue = 'general' | 'book' | 'cd'

const router = useRouter()
const activeTab = ref<TabValue>('general')
const searchInput = ref('')
const currentAsset = ref<Asset | Book | CD | null>(null)
const errorMsg = ref('')
const showPositionModal = ref(false)

const { activeMode, positionId, positionName, toggleMode, confirmInbound, clearMode } = useAssetMode()

const searchSectionRef = ref<InstanceType<typeof SearchSection> | null>(null)

// Clear state on tab switch
watch(activeTab, () => {
  currentAsset.value = null
  searchInput.value = ''
  errorMsg.value = ''
  clearMode()
  nextTick(() => searchSectionRef.value?.focus())
})

function getAssetType(tab: TabValue) {
  if (tab === 'book') return 'BOOK' as const
  if (tab === 'cd') return 'CD' as const
  return undefined
}

function getBaseAsset(result: Asset | Book | CD): Asset {
  const r = result as any
  return r.asset ?? r
}

async function handleSearch(query: string) {
  const q = query.trim()
  if (!q) return
  errorMsg.value = ''

  try {
    const type = getAssetType(activeTab.value)
    const result = await api.getAsset(q, type)

    if (activeMode.value !== null) {
      const base = getBaseAsset(result as Asset | Book | CD)
      const pos = activeMode.value === 'I' ? positionId.value : base.position_id
      await api.performAction(base.id, activeMode.value, pos)
      // Refresh
      const refreshed = await api.getAsset(base.id, type)
      currentAsset.value = refreshed as Asset | Book | CD
      Message.success('操作成功')
      clearMode()
    } else {
      currentAsset.value = result as Asset | Book | CD
    }
  } catch (err: any) {
    const msg =
      err?.response?.data?.message ||
      err?.response?.data?.msg ||
      err?.message ||
      '查询失败'
    errorMsg.value = msg
    currentAsset.value = null
  } finally {
    searchSectionRef.value?.focus()
  }
}

function handleModeClick(mode: 'I' | 'O' | 'A') {
  const needsModal = toggleMode(mode)
  if (needsModal) {
    showPositionModal.value = true
  }
}

function handlePositionConfirm(position: { id: string; name: string }) {
  confirmInbound(position.id, position.name)
  nextTick(() => searchSectionRef.value?.focus())
}

function handleCheckout() {
  if (!currentAsset.value) return
  const base = getBaseAsset(currentAsset.value)
  toggleMode('O')
  handleSearch(base.id)
}

function handleViewRecords() {
  if (!currentAsset.value) return
  const base = getBaseAsset(currentAsset.value)
  router.push({
    path: `/records/${base.id}`,
    query: {
      assetName: base.name,
      assetCode: base.code,
    },
  })
}

function handleGlobalRecords() {
  if (currentAsset.value) {
    handleViewRecords()
  } else {
    Message.info('请先查询一个资产以查看操作记录')
  }
}
</script>