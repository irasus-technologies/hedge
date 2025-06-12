#
# Original contributor
# Copyright (c) 2006-2020, Salvatore Sanfilippo
#
# Contributors: BMC Helix, Inc.
#
# (c) Copyright 2020-2025 BMC Helix, Inc.
#
FROM cgr.dev/chainguard/wolfi-base:latest

ARG REDIS_VERSION=7.4.0-r0

RUN apk update && apk add --no-cache \
    libgcc \
    bash \
    curl \
    redis-cli=${REDIS_VERSION} \
    redis=${REDIS_VERSION}

WORKDIR /data

# Create a group 'edgex' with GID 2001 and a user 'edgex' with UID 2002
RUN addgroup -S -g 2001 edgex && adduser -S -u 2002 -G edgex edgex

RUN mkdir /usr/local/bin

COPY --chown=edgex:edgex external/redis/docker-entrypoint.sh /usr/local/bin

RUN chmod +x /usr/local/bin/docker-entrypoint.sh

# Switch to non-root user
USER 2002

EXPOSE 6379