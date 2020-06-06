FROM golang:1-alpine

LABEL maintainer="endpot@gmail.com"

WORKDIR /go/src/app
COPY . .

RUN go get -v ./...

CMD ["/bin/sh", "-c", "go run main.go"]
