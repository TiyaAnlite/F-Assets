<template>
  <span
    :style="{ backgroundColor: color, borderRadius: '100px', padding: '4px 12px', color: '#fff', fontSize: '13px', fontWeight: 500, display: 'inline-block', whiteSpace: 'nowrap' }"
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
