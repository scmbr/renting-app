# Build stage
FROM node:20-alpine AS builder
WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .
RUN npm run build

# Run stage
FROM nginx:alpine

# Копируем своё nginx.conf в образ
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Копируем билд
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
