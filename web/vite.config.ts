import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
  server: {
    proxy: {
      "/go" : "http://localhost:3000", // it works like prefix (/go*)
      "/link" : "http://localhost:3000",
      "/tag" : "http://localhost:3000",
    }
  }
})
