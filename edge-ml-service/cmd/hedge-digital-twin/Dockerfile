#
# Contributors: BMC Helix, Inc.
#
# (c) Copyright 2020-2025 BMC Helix, Inc.
#
# SPDX-License-Identifier: Apache-2.0
#

#build stage
# Default values: placeholders to pull latest version for dev, will be overridden from Makefile
ARG GO_BASE=golang:alpine
ARG ALPINE_BASE=alpine:latest

FROM ${GO_BASE} AS builder

# Switch to root user for installation
USER root

ARG ALPINE_PKG_BASE="make git"
ARG ALPINE_PKG_EXTRA=""

# RUN sed -e 's/dl-cdn[.]alpinelinux.org/nl.alpinelinux.org/g' -i~ /etc/apk/repositories
RUN apk add --update --no-cache ${ALPINE_PKG_BASE} ${ALPINE_PKG_EXTRA}

# add git for go modules
WORKDIR /hedge
COPY go.mod .
RUN go mod download all
COPY . .

#create binary file in the container
RUN GO111MODULE=on go build -ldflags "-X hedge.Version=latest" -o hedge-digital-twin ./edge-ml-service/cmd/hedge-digital-twin/main.go

#final stage
FROM ${ALPINE_BASE}

# Switch to root user for installation
USER root

# dumb-init is required as security-bootstrapper uses it in the entrypoint script
RUN apk add --update --no-cache ca-certificates dumb-init

# Turn off secure mode for examples. Not recommended for production
#ENV EDGEX_SECURITY_SECRET_STORE=false

# Create a group 'edgex' with GID 2001 and a user 'edgex' with UID 2002
RUN addgroup -S -g 2001 edgex && adduser -S -u 2002 -G edgex edgex

WORKDIR /
# Copy application files and configuration with 'edgex' user and group ownership
COPY --from=builder --chown=edgex:edgex /hedge/edge-ml-service/cmd/hedge-digital-twin/res /res
COPY --from=builder --chown=edgex:edgex /hedge/hedge-digital-twin /hedge-digital-twin

# Ensure the main binary is executable
RUN chmod -R u+rwx /res
RUN chmod u+x /hedge-digital-twin

EXPOSE 59860

# Switch to the non-root 'edgex' user
USER 2002

ENTRYPOINT ["./hedge-digital-twin"]
CMD ["--cp=consul.http://edgex-core-consul:8500","--registry","--configDir=./res"]
