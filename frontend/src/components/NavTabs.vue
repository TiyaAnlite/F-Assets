<template>
  <nav
    class="h-14 bg-dark flex items-center px-4 md:px-8 py-2 gap-2 border-b border-color overflow-x-auto"
  >
    <button
      v-for="tab in tabs"
      :key="tab.value"
      :style="tabStyle(tab.value)"
      class="h-10 px-4 rounded-lg cursor-pointer text-sm transition-all duration-150 whitespace-nowrap flex-1 md:flex-none md:w-auto min-w-[80px]"
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
