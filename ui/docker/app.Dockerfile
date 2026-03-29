FROM node:22-bookworm

WORKDIR /app

COPY package.json ./
COPY package-lock.json ./
COPY tsconfig.json ./
COPY tsconfig.app.json ./
COPY tsconfig.node.json ./
COPY vite.config.ts ./

RUN npm install

COPY index.html ./
COPY src ./src

CMD npm run dev -- --host 0.0.0.0 --port ${APP_PORT}
