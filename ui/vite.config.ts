import {defineConfig} from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig({
    plugins: [react()],
    server: {
        allowedHosts: true,
        host: "0.0.0.0",
        port: Number(process.env.APP_PORT ?? 3000),
    },
});
