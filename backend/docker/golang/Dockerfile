FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o apiserver .

FROM alpine:latest AS runner

WORKDIR /app

COPY --from=builder /app/apiserver .

EXPOSE 5050

RUN adduser -D nonroot

USER nonroot

CMD ["./apiserver"]


