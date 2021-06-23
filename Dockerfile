FROM node:16.3-alpine3.13 AS build_frontend

WORKDIR /app
COPY ./frontend /app

RUN yarn
RUN yarn build

FROM golang:1.16-alpine AS build_backend

WORKDIR /app
COPY ./backend /app

COPY --from=build_frontend /app/dist/ /app/web/

RUN go mod tidy
RUN go mod vendor
RUN go mod download

RUN go build -o cao

FROM alpine AS backend

COPY --from=build_backend /app/cao /cao

ENTRYPOINT [ "/cao" ]
