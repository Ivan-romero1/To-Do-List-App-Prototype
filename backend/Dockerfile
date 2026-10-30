
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod main.go ./

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o server main.go

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
