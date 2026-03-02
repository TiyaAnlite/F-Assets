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
import type { AssetStatus } from '@/types'
import { STATUS_LABELS, STATUS_COLORS } from '@/types'

const props = defineProps<{
  status: AssetStatus
}>()

const label = computed(() => STATUS_LABELS[props.status] ?? props.status)
const color = computed(() => STATUS_COLORS[props.status] ?? '#6B7280')
</script>