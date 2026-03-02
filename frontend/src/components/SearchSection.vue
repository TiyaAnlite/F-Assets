<template>
  <div class="px-4 md:px-8 pt-6">
    <!-- Title row -->
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3 mb-4">
      <h2 class="text-base md:text-lg font-semibold text-primary m-0">资产查询</h2>

      <div class="flex flex-wrap items-center gap-2">
        <!-- Position label (shown only when 'I' mode active) -->
        <span
          v-if="activeMode === 'I' && positionName"
          class="text-xs text-muted mr-1"
        >
          位置：{{ positionName }}
        </span>

        <!-- Mode buttons -->
        <button
          v-for="btn in modeButtons"
          :key="btn.mode"
          :style="modeButtonStyle(btn.mode)"
          class="w-14 h-8 rounded-md cursor-pointer text-xs font-medium transition-all duration-150"
          @click="handleModeClick(btn.mode)"
        >
          {{ btn.label }}
        </button>
      </div>
    </div>

    <!-- Search input -->
    <div class="relative mb-6">
      <input
        ref="inputRef"
        v-model="localInput"
        type="text"
        placeholder="输入资产 ID 或资产码..."
        :style="inputStyle"
        class="w-full h-14 rounded-xl bg-card text-primary text-base md:text-lg px-5 pr-14 outline-none box-border transition-border-color"
        @keydown.enter="handleEnter"
        @input="emit('update:modelValue', localInput)"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { Search } from 'lucide-vue-next'
import type { ActionMode } from '@/types'

const props = defineProps<{
  modelValue: string
  activeMode: ActionMode
  positionName: string
}>()

const emit = defineEmits<{
  'update:modelValue': [value: string]
  search: [query: string]
  modeClick: [mode: 'I' | 'O' | 'A']
}>()

const localInput = ref(props.modelValue)
const inputRef = ref<HTMLInputElement | null>(null)

defineExpose({ focus: () => inputRef.value?.focus() })

watch(
  () => props.modelValue,
  (v) => { localInput.value = v },
)

const modeButtons = [
  { mode: 'O' as const, label: '出' },
  { mode: 'I' as const, label: '入' },
  { mode: 'A' as const, label: '销' },
]

const modeColors: Record<'O' | 'I' | 'A', string> = {
  O: '#2563EB',
  I: '#059669',
  A: '#DC2626',
}

function modeButtonStyle(mode: 'I' | 'O' | 'A') {
  if (props.activeMode === mode) {
    return {
      background: modeColors[mode],
      border: 'none',
      color: '#fff',
    }
  }
  return {
    background: '#1E293B',
    border: '1px solid #334155',
    color: '#94A3B8',
  }
}

const inputStyle = computed(() => ({
  border: localInput.value
    ? '1px solid #2563EB'
    : '1px solid #334155',
}))

function handleModeClick(mode: 'I' | 'O' | 'A') {
  emit('modeClick', mode)
}

function handleEnter() {
  const val = localInput.value
  emit('search', val)
  localInput.value = ''
  emit('update:modelValue', '')
}
</script>
