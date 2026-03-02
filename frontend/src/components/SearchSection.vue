<template>
  <div style="padding: 24px 32px 0;">
    <!-- Title row -->
    <div style="display: flex; align-items: center; justify-content: space-between; margin-bottom: 16px;">
      <h2 style="font-size: 18px; font-weight: 600; color: #F8FAFC; margin: 0;">资产查询</h2>

      <div style="display: flex; align-items: center; gap: 8px;">
        <!-- Position label (shown only when 'I' mode active) -->
        <span
          v-if="activeMode === 'I' && positionName"
          style="font-size: 13px; color: #64748B; margin-right: 8px;"
        >
          位置: {{ positionName }}
        </span>

        <!-- Mode buttons -->
        <button
          v-for="btn in modeButtons"
          :key="btn.mode"
          :style="modeButtonStyle(btn.mode)"
          style="width: 56px; height: 32px; border-radius: 6px; cursor: pointer; font-size: 13px; font-weight: 500; transition: all 0.15s;"
          @click="handleModeClick(btn.mode)"
        >
          {{ btn.label }}
        </button>
      </div>
    </div>

    <!-- Search input -->
    <div style="position: relative; margin-bottom: 24px;">
      <input
        ref="inputRef"
        v-model="localInput"
        type="text"
        placeholder="输入资产 ID 或资产码..."
        :style="inputStyle"
        style="width: 100%; height: 56px; border-radius: 12px; background: #1E293B; color: #F8FAFC; font-size: 15px; padding: 0 56px 0 20px; outline: none; box-sizing: border-box; transition: border-color 0.15s;"
        @keydown.enter="handleEnter"
        @input="emit('update:modelValue', localInput)"
      />
      <button
        style="position: absolute; right: 16px; top: 50%; transform: translateY(-50%); background: none; border: none; cursor: pointer; color: #64748B; display: flex; align-items: center;"
        @click="emit('search', localInput)"
      >
        <Search :size="20" />
      </button>
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
