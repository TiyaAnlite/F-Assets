<template>
  <nav
    style="height: 56px; background: #0F172A; display: flex; align-items: center; padding: 0 32px; gap: 8px; border-bottom: 1px solid #334155;"
  >
    <button
      v-for="tab in tabs"
      :key="tab.value"
      :style="tabStyle(tab.value)"
      style="width: 120px; height: 40px; border-radius: 8px; cursor: pointer; font-size: 14px; transition: all 0.15s;"
      @click="emit('update:modelValue', tab.value)"
    >
      {{ tab.label }}
    </button>
  </nav>
</template>

<script setup lang="ts">
import { computed } from 'vue'

type TabValue = 'general' | 'book' | 'cd'

const props = defineProps<{
  modelValue: TabValue
}>()

const emit = defineEmits<{
  'update:modelValue': [value: TabValue]
}>()

const tabs: { label: string; value: TabValue }[] = [
  { label: '通用', value: 'general' },
  { label: '图书', value: 'book' },
  { label: '专辑', value: 'cd' },
]

function tabStyle(value: TabValue) {
  if (props.modelValue === value) {
    return {
      background: '#2563EB',
      color: '#fff',
      fontWeight: '600',
      border: 'none',
    }
  }
  return {
    background: 'transparent',
    color: '#94A3B8',
    fontWeight: '400',
    border: '1px solid #334155',
  }
}
</script>
