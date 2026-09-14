import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  // GitHub Pages serves this repository from /IPO-Radar/, while local/Vercel
  // previews are served from /. Keeping this conditional avoids broken assets
  // in both environments.
  base: process.env.GITHUB_ACTIONS === 'true' ? '/IPO-Radar/' : '/',
})
