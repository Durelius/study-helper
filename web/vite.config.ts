import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'

// The build lands straight in the Go module so `go:embed` picks it up and one binary
// serves both the app and the API.
export default defineConfig({
  plugins: [react(), tailwindcss()],
  build: {
    outDir: '../internal/web/dist',
    // The directory has to survive a clean checkout or `go:embed all:dist` fails to
    // compile, so .gitkeep stays put and `npm run build` clears assets/ itself.
    emptyOutDir: false,
  },
  server: {
    port: 5173,
    proxy: { '/api': 'http://127.0.0.1:8080' },
  },
})
