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
#COPY app-services/helix-remediate/_build ./edgexfoundry/
RUN go mod download all
COPY . .
#RUN apk info -a zeromq-dev

#create binary file in the container
RUN GO111MODULE=on go build -ldflags "-X hedge.Version=latest" -o hedge-remediate ./app-services/hedge-remediate/main.go

#==========================================================================================================================================================

# Next image - Copy built Go binary into new workspace
FROM ${ALPINE_BASE}

# Switch to root user for installation
USER root

#WORKDIR /hedge/app-services/hedge-event

# dumb-init is required as security-bootstrapper uses it in the entrypoint script
RUN apk add --update --no-cache ca-certificates dumb-init

#WORKDIR /hedge/app-services/helix-remediate
#RUN apk --no-cache add zeromq

# Turn off secure mode for examples. Not recommended for production
#ENV EDGEX_SECURITY_SECRET_STORE=false

# Create a group 'edgex' with GID 2001 and a user 'edgex' with UID 2002
RUN addgroup -S -g 2001 edgex && adduser -S -u 2002 -G edgex edgex

WORKDIR /
# Copy application files and configuration with 'edgex' user and group ownership
COPY --from=builder --chown=edgex:edgex /hedge/app-services/hedge-remediate/res /res
COPY --from=builder --chown=edgex:edgex /hedge/hedge-remediate /hedge-remediate

# Ensure the main binary is executable
RUN chmod u+x /hedge-remediate

# Expose the application port
EXPOSE 48105

# Switch to the non-root 'edgex' user
USER 2002

ENTRYPOINT ["./hedge-remediate"]
CMD ["-cp=consul.http://edgex-core-consul:8500", "--registry", "--configDir=./res"]
