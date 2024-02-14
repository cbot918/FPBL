import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      // Using the '/api' prefix as an example. Requests to '/api/xxx' will be proxied to 'https://some.server.com/xxx'.
      '/api': {
        target: 'http://localhost:3001',
        changeOrigin: true, // needed for virtual hosted sites
        rewrite: (path) => path.replace(/^\/api/, '') // rewrite the paths to remove the /api prefix
      },
    },
  },
})
