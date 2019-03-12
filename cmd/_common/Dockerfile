# Common Dockerfile loads all the modules
# Copyright (c) 2019 ObjectBox

FROM golang:1.12-stretch
ARG service

RUN apt-get update && apt-get install -y libzmq3-dev

ENV GO111MODULE=on
WORKDIR /go/src/github.com/edgexfoundry/edgex-go

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./objectbox-download.sh .
RUN ./objectbox-download.sh --install 0.5
