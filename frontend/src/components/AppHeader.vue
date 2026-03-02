<template>
  <header
    style="height: 64px; background: #0F172A; padding: 0 32px; display: flex; align-items: center; justify-content: space-between; border-bottom: 1px solid #334155;"
  >
    <span style="font-size: 20px; font-weight: 600; color: #F8FAFC; letter-spacing: 0.5px;">
      F-Assets
    </span>
    <button
      :style="recordsBtnStyle"
      style="display: flex; align-items: center; gap: 6px; padding: 8px 16px; border-radius: 8px; cursor: pointer; font-size: 14px; font-weight: 500; transition: opacity 0.15s;"
      @click="emit('clickRecords')"
    >
      <History :size="16" />
      操作记录
    </button>
  </header>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { History } from 'lucide-vue-next'

const props = defineProps<{
  recordsVariant?: 'filled' | 'outline'
}>()

const emit = defineEmits<{
  clickRecords: []
}>()

const recordsBtnStyle = computed(() => {
  if (props.recordsVariant === 'outline') {
    return {
      background: 'transparent',
      border: '1px solid #2563EB',
      color: '#2563EB',
    }
  }
  return {
    background: '#2563EB',
    border: 'none',
    color: '#fff',
  }
})
</script>
