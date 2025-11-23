FROM golang:1.25-alpine AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o backend ./cmd/server

FROM alpine:3.19

WORKDIR /root/

COPY --from=builder /app/backend .

CMD ["./backend"]