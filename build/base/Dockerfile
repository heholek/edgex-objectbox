# Common Dockerfile loads all the modules
# Copyright (c) 2019 ObjectBox

FROM golang:1.11-stretch
ARG service

RUN apt-get update && apt-get install -y libzmq3-dev

ENV GO111MODULE=on
WORKDIR /go/src/github.com/objectbox/edgex-objectbox

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./build/base/objectbox-download.sh .
RUN ./objectbox-download.sh --install 0.5
