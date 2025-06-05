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

ARG ALPINE_PKG_BASE="make git"
ARG ALPINE_PKG_EXTRA=""

# RUN sed -e 's/dl-cdn[.]alpinelinux.org/nl.alpinelinux.org/g' -i~ /etc/apk/repositories
RUN apk add --update --no-cache ${ALPINE_PKG_BASE} ${ALPINE_PKG_EXTRA}

WORKDIR /hedge
COPY go.mod .
RUN go mod download all
COPY . .

#create binary file for container
RUN GO111MODULE=on go build -ldflags "-X hedge.Version=1.0.0" -o hedge-exp ./app-services/hedge-export/cmd/main.go

#=========================================================================================================================================================

# Next image - Copy built Go binary into new workspace

FROM ${ALPINE_BASE}

# Switch to root user for installation
USER root

# dumb-init is required as security-bootstrapper uses it in the entrypoint script
RUN apk add --update --no-cache ca-certificates dumb-init

# Create a group 'edgex' with GID 2001 and a user 'edgex' with UID 2002
RUN addgroup -S -g 2001 edgex && adduser -S -u 2002 -G edgex edgex

WORKDIR /
# Copy application files and configuration with 'edgex' user and group ownership
COPY --from=builder --chown=edgex:edgex /hedge/app-services/hedge-export/cmd/res /res
COPY --from=builder --chown=edgex:edgex /hedge/hedge-exp /hedge-export

# Ensure the main binary is executable
RUN chmod u+x /hedge-export

EXPOSE 48099

# Switch to the non-root 'edgex' user
USER 2002

ENTRYPOINT ["./hedge-export"]
CMD ["-cp=consul.http://edgex-core-consul:8500", "--registry", "--configDir=./res"]
