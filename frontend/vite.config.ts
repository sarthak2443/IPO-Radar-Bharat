import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  // GitHub Pages project sites are served from /<repository-name>/.
  // GitHub Actions exposes the actual repository, so this also works if the
  // project is renamed or deployed from a fork.
  base: process.env.GITHUB_ACTIONS === 'true'
    ? `/${process.env.GITHUB_REPOSITORY?.split('/')[1] ?? 'IPO-Radar'}/`
    : '/',
})
