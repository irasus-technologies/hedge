#
# Contributors: BMC Helix, Inc.
#
# (c) Copyright 2020-2025 BMC Helix, Inc.
#
# SPDX-License-Identifier: Apache-2.0
#
# Default values: placeholders to pull latest version for dev, will be overridden from Makefile

ARG GO_BASE=golang:latest
ARG DOCKER_BASE=docker:dind

FROM ${GO_BASE} AS builder

# Switch to root user for installation
USER root

ARG ALPINE_PKG_BASE="make git"
ARG ALPINE_PKG_EXTRA=""

# RUN sed -e 's/dl-cdn[.]alpinelinux.org/nl.alpinelinux.org/g' -i~ /etc/apk/repositories
RUN apk add --update --no-cache ${ALPINE_PKG_BASE} ${ALPINE_PKG_EXTRA}
RUN apk add --update --no-cache pkgconfig build-base openssh-client ca-certificates
RUN apk add --no-cache fuse-overlayfs

# add git for go modules
WORKDIR /hedge
COPY go.mod .
RUN go mod download all
COPY . .

#create binary file in the container
RUN GO111MODULE=on go build -ldflags "-X hedge.Version=latest" -o hedge-ml-sandbox ./edge-ml-service/cmd/hedge-ml-sandbox/main.go

#==========================================================================================================================================================

FROM docker:dind-rootless

# Switch to root user for installation
USER root

# Create a group 'edgex' with GID 2001 and a user 'edgex' with UID 2002
RUN addgroup -S -g 2001 edgex && adduser -S -u 2002 -G edgex edgex

WORKDIR /
COPY --from=builder /hedge/hedge-ml-sandbox /hedge-ml-sandbox
COPY --from=builder /hedge/edge-ml-service/cmd/hedge-ml-sandbox/res /res
COPY --from=builder /hedge/edge-ml-service/cmd/hedge-ml-sandbox/entrypoint.sh /entrypoint.sh

# Change ownership of the application to the non-root user
RUN chown -R edgex:edgex /res /hedge-ml-sandbox /etc/ssl/certs
#RUN mkdir -p /res/traning/data
RUN chmod -R u+rwx /res
RUN chmod u+x /hedge-ml-sandbox

# Ensure the XDG_RUNTIME_DIR exists with correct permissions
ENV XDG_RUNTIME_DIR=/home/edgex/.run/2002
ENV DOCKER_HOST=unix:///home/edgex/.run/2002/docker.sock

RUN mkdir -p /home/edgex/.run/2002 && \
    chown -R edgex:edgex /home/edgex/.run/2002

RUN chown edgex:edgex /entrypoint.sh && chmod +x /entrypoint.sh

RUN echo "2002:100000:65536" >> /etc/subuid \
 && echo "2002:100000:65536" >> /etc/subgid

# Create Docker data directory and fix ownership
RUN mkdir -p /home/edgex/.local/share/docker && \
    chown -R edgex:edgex /home/edgex/.local/share

# Create job directory directory and fix ownership
RUN mkdir -p /tmp/jobs && \
    chown -R edgex:edgex /tmp/jobs

# Switch to the non-root 'edgex' user
USER 2002

ENTRYPOINT ["/entrypoint.sh"]
# dockerd-entrypoint.sh is part of docker:dind-rootless and we need it to start as non-root user
#ENTRYPOINT ["sh", "-c", "dockerd-entrypoint.sh & ./hedge-ml-sandbox"]
CMD ["-cp=consul.http://edgex-core-consul:8500", "--registry", "--configDir=./res"]

