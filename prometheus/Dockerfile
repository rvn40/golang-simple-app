FROM golang:alpine3.12 AS builder

RUN mkdir -p /golang-simple-app
WORKDIR /golang-simple-app
ADD . /golang-simple-app

RUN apk update && apk add --no-cache git openssh-client gcc g++ mercurial \
    tzdata ca-certificates curl 

RUN go get github.com/prometheus/client_golang/prometheus && go get github.com/prometheus/client_golang/prometheus/promhttp
	

RUN go build -o /golang-simple-app/prometheus/main ./prometheus

FROM alpine:3.12

COPY --from=builder /golang-simple-app/prometheus/main /golang-simple-app/prometheus/main

EXPOSE 9110

CMD ./golang-simple-app/prometheus/main 
