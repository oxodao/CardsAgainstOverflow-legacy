// vite.config.js
import * as path from 'path';

import { defineConfig } from 'vite';
// import vue from '@vitejs/plugin-vue' // Until we migrate to vue3
import vue from '@vitejs/plugin-vue2';

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
            '/api': {
                target: 'ws://backend:8000/',
                ws: true,
                changeOrigin: true,
            },
            '/deporte': {
                target: 'ws://backend:8000/',
                ws: true,
                changeOrigin: true,
            }
        }
    }
});