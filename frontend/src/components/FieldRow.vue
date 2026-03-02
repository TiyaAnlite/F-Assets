<template>
  <div style="display: flex; align-items: flex-start; gap: 8px;">
    <span style="color: #94A3B8; font-size: 13px; min-width: 80px; flex-shrink: 0;">{{ label }}</span>
    <span style="color: #F8FAFC; font-size: 13px; word-break: break-all;">{{ value }}</span>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  label: string
  value: string
}>()
</script>
