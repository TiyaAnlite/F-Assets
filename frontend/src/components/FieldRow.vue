<template>
  <div class="flex items-start gap-2">
    <span class="text-secondary text-xs md:text-sm min-w-[70px] md:min-w-[80px] flex-shrink-0">{{ label }}</span>
    <span class="text-primary text-xs md:text-sm break-all">{{ value }}</span>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  label: string
  value: string
}>()
</script>