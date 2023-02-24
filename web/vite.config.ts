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
  base: "",
  server: {
    // port: 3000,
    proxy: {
      "/go" : "http://127.0.0.1:3000", // it works like prefix (/go*)
      "/link" : "http://127.0.0.1:3000",
      "/tag" : "http://127.0.0.1:3000",
    }
  }
})
