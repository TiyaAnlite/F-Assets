<template>
  <header
    class="h-16 md:h-20 bg-dark px-4 md:px-8 flex items-center justify-between border-b border-color"
  >
    <span class="text-lg md:text-xl font-semibold text-primary tracking-wide">
      F-Assets
    </span>
    <button
      :style="recordsBtnStyle"
      class="flex items-center gap-1.5 px-3 md:px-4 py-2 rounded-lg cursor-pointer text-sm md:text-base font-medium transition-opacity duration-150"
      @click="emit('clickRecords')"
    >
      <History :size="16" />
      <span class="hidden sm:inline">操作记录</span>
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
