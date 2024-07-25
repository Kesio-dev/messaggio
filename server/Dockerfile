FROM golang:1.22.5-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main ./cmd/server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 5001
CMD ["./main"]
