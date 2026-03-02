<template>
  <span
    :style="{ backgroundColor: color, borderRadius: '100px', padding: '6px 12px', color: '#fff', fontSize: '13px', fontWeight: 500, display: 'inline-block', whiteSpace: 'nowrap' }"
  >
    {{ label }}
  </span>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { AssetStatus } from '@/types'
import { STATUS_LABELS, STATUS_COLORS } from '@/types'

const props = defineProps<{
  status: AssetStatus
}>()

const label = computed(() => STATUS_LABELS[props.status] ?? props.status)
const color = computed(() => STATUS_COLORS[props.status] ?? '#6B7280')
</script>
