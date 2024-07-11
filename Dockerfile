FROM node:14 AS BUILDFRONT

WORKDIR /app
COPY ./frontend /app

RUN npm i
RUN npm run build

FROM golang:1.22-alpine AS BUILDBACK

WORKDIR /app
COPY . /app

COPY --from=BUILDFRONT /app/dist/ /app/data/

RUN go mod tidy
RUN go mod vendor
RUN go mod download

RUN go build -o cao

FROM alpine

COPY --from=BUILDBACK /app/cao /cao

ENTRYPOINT [ "/cao" ]
