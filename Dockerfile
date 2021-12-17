FROM golang:1.17.5-alpine3.15 AS builder

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go build app.go

CMD ["./app"]
