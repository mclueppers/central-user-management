FROM golang:1.18.1-alpine as builder

WORKDIR /app

COPY src/cum/go.* ./

RUN go mod download

COPY src/cum/. .

RUN go build -o main ./cum/cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY .docker/entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

CMD ["./main"]
