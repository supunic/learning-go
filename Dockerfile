FROM golang:1.17.2-alpine

RUN apk add --update && apk add git gcc g++ sqlite

RUN mkdir -p /go/src/app

WORKDIR /go/src/app
