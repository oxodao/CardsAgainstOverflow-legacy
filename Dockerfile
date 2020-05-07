FROM golang:1.14.2-alpine as build
RUN apk add --no-cache gcc musl-dev
RUN mkdir /app
COPY . /app/
WORKDIR /app
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' -o cao .

FROM scratch
COPY --from=build /app/cao /cao
EXPOSE 8000
CMD ["./cao"]
