FROM node:14 AS BUILDFRONT

WORKDIR /app
COPY ./frontend /app

RUN yarn
RUN yarn build

FROM golang:1.14 AS BUILDBACK

WORKDIR /app
COPY . /app

RUN go mod vendor
RUN go get github.com/markbates/pkger/cmd/pkger

COPY --from=BUILDFRONT /app/dist/ /app/data/
RUN pkger

RUN go build -o cao

FROM debian:bullseye-20200803-slim

MAINTAINER Nathan JANCZEWSKI <nathan@janczewski.fr>

COPY --from=BUILDBACK /app/cao /cao

ENTRYPOINT [ "/cao" ]
