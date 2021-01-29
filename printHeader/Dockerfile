FROM golang:alpine3.12 AS builder

RUN mkdir -p /golang-simple-app
WORKDIR /golang-simple-app
ADD . /golang-simple-app

RUN apk update && apk add --no-cache git openssh-client gcc g++ mercurial \
    tzdata ca-certificates curl 

RUN go get github.com/olekukonko/tablewriter && go get github.com/prometheus/client_golang/prometheus

RUN go build -o /golang-simple-app/printHeader/main ./printHeader 

FROM alpine:3.12

COPY --from=builder /golang-simple-app/printHeader/main /golang-simple-app/printHeader/main

EXPOSE 8080 

CMD ./golang-simple-app/printHeader/main 
