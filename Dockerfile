FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o solution-service ./cmd/main.go

FROM alpine:latest

WORKDIR /
COPY --from=builder /app/solution-service .
ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["./solution-service"]
