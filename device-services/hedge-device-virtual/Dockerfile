#
# Contributors: BMC Helix, Inc.
#
# (c) Copyright 2020-2025 BMC Helix, Inc.
#
# SPDX-License-Identifier: Apache-2.0
#

# Default values: placeholders to pull latest version for dev, will be overridden from Makefile
ARG GO_BASE=golang:latest
ARG ALPINE_BASE=alpine:latest

FROM ${GO_BASE} AS builder

# Switch to root user for installation
USER root

ARG ALPINE_PKG_BASE="make git openssh-client gcc libc-dev zeromq-dev libsodium-dev"
ARG ALPINE_PKG_EXTRA=""
RUN apk add --update --no-cache ${ALPINE_PKG_BASE} ${ALPINE_PKG_EXTRA}

WORKDIR /hedge
COPY go.mod .
#COPY device-services/hedge-device-virtual/_build ./edgexfoundry/
RUN go mod download all
COPY . .

RUN GO111MODULE=on go build -ldflags "-X hedge.Version=latest" -o hedge-device-virtual ./device-services/hedge-device-virtual/cmd/main.go

#==========================================================================================================================================================

FROM ${ALPINE_BASE}

# Switch to root user for installation
USER root

RUN apk add --update --no-cache ca-certificates dumb-init

# Create a group 'edgex' with GID 2001 and a user 'edgex' with UID 2002
RUN addgroup -S -g 2001 edgex && adduser -S -u 2002 -G edgex edgex

ENV APP_PORT=49991
EXPOSE $APP_PORT

WORKDIR /
COPY --from=builder /hedge/device-services/hedge-device-virtual/cmd /cmd
COPY --from=builder /hedge/device-services/hedge-device-virtual/cmd/res /res
COPY --from=builder /hedge/hedge-device-virtual /hedge-device-virtual

# Change ownership of the application to the non-root user
RUN chown -R edgex:edgex /cmd /res /hedge-device-virtual /etc/ssl/certs

# Ensure the binary is executable by the owner
RUN chmod u+x /hedge-device-virtual

# Switch to the non-root 'edgex' user
USER 2002

ENTRYPOINT ["./hedge-device-virtual"]
CMD ["--cp=consul.http://edgex-core-consul:8500","--registry","--configDir=./res"]
