# Common Dockerfile for commands using ObjectBox as data storage
# Copyright (c) 2017-2019
# Mainflux
# Cavium
# ObjectBox
#
# SPDX-License-Identifier: Apache-2.0
#

FROM golang:1.11-stretch AS builder
ARG service

RUN apt-get update && apt-get install -y libzmq3-dev

ENV GO111MODULE=on
WORKDIR /go/src/github.com/edgexfoundry/edgex-go

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./objectbox-download.sh .
RUN ./objectbox-download.sh --install 0.5

COPY . .

RUN make cmd/$service/$service

# find all libraries linked to the service binary & copy them to the /deps folder, used bellow
RUN ldd cmd/$service/$service | tr -s '[:blank:]' '\n' | grep '^/' | \
    xargs -I % sh -c 'mkdir -p $(dirname /deps%); cp % /deps%;'

FROM scratch
ARG service

LABEL license='SPDX-License-Identifier: Apache-2.0' \
      copyright='Copyright (c) 2017-2019: Mainflux, Cavium, Dell, ObjectBox'

# service
COPY --from=builder /go/src/github.com/edgexfoundry/edgex-go/cmd/$service/Attribution.txt /
COPY --from=builder /go/src/github.com/edgexfoundry/edgex-go/cmd/$service/$service /service
COPY --from=builder /go/src/github.com/edgexfoundry/edgex-go/cmd/$service/res/docker/configuration.toml /res/docker/configuration.toml

# all service binary dependencies (shared objects)
COPY --from=builder /deps /

# run the service using ld.so
ENV LD_LIBRARY_PATH /usr/local/lib:/usr/lib:/lib:/usr/lib64:/lib64
COPY --from=builder /lib64/ld-linux-x86-64.so.2 /lib64/
ENTRYPOINT ["/lib64/ld-linux-x86-64.so.2","/service","--registry","--profile=docker","--confdir=/res"]