import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      "/go" : "http://localhost:3000",
      "/link" : "http://localhost:3000",
      "/tag" : "http://localhost:3000",
    }
  }
})
