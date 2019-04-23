###############################################################################
# Copyright 2018 Dell Technologies, Inc All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
###############################################################################

# Docker image for building EdgeX Foundry Config Seed
ARG git_sha
FROM objectboxio/edge-build-base:${git_sha} AS build-env

COPY . .

# build
RUN make prepare
RUN make cmd/config-seed/config-seed

# Consul Docker image for EdgeX Foundry
FROM golang:1.11-alpine

LABEL license='SPDX-License-Identifier: Apache-2.0' \
      copyright='Copyright (c) 2017: Samsung'

RUN apk add --no-cache bash

# environment variables
ENV APP_DIR=/edgex/cmd/config-seed

# set the working directory
WORKDIR $APP_DIR

# copy files
COPY --from=build-env $GOPATH/src/github.com/objectbox/edgex-objectbox/cmd/config-seed/Attribution.txt  .
COPY --from=build-env $GOPATH/src/github.com/objectbox/edgex-objectbox/cmd/config-seed/config-seed  .
COPY --from=build-env $GOPATH/src/github.com/objectbox/edgex-objectbox/cmd/config-seed/res ./res
COPY --from=build-env $GOPATH/src/github.com/objectbox/edgex-objectbox/cmd/core-command/res /edgex/cmd/core-command/res
COPY --from=build-env $GOPATH/src/github.com/objectbox/edgex-objectbox/cmd/core-data/res /edgex/cmd/core-data/res
COPY --from=build-env $GOPATH/src/github.com/objectbox/edgex-objectbox/cmd/core-metadata/res /edgex/cmd/core-metadata/res
COPY --from=build-env $GOPATH/src/github.com/objectbox/edgex-objectbox/cmd/export-client/res /edgex/cmd/export-client/res
COPY --from=build-env $GOPATH/src/github.com/objectbox/edgex-objectbox/cmd/export-distro/res /edgex/cmd/export-distro/res
COPY --from=build-env $GOPATH/src/github.com/objectbox/edgex-objectbox/cmd/support-logging/res /edgex/cmd/support-logging/res
COPY --from=build-env $GOPATH/src/github.com/objectbox/edgex-objectbox/cmd/support-notifications/res /edgex/cmd/support-notifications/res
COPY --from=build-env $GOPATH/src/github.com/objectbox/edgex-objectbox/cmd/support-scheduler/res /edgex/cmd/support-scheduler/res
COPY --from=build-env $GOPATH/src/github.com/objectbox/edgex-objectbox/cmd/sys-mgmt-agent/res /edgex/cmd/sys-mgmt-agent/res

ENTRYPOINT /edgex/cmd/config-seed/config-seed --profile=docker --cmd=/edgex/cmd
