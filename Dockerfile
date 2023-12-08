FROM golang:1.18.3-alpine AS builder

COPY . /github.com/kaspeya/go-url-shortener/
WORKDIR /github.com/kaspeya/go-url-shortener/

RUN go mod download
RUN go build -o ./bin/go-url-shortener cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/kaspeya/go-url-shortener/bin/url_shortener .
COPY --from=builder /github.com/kaspeya/go-url-shortener/config/ /root/config/

EXPOSE 50051
EXPOSE 8080

CMD ["./go-url-shortener"]