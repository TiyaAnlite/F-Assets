<template>
  <a-modal
    v-model:visible="visible"
    title="选择入库位置"
    :mask-closable="false"
    @ok="handleConfirm"
    @cancel="handleCancel"
    :ok-loading="loading"
  >
    <div style="display: flex; flex-direction: column; gap: 12px;">
      <p style="color: #64748B; margin: 0; font-size: 14px;">请输入 5 位位置代码</p>
      <a-input
        v-model="positionCode"
        placeholder="例如: 00001"
        maxlength="5"
        :error="!!errorMsg"
        @keydown.enter="handleConfirm"
        allow-clear
      />
      <p v-if="errorMsg" style="color: #DC2626; margin: 0; font-size: 13px;">{{ errorMsg }}</p>
    </div>
  </a-modal>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { api } from '@/api'

const props = defineProps<{
  modelValue: boolean
}>()

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
  confirm: [position: { id: string; name: string }]
}>()

const visible = ref(props.modelValue)
const positionCode = ref('')
const errorMsg = ref('')
const loading = ref(false)

watch(
  () => props.modelValue,
  (v) => {
    visible.value = v
    if (v) {
      positionCode.value = ''
      errorMsg.value = ''
    }
  },
)

watch(visible, (v) => {
  emit('update:modelValue', v)
})

async function handleConfirm() {
  const code = positionCode.value.trim()
  if (!code || code.length !== 5) {
    errorMsg.value = '请输入 5 位位置代码'
    return
  }

  loading.value = true
  errorMsg.value = ''

  try {
    const pos = await api.getPosition(code)
    emit('confirm', { id: pos.id, name: pos.name })
    visible.value = false
  } catch {
    errorMsg.value = '位置不存在'
  } finally {
    loading.value = false
  }
}

function handleCancel() {
  visible.value = false
}
</script>
