## Build
FROM golang:1.19.3-alpine3.16 AS build
WORKDIR /usr/local/app/
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o ./build/http_service ./cmd/http_service/main.go

## Deploy
FROM alpine:3.16.2
WORKDIR /usr/local/app/
COPY --from=build /usr/local/app/build/http_service ./http_service
ENTRYPOINT ["./http_service"]