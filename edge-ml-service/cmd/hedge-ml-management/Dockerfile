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

RUN apk add --update --no-cache ${ALPINE_PKG_BASE} ${ALPINE_PKG_EXTRA}

WORKDIR /hedge
COPY go.mod .
RUN go mod download all
COPY . .

RUN GO111MODULE=on go build -ldflags "-X hedge.Version=latest" -o hedge-ml-management ./edge-ml-service/cmd/hedge-ml-management/main.go

#==========================================================================================================================================================

FROM ${ALPINE_BASE}

# Switch to root user for installation
USER root

RUN apk add --update --no-cache ca-certificates dumb-init zip attr

# Create a group 'edgex' with GID 2001 and a user 'edgex' with UID 2002
RUN addgroup -S -g 2001 edgex && adduser -S -u 2002 -G edgex edgex

EXPOSE 48095

# Turn off secure mode for examples. Not recommended for production

WORKDIR /
# Copy application files and configuration with 'edgex' user and group ownership
COPY --from=builder /hedge/edge-ml-service/cmd/hedge-ml-management/res /res
COPY --from=builder /hedge/hedge-ml-management /hedge-ml-management

# Copy and make the cleanup script executable
COPY --from=builder /hedge/edge-ml-service/cmd/hedge-ml-management/scripts/clean_zip.sh /usr/local/bin/clean_zip.sh
RUN chown edgex:edgex /usr/local/bin/clean_zip.sh && chmod +x /usr/local/bin/clean_zip.sh

# Change ownership to 'edgex' user and set broad permissions on /res
RUN chown -R edgex:edgex /res /hedge-ml-management

# Create job directory directory and fix ownership
RUN mkdir -p /tmp/jobs && \
    chown -R edgex:edgex /tmp/jobs

# Create job directory directory and fix ownership
RUN mkdir -p /res/models && \
    chown -R edgex:edgex /res/models

RUN mkdir -p /res/models
RUN chmod -R u+rwx /res
RUN chmod u+x /hedge-ml-management

# Switch to the non-root 'edgex' user
USER 2002

ENTRYPOINT ["./hedge-ml-management"]
CMD ["--cp=consul.http://edgex-core-consul:8500","--registry","--configDir=./res"]
