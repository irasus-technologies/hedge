#
# Contributors: BMC Helix, Inc.
#
# (c) Copyright 2020-2025 BMC Helix, Inc.
#
# SPDX-License-Identifier: Apache-2.0
#

#build stage
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
#COPY hedge-admin/_build ./edgexfoundry/
RUN go mod download all
COPY . .
#RUN apk info -a zeromq-dev

#create binary file for container
RUN GO111MODULE=on go build -ldflags "-X hedge.Version=1.0.0" -o hedge-admin-service ./app-services/hedge-admin/cmd/main.go

#=========================================================================================================================================================

# Next image - Copy built Go binary into new workspace
FROM ${ALPINE_BASE}

# Switch to root user for installation
USER root

# dumb-init is required as security-bootstrapper uses it in the entrypoint script
RUN apk add --update --no-cache ca-certificates dumb-init curl

# Turn off secure mode for examples. Not recommended for production
#ENV EDGEX_SECURITY_SECRET_STORE=false

# Create a group 'edgex' with GID 2001 and a user 'edgex' with UID 2002
RUN addgroup -S -g 2001 edgex && adduser -S -u 2002 -G edgex edgex

WORKDIR /
# Copy application files and configuration with 'edgex' user and group ownership
COPY --from=builder --chown=edgex:edgex /hedge/app-services/hedge-admin/cmd/res /res
COPY --from=builder --chown=edgex:edgex /hedge/hedge-admin-service /hedge-admin

# Create contents directory and set ownership to 'edgex' user and group
RUN mkdir ./contents

# Copy the contents and vertical specific solutions with 'edgex' user and group ownership
COPY --from=builder --chown=edgex:edgex /hedge/contents/ ./contents/

# Ensure the main binary is executable
RUN chown -R edgex:edgex ./contents/
RUN chmod u+x /hedge-admin

#Only for grafana directories within /contents, update the permissions to 472:0 ie grafana:root
RUN chown -R 472:0 ./contents/grafana

EXPOSE 48098

# Switch to the non-root 'edgex' user
USER 2002

ENTRYPOINT ["./hedge-admin"]
CMD ["-cp=consul.http://edgex-core-consul:8500", "--registry", "--configDir=./res"]
