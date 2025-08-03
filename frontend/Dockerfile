FROM node:24 AS builder

WORKDIR /app

ARG VITE_BACKEND_URL
ENV VITE_BACKEND_URL=$VITE_BACKEND_URL

ARG VITE_ALLOW_FROM
ENV VITE_ALLOW_FROM=$VITE_ALLOW_FROM

ARG VITE_ALLOW_TO
ENV VITE_ALLOW_TO=$VITE_ALLOW_TO

COPY package.json package-lock.json ./
RUN npm ci

COPY . .
RUN npm run build

FROM nginx:stable-alpine

RUN apk add --no-cache curl

RUN rm -rf /etc/nginx/conf.d/*
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d/sites-enabled.conf

COPY --from=builder /app/dist /usr/share/nginx/html