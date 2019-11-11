# Common Dockerfile loads all the modules
# Copyright (c) 2019 ObjectBox

FROM golang:1.12-stretch
ARG service

RUN apt-get update && apt-get install -y libzmq3-dev

ENV GO111MODULE=on
WORKDIR /go/src/github.com/objectbox/edgex-objectbox

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./build/base/objectbox-download.sh .
RUN ./objectbox-download.sh --install 0.7.1

RUN if [ "$(uname -m)" = "armv7l" ]; \
        then cp /lib/ld-linux-armhf.so.3 /ld-linux.so; \
        else cp /lib64/ld-linux-x86-64.so.2 /ld-linux.so; fi
