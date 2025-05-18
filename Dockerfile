FROM golang:1.24.2-alpine AS build
RUN apk --no-cache add ca-certificates zstd build-base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -o /main ./cmd/main.go

FROM ubuntu
WORKDIR /app
RUN apt-get update && apt-get install -y ca-certificates zstd wget
COPY --from=build /main /main
RUN useradd -u 1001 -s /bin/bash manager
USER 1001
ENTRYPOINT ["/main"]