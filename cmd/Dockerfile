# Common Dockerfile for commands using ObjectBox as data storage
# Copyright (c) 2017-2019
# Mainflux
# Cavium
# ObjectBox
#
# SPDX-License-Identifier: Apache-2.0
#
ARG git_sha
FROM objectboxio/edgex-build-base:${git_sha} AS builder

ARG service
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
COPY --from=builder /go/src/github.com/objectbox/edgex-objectbox/cmd/$service/Attribution.txt /
COPY --from=builder /go/src/github.com/objectbox/edgex-objectbox/cmd/$service/$service /service
COPY --from=builder /go/src/github.com/objectbox/edgex-objectbox/cmd/$service/res/docker/configuration.toml /res/docker/configuration.toml

# all service binary dependencies (shared objects)
COPY --from=builder /deps /

# run the service using ld.so
ENV LD_LIBRARY_PATH /usr/local/lib:/usr/lib:/lib:/usr/lib64:/lib64
COPY --from=builder /ld-linux.so /ld-linux.so
ENTRYPOINT ["/ld-linux.so","/service","--registry","--profile=docker","--confdir=/res"]