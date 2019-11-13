#
# Copyright (c) 2018 Cavium
#
# SPDX-License-Identifier: Apache-2.0
#


.PHONY: build clean test docker run

GO=CGO_ENABLED=0 GO111MODULE=on go
GOCGO=CGO_ENABLED=1 GO111MODULE=on go

DOCKERS=docker_build_base docker_volume docker_config_seed docker_export_client docker_export_distro docker_core_data docker_core_metadata docker_core_command docker_support_logging docker_support_notifications docker_sys_mgmt_agent docker_support_scheduler docker_security_secrets_setup docker_security_proxy_setup docker_security_secretstore_setup
# DOCKERS+=docker_app_service_configurable docker_support_rulesengine
DOCKERS+=docker_consul docker_volume docker_devices docker_ui
.PHONY: $(DOCKERS)

MICROSERVICES=cmd/config-seed/config-seed cmd/export-client/export-client cmd/export-distro/export-distro cmd/core-metadata/core-metadata cmd/core-data/core-data cmd/core-command/core-command cmd/support-logging/support-logging cmd/support-notifications/support-notifications cmd/sys-mgmt-executor/sys-mgmt-executor cmd/sys-mgmt-agent/sys-mgmt-agent cmd/support-scheduler/support-scheduler cmd/security-secrets-setup/security-secrets-setup cmd/security-proxy-setup/security-proxy-setup cmd/security-secretstore-setup/security-secretstore-setup

.PHONY: $(MICROSERVICES)

VERSION=$(shell cat ./VERSION)
DOCKER_TAG=$(shell uname -m)-$(VERSION)

GOFLAGS=-ldflags "-X github.com/objectbox/edgex-objectbox.Version=$(VERSION)"

GIT_SHA=$(shell git rev-parse HEAD)

ARCH=$(shell uname -m)

build: $(MICROSERVICES)

cmd/config-seed/config-seed:
	$(GO) build $(GOFLAGS) -o $@ ./cmd/config-seed

cmd/core-metadata/core-metadata:
	$(GOCGO) build $(GOFLAGS) -o $@ ./cmd/core-metadata

cmd/core-data/core-data:
	$(GOCGO) build $(GOFLAGS) -o $@ ./cmd/core-data

cmd/core-command/core-command:
	$(GOCGO) build $(GOFLAGS) -o $@ ./cmd/core-command

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

cmd/security-secrets-setup/security-secrets-setup:
	$(GO) build $(GOFLAGS) -o ./cmd/security-secrets-setup/security-secrets-setup ./cmd/security-secrets-setup

cmd/security-proxy-setup/security-proxy-setup:
	$(GO) build $(GOFLAGS) -o ./cmd/security-proxy-setup/security-proxy-setup ./cmd/security-proxy-setup

cmd/security-secretstore-setup/security-secretstore-setup:
	$(GO) build $(GOFLAGS) -o ./cmd/security-secretstore-setup/security-secretstore-setup ./cmd/security-secretstore-setup


clean:
	rm -f $(MICROSERVICES)

test:
	GO111MODULE=on go test -coverprofile=coverage.out ./...
	GO111MODULE=on go vet ./...
	gofmt -l .
	[ "`gofmt -l .`" = "" ]
	./bin/test-go-mod-tidy.sh
	./bin/test-attribution-txt.sh

run:
	cd bin && ./edgex-launch.sh

run_docker:
	bin/edgex-docker-launch.sh $(EDGEX_DB)

docker: $(DOCKERS)

docker_build_base:
	docker build \
		-f build/base/Dockerfile \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edgex-build-base:$(GIT_SHA) \
		.

# docker_app_service_configurable:
# 	$(eval NAME := app-service-configurable)
# 	./build/get-upstream-repo.sh \
# 		$(NAME) \
# 		407c5695040f00cf2c69603f5ad8d39d9232cb8d \
# 		https://github.com/edgexfoundry/app-service-configurable.git
# 	docker build \
# 		--build-arg http_proxy \
# 		--build-arg https_proxy \
# 		--label "git_sha=$(GIT_SHA)" \
# 		-t objectboxio/edgex-$(NAME):$(GIT_SHA) \
# 		-t objectboxio/edgex-$(NAME):$(DOCKER_TAG) \
# 		build/checkouts/$(NAME)

# currently fails to build: Step 7/15 : COPY target/*.jar $APP_DIR/$APP - no source files were specified
# docker_support_rulesengine:
# 	$(eval NAME := support-rulesengine)
# 	./build/get-upstream-repo.sh \
# 		$(NAME) \
# 		35f17d8e59789244a7d634909f958bd3b73e8e81 \
# 		https://github.com/edgexfoundry/support-rulesengine.git
# 	docker build \
# 		--build-arg http_proxy \
# 		--build-arg https_proxy \
# 		--label "git_sha=$(GIT_SHA)" \
# 		-t objectboxio/edgex-$(NAME):$(GIT_SHA) \
# 		-t objectboxio/edgex-$(NAME):$(DOCKER_TAG) \
# 		build/checkouts/$(NAME)

docker_consul:
	./build/build-upstream-repo.sh \
		consul \
		50b620a4478e2ed0408c29de667cb3e11d03cd8a \
		https://github.com/edgexfoundry/docker-edgex-consul.git \
		$(GIT_SHA) \
		$(DOCKER_TAG)

docker_volume:
	./build/build-upstream-repo.sh \
		volume \
		9cc8536c83189269461638e6684ed6435b9b3075 \
		https://github.com/edgexfoundry/docker-edgex-volume.git \
		$(GIT_SHA) \
		$(DOCKER_TAG)

docker_devices:
	./build/build-upstream-repo.sh \
		device-virtual \
		b4f3ea612a0bba9fde6d44447b3e9c68bf64be13 \
		https://github.com/edgexfoundry/device-virtual-go.git \
		$(GIT_SHA) \
		$(DOCKER_TAG)

docker_ui:
	./build/build-upstream-repo.sh \
		ui \
		277361c19809286defc6bccd20bc08e304c6ef7c \
		https://github.com/edgexfoundry/edgex-ui-go.git \
		$(GIT_SHA) \
		$(DOCKER_TAG)

docker_config_seed:
	docker build \
		-f cmd/config-seed/Dockerfile \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edgex-core-config-seed:$(GIT_SHA) \
		-t objectboxio/edgex-core-config-seed:$(DOCKER_TAG) \
		.

docker_core_metadata:
	docker build \
		-f cmd/Dockerfile --build-arg service=core-metadata \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edgex-core-metadata:$(GIT_SHA) \
		-t objectboxio/edgex-core-metadata:$(DOCKER_TAG) \
		.

docker_core_data:
	docker build \
		-f cmd/Dockerfile --build-arg service=core-data \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edgex-core-data:$(GIT_SHA) \
		-t objectboxio/edgex-core-data:$(DOCKER_TAG) \
		.

docker_core_command:
	docker build \
		-f cmd/Dockerfile --build-arg service=core-command \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edgex-core-command:$(GIT_SHA) \
		-t objectboxio/edgex-core-command:$(DOCKER_TAG) \
		.

docker_export_client:
	docker build \
		-f cmd/Dockerfile --build-arg service=export-client \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edgex-export-client:$(GIT_SHA) \
		-t objectboxio/edgex-export-client:$(DOCKER_TAG) \
		.

docker_export_distro:
	docker build \
		-f cmd/export-distro/Dockerfile \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edgex-export-distro:$(GIT_SHA) \
		-t objectboxio/edgex-export-distro:$(DOCKER_TAG) \
		.

docker_support_logging:
	docker build \
		-f cmd/support-logging/Dockerfile \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edgex-support-logging:$(GIT_SHA) \
		-t objectboxio/edgex-support-logging:$(DOCKER_TAG) \
		.

docker_support_notifications:
	docker build \
		-f cmd/Dockerfile --build-arg service=support-notifications \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edgex-support-notifications:$(GIT_SHA) \
		-t objectboxio/edgex-support-notifications:$(DOCKER_TAG) \
		.

docker_support_scheduler:
	docker build \
		-f cmd/Dockerfile --build-arg service=support-scheduler \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edgex-support-scheduler:$(GIT_SHA) \
		-t objectboxio/edgex-support-scheduler:$(DOCKER_TAG) \
		.

docker_sys_mgmt_agent:
	docker build \
		-f cmd/sys-mgmt-agent/Dockerfile \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edgex-sys-mgmt-agent:$(GIT_SHA) \
		-t objectboxio/edgex-sys-mgmt-agent:$(DOCKER_TAG) \
		.

docker_security_secrets_setup:
	# TODO: split this up and rename it when security-secrets-setup is a
	# different container
	docker build \
		-f cmd/security-secrets-setup/Dockerfile \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edgex-secret-store:$(GIT_SHA) \
		-t objectboxio/edgex-secret-store:$(DOCKER_TAG) \
		.

docker_security_proxy_setup:
	docker build \
		-f cmd/security-proxy-setup/Dockerfile \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edgex-security-proxy-setup:$(GIT_SHA) \
		-t objectboxio/edgex-security-proxy-setup:$(DOCKER_TAG) \
		.

docker_security_secretstore_setup:
		docker build \
		-f cmd/security-secretstore-setup/Dockerfile \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edgex-security-secretstore-setup:$(GIT_SHA) \
		-t objectboxio/edgex-security-secretstore-setup:$(DOCKER_TAG) \
		.

raml_verify:
	docker run --rm --privileged \
		-v $(PWD):/raml-verification -w /raml-verification \
		nexus3.edgexfoundry.org:10003/edgex-docs-builder:$(ARCH) \
		/scripts/raml-verify.sh