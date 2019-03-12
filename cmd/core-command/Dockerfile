#
# Copyright (c) 2018
# Dell
# Cavium
#
# SPDX-License-Identifier: Apache-2.0
#

ARG git_sha
FROM edgexfoundry/build-base:${git_sha} AS builder

COPY . .

RUN make cmd/core-command/core-command

FROM scratch

LABEL license='SPDX-License-Identifier: Apache-2.0' \
      copyright='Copyright (c) 2018: Dell, Cavium'

ENV APP_PORT=48082
#expose command data port
EXPOSE $APP_PORT

WORKDIR /
COPY --from=builder /go/src/github.com/edgexfoundry/edgex-go/cmd/core-command/Attribution.txt /
COPY --from=builder /go/src/github.com/edgexfoundry/edgex-go/cmd/core-command/core-command /
COPY --from=builder /go/src/github.com/edgexfoundry/edgex-go/cmd/core-command/res/docker/configuration.toml /res/docker/configuration.toml
ENTRYPOINT ["/core-command","--registry","--profile=docker","--confdir=/res"]
