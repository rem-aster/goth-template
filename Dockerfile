FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd/
COPY internal ./internal/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/server/server.go

FROM scratch
COPY web ./web/
COPY --from=builder /app/main .
CMD ["./main"]