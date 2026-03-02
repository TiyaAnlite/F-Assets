<template>
  <span
    :style="{ backgroundColor: color }"
    class="rounded-full px-3 py-1 text-white text-xs md:text-sm font-medium inline-block whitespace-nowrap"
  >
    {{ label }}
  </span>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { AssetOperation } from '@/types'
import { OPERATION_LABELS, OPERATION_COLORS } from '@/types'

const props = defineProps<{
  operation: AssetOperation
}>()

const label = computed(() => OPERATION_LABELS[props.operation] ?? props.operation)
const color = computed(() => OPERATION_COLORS[props.operation] ?? '#6B7280')
</script>