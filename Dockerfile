FROM golang:1.24-alpine AS builder

COPY go.mod go.sum ./
RUN go mod download
COPY . .

WORKDIR /app
COPY . .
RUN go build -o main ./cmd


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 3000
CMD ["./main"]
