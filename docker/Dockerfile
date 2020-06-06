FROM golang:1-alpine

LABEL maintainer="endpot@gmail.com"

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./... \
    && go install -v ./... \
    && swag init

CMD ["go", "run main.go"]
