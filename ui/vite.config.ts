import {defineConfig} from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig({
    plugins: [react()],
    server: {
        allowedHosts: true,
        host: "0.0.0.0",
        port: 3000,
        proxy: {
            "/api/v1/auth": {
                target: "http://acta-auth-app-development:10000",
                changeOrigin: true,
            },
            "/api/v1/journal": {
                target: "http://acta-journal-app-development:10010",
                changeOrigin: true,
            },
        },
    },
});
