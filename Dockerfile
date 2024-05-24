FROM golang:1.21-alpine3.20 as builder

WORKDIR /github.com/rogaliiik/library/
COPY . /github.com/rogaliiik/library/

RUN go mod download
RUN GOOS=linux go build -o ./app ./cmd/app/main.go

FROM alpine:3.20 as app

WORKDIR /github.com/rogaliiik/library/

COPY --from=builder /github.com/rogaliiik/library/app .
COPY --from=builder /github.com/rogaliiik/library/config ./config
COPY --from=builder /github.com/rogaliiik/library/migrations ./migrations

CMD ["./app", "--config", "config/config-compose.yaml"]





