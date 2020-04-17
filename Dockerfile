FROM golang:1.13.8-alpine3.11

RUN apk upgrade \
    && apk update \
    && apk add --no-cache git