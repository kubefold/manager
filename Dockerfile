FROM golang:1.24.2-alpine AS build
RUN apk --no-cache add ca-certificates zstd build-base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -o /main ./cmd/main.go

FROM alpine:3.19
WORKDIR /app
RUN apk --no-cache add ca-certificates zstd wget
RUN adduser -u 1001 -D manager
USER 1001
COPY --from=build /main /main
ENTRYPOINT ["/main"]