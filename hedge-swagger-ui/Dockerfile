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

ARG ALPINE_PKG_BASE="make git gcc libc-dev libsodium-dev jq"
ARG ALPINE_PKG_EXTRA=""

RUN apk add --update --no-cache ${ALPINE_PKG_BASE} ${ALPINE_PKG_EXTRA}

# add git for go modules
WORKDIR /hedge
COPY go.mod .
RUN go mod download all
COPY . .

RUN go get github.com/swaggo/swag/cmd/swag
RUN go run github.com/swaggo/swag/cmd/swag init --parseInternal=true --generalInfo=doc.go --pd=true --ot=json --output=/hedge/hedge-swagger-ui/res/swagger/.tmp/

RUN jq -S . /hedge/hedge-swagger-ui/res/swagger/swagger.json > /hedge/hedge-swagger-ui/res/swagger/.tmp/f1.json && jq -S . /hedge/hedge-swagger-ui/res/swagger/.tmp/swagger.json > /hedge/hedge-swagger-ui/res/swagger/.tmp/f2.json

RUN	if diff /hedge/hedge-swagger-ui/res/swagger/.tmp/f1.json /hedge/hedge-swagger-ui/res/swagger/.tmp/f2.json >/dev/null; then \
  		echo "Swagger specification is up-to-date."; \
	else \
  		echo "Warning: Discrepancy found between ./hedge-swagger-ui/res/swagger/swagger.json in Git and generated swagger specification. Please regenerate swagger using make generate-swagger-spec and execute Git commit and push. "; \
  		cp "/hedge/hedge-swagger-ui/res/swagger/.tmp/swagger.json" "/hedge/hedge-swagger-ui/res/swagger/swagger.json"; \
  	fi

#create binary file in the container
RUN CGO_ENABLED=1 GO111MODULE=on go build -ldflags "-X hedge.Version=latest" -o swagger-ui ./hedge-swagger-ui/main.go

#==========================================================================================================================================================
FROM ${ALPINE_BASE}

# Switch to root user for installation
USER root

RUN apk add --update --no-cache ca-certificates dumb-init

# Turn off secure mode for examples. Not recommended for production
#ENV EDGEX_SECURITY_SECRET_STORE=false

# Create a group 'edgex' with GID 2001 and a user 'edgex' with UID 2002
RUN addgroup -S -g 2001 edgex && adduser -S -u 2002 -G edgex edgex

WORKDIR /
# Copy application files and configuration with 'edgex' user and group ownership
COPY --from=builder --chown=edgex:edgex /hedge/swagger-ui /hedge-swagger-ui
COPY --from=builder --chown=edgex:edgex /hedge/hedge-swagger-ui/res /res

# Ensure the main binary is executable
RUN chmod u+x /hedge-swagger-ui

# Switch to the non-root 'edgex' user
USER 2002

WORKDIR /
ENTRYPOINT ["./hedge-swagger-ui"]
