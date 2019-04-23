#
# Copyright (c) 2018 Cavium
#
# SPDX-License-Identifier: Apache-2.0
#


.PHONY: build clean test docker run


GO=CGO_ENABLED=0 GO111MODULE=on go
GOCGO=CGO_ENABLED=1 GO111MODULE=on go

DOCKERS=docker__common docker_config_seed docker_export_client docker_export_distro docker_core_data docker_core_metadata docker_core_command docker_support_logging docker_support_notifications docker_sys_mgmt_agent docker_support_scheduler
.PHONY: $(DOCKERS)

MICROSERVICES=cmd/config-seed/config-seed cmd/export-client/export-client cmd/export-distro/export-distro cmd/core-metadata/core-metadata cmd/core-data/core-data cmd/core-command/core-command cmd/support-logging/support-logging cmd/support-notifications/support-notifications cmd/sys-mgmt-executor/sys-mgmt-executor cmd/sys-mgmt-agent/sys-mgmt-agent cmd/support-scheduler/support-scheduler

.PHONY: $(MICROSERVICES)

VERSION=$(shell cat ./VERSION)
VERSION_SUFFIX=-dev

GOFLAGS=-ldflags "-X github.com/objectbox/edgex-objectbox.Version=$(VERSION)"

GIT_SHA=$(shell git rev-parse HEAD)

build: $(MICROSERVICES)

cmd/config-seed/config-seed:
	$(GO) build $(GOFLAGS) -o $@ ./cmd/config-seed

cmd/core-metadata/core-metadata:
	$(GOCGO) build $(GOFLAGS) -o $@ ./cmd/core-metadata

cmd/core-data/core-data:
	$(GOCGO) build $(GOFLAGS) -o $@ ./cmd/core-data

cmd/core-command/core-command:
	$(GO) build $(GOFLAGS) -o $@ ./cmd/core-command

cmd/export-client/export-client:
	$(GOCGO) build $(GOFLAGS) -o $@ ./cmd/export-client

cmd/export-distro/export-distro:
	$(GOCGO) build $(GOFLAGS) -o $@ ./cmd/export-distro

cmd/support-logging/support-logging:
	$(GO) build $(GOFLAGS) -o $@ ./cmd/support-logging

cmd/support-notifications/support-notifications:
	$(GOCGO) build $(GOFLAGS) -o $@ ./cmd/support-notifications

cmd/sys-mgmt-executor/sys-mgmt-executor:
	$(GO) build $(GOFLAGS) -o $@ ./cmd/sys-mgmt-executor

cmd/sys-mgmt-agent/sys-mgmt-agent:
	$(GO) build $(GOFLAGS) -o $@ ./cmd/sys-mgmt-agent

cmd/support-scheduler/support-scheduler:
	$(GOCGO) build $(GOFLAGS) -o $@ ./cmd/support-scheduler

clean:
	rm -f $(MICROSERVICES)

test:
	GO111MODULE=on go test -cover ./...
	GO111MODULE=on go vet ./...

prepare:

run:
	cd bin && ./edgex-launch.sh

run_docker:
	cd bin && ./edgex-docker-launch.sh

docker: $(DOCKERS)

docker__common:
	docker build \
		-f cmd/_common/Dockerfile \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edge-build-base:$(GIT_SHA) \
		.

docker_config_seed:
	docker build \
		-f cmd/config-seed/Dockerfile \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edge-core-config-seed:$(GIT_SHA) \
		-t objectboxio/edge-core-config-seed:$(VERSION)$(VERSION_SUFFIX) \
		.

docker_core_metadata:
	docker build \
		-f cmd/Dockerfile --build-arg service=core-metadata \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edge-core-metadata:$(GIT_SHA) \
		-t objectboxio/edge-core-metadata:$(VERSION)$(VERSION_SUFFIX) \
		.

docker_core_data:
	docker build \
		-f cmd/Dockerfile --build-arg service=core-data \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edge-core-data:$(GIT_SHA) \
		-t objectboxio/edge-core-data:$(VERSION)$(VERSION_SUFFIX) \
		.

docker_core_command:
	docker build \
		-f cmd/core-command/Dockerfile \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edge-core-command:$(GIT_SHA) \
		-t objectboxio/edge-core-command:$(VERSION)$(VERSION_SUFFIX) \
		.

docker_export_client:
	docker build \
		-f cmd/Dockerfile --build-arg service=export-client \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edge-export-client:$(GIT_SHA) \
		-t objectboxio/edge-export-client:$(VERSION)$(VERSION_SUFFIX) \
		.

docker_export_distro:
	docker build \
		-f cmd/export-distro/Dockerfile \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edge-export-distro:$(GIT_SHA) \
		-t objectboxio/edge-export-distro:$(VERSION)$(VERSION_SUFFIX) \
		.

docker_support_logging:
	docker build \
		-f cmd/support-logging/Dockerfile \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edge-support-logging:$(GIT_SHA) \
		-t objectboxio/edge-support-logging:$(VERSION)$(VERSION_SUFFIX) \
		.

docker_support_notifications:
	docker build \
		-f cmd/Dockerfile --build-arg service=support-notifications \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edge-support-notifications:$(GIT_SHA) \
		-t objectboxio/edge-support-notifications:$(VERSION)$(VERSION_SUFFIX) \
		.

docker_support_scheduler:
	docker build \
		-f cmd/Dockerfile --build-arg service=support-scheduler \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edge-support-scheduler:$(GIT_SHA) \
		-t objectboxio/edge-support-scheduler:$(VERSION)$(VERSION_SUFFIX) \
		.

docker_sys_mgmt_agent:
	docker build \
		-f cmd/sys-mgmt-agent/Dockerfile \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edge-sys-mgmt-agent:$(GIT_SHA) \
		-t objectboxio/edge-sys-mgmt-agent:$(VERSION)$(VERSION_SUFFIX) \
		.
