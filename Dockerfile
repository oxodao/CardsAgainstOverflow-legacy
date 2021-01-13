FROM node:14 AS BUILDFRONT

WORKDIR /app
COPY ./frontend /app

RUN yarn
RUN yarn build

FROM golang:1.15-alpine AS BUILDBACK

WORKDIR /app
COPY . /app

RUN go get github.com/markbates/pkger/cmd/pkger

COPY --from=BUILDFRONT /app/dist/ /app/data/

RUN pkger
RUN go mod tidy
RUN go mod vendor
RUN go mod download

RUN go build -o cao

FROM debian:bullseye-20200803-slim

COPY --from=BUILDBACK /app/cao /cao

ENTRYPOINT [ "/cao" ]
