FROM golang:1.18.2-alpine as builder

WORKDIR /var/www/html

COPY . /var/www/html

RUN go mod download

RUN go build -o /var/www/html/spa-server main.go

FROM alpine:3.16

WORKDIR /data

COPY --from=builder /var/www/html/spa-server /usr/local/bin/spa-server

RUN chmod +x /usr/local/bin/spa-server

VOLUME /data

ENTRYPOINT ["spa-server"]