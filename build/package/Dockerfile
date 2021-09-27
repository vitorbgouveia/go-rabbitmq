FROM golang

WORKDIR /go/src/web-api

ADD . /go/src/web-api

RUN go mod tidy

EXPOSE 5000

CMD ["sh", "-c", "go run cmd/go-rabbitmq/main.go"]
