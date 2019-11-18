# Builds https://github.com/edgexfoundry/edgex-ui-go
# Docker image for building EdgeX Foundry Config Seed
FROM golang:1.12-alpine AS builder

RUN apk update && apk add make git

# set the working directory
WORKDIR /build

ARG src_git_sha
RUN git init && \
	git remote add origin https://github.com/edgexfoundry/edgex-ui-go.git && \
	git fetch origin $src_git_sha --depth=1 && \
	git reset --hard FETCH_HEAD

RUN make prepare
RUN make cmd/edgex-ui-server/edgex-ui-server

FROM alpine:3.6

EXPOSE 4000

COPY --from=builder /build/cmd/edgex-ui-server /edgex-ui
WORKDIR /edgex-ui
ENTRYPOINT ["./edgex-ui-server"]