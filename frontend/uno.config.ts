import { defineConfig, presetUno, presetAttributify } from 'unocss'

export default defineConfig({
  presets: [
    presetUno(),
    presetAttributify(),
  ],
  theme: {
    colors: {
      'bg-darker': '#020617',
      'bg-dark': '#0F172A',
      'bg-card': '#1E293B',
      'accent-blue': '#2563EB',
      'border-color': '#334155',
      'text-primary': '#F8FAFC',
      'text-secondary': '#94A3B8',
      'text-muted': '#64748B',
    },
  },
})
