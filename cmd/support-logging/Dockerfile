#
# Copyright (c) 2018
# Cavium
#
# SPDX-License-Identifier: Apache-2.0
#
ARG git_sha
FROM edgexfoundry/build-base:${git_sha} AS builder

COPY . .
RUN make cmd/support-logging/support-logging

FROM scratch

LABEL license='SPDX-License-Identifier: Apache-2.0' \
      copyright='Copyright (c) 2018: Cavium'

COPY --from=builder /bin/bash /bin/bash
COPY --from=builder /go/src/github.com/edgexfoundry/edgex-go/cmd/support-logging/Attribution.txt /
COPY --from=builder /go/src/github.com/edgexfoundry/edgex-go/cmd/support-logging/support-logging /
COPY --from=builder /go/src/github.com/edgexfoundry/edgex-go/cmd/support-logging/res/docker/configuration.toml /res/docker/configuration.toml
ENTRYPOINT ["/support-logging","--registry","--profile=docker","--confdir=/res"]
