import { ref } from 'vue'
import type { ActionMode } from '@/types'

export function useAssetMode() {
  const activeMode = ref<ActionMode>(null)
  const positionId = ref<string>('')
  const positionName = ref<string>('')

  // Returns true if the 'I' (入库) mode needs a position modal
  function toggleMode(mode: 'I' | 'O' | 'A'): boolean {
    if (activeMode.value === mode) {
      // Click active button → deactivate
      clearMode()
      return false
    }

    if (mode === 'I') {
      // Needs modal confirmation before activating
      return true
    }

    activeMode.value = mode
    positionId.value = ''
    positionName.value = ''
    return false
  }

  function confirmInbound(id: string, name: string) {
    activeMode.value = 'I'
    positionId.value = id
    positionName.value = name
  }

  function clearMode() {
    activeMode.value = null
    positionId.value = ''
    positionName.value = ''
  }

  return {
    activeMode,
    positionId,
    positionName,
    toggleMode,
    confirmInbound,
    clearMode,
  }
}
