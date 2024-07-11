FROM node:22 AS buildfront

WORKDIR /app
COPY ./frontend /app

RUN npm i
RUN npm run build

FROM golang:1.22-alpine AS buildback

WORKDIR /app
COPY . /app

COPY --from=buildfront /app/dist/ /app/frontend/dist/

RUN go mod tidy
RUN go mod vendor
RUN go mod download

RUN go build -o cao

FROM alpine

COPY --from=buildback /app/cao /cao

ENTRYPOINT [ "/cao" ]
