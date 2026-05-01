FROM golang:1.24.2-alpine AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o server .

FROM alpine:latest
RUN addgroup -S -g 1000 blurt && adduser -S -u 1000 -G blurt blurt
WORKDIR /app
COPY --from=builder /build/server .
USER blurt
EXPOSE 3320
CMD ["./server"]
